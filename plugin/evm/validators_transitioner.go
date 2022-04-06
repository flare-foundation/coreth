// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"bytes"
	"errors"
	"fmt"
	"sort"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
	"github.com/flare-foundation/flare/utils/logging"
)

var DefaultTransitionConfig = TransitionConfig{
	StepSize: 1,
}

type TransitionConfig struct {
	StepSize uint
}

type TransitionOption func(*TransitionConfig)

func WithStepSize(size uint) TransitionOption {
	return func(cfg *TransitionConfig) {
		cfg.StepSize = size
	}
}

type ValidatorsPersister interface {
	Persist(epoch uint64, validators map[ids.ShortID]uint64) error
}

// ValidatorsTransitioner transitions validators from a static set of validators
// to a growing set of dynamic validators over a number of smooth steps.
type ValidatorsTransitioner struct {
	log             logging.Logger
	retrieveActive  ValidatorsRetriever
	retrieveDefault ValidatorsRetriever
	retrieveFTSO    ValidatorsRetriever
	store           ValidatorsPersister
	cfg             TransitionConfig
}

// NewValidatorsTransitioner creates a transition from the given default validators
// to the validators retrieved from the given FTSO validators retriever.
func NewValidatorsTransitioner(log logging.Logger, retrieveDefault ValidatorsRetriever, retrieveFTSO ValidatorsRetriever, retrieveActive ValidatorsRetriever, store ValidatorsPersister, options ...TransitionOption) *ValidatorsTransitioner {

	cfg := DefaultTransitionConfig
	for _, opt := range options {
		opt(&cfg)
	}

	v := ValidatorsTransitioner{
		log:             log,
		retrieveDefault: retrieveDefault,
		retrieveFTSO:    retrieveFTSO,
		retrieveActive:  retrieveActive,
		store:           store,
		cfg:             cfg,
	}

	return &v

}

// ByEpoch returns a set of validators that is given by a smooth transition from an
// original static set of validators, to a new dynamic set of validators that is
// growing and continuously takes up more of the total set until the original static
// validators have been entirely phased out.
func (v *ValidatorsTransitioner) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	fmt.Println("=========================================")

	// Get the default validators for the requested epoch.
	defaultValidators, err := v.retrieveDefault.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get current default validators for transition: %w", err)
	}

	// At epoch zero, we return the default validators.
	if epoch == 0 {
		v.log.Debug("returning default validators for epoch zero (%d)", len(defaultValidators))
		return defaultValidators, nil
	}

	// Now that we know we are not at epoch zero, we can get the FTSO validators
	// from the previous epoch. We have to use the previous epoch's FTSO validators
	// because we need access to the full distributed rewards for the epoch, and
	// they have not yet been determined for the currently retrieveActive epoch.
	ftsoValidators, err := v.retrieveFTSO.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not get previous FTSO validators for transition: %w", err)
	}

	// If there are no FTSO validators for the previous epoch, we return the default
	// validators, as none of them can currently be phased out.
	if len(ftsoValidators) == 0 {
		v.log.Debug("returning default validators in absence of FTSO validators (%d)", len(defaultValidators))
		return defaultValidators, nil
	}

	// Otherwise, get the retrieveActive validators from the previous epoch to see how
	// many we have to transition. If none are available, we have to recurse.
	previousValidators, err := v.retrieveActive.ByEpoch(epoch - 1)
	if errors.Is(err, leveldb.ErrNotFound) {
		v.log.Debug("retrieveActive validators not available, recursing into epoch %d", epoch-1)
		previousValidators, err = v.ByEpoch(epoch - 1)
	}
	if err != nil {
		return nil, fmt.Errorf("could not get previous active validators for transition: %w", err)
	}

	// We can then count the number of default validators that were included in the
	// retrieveActive validators from the previous epoch.
	included := 0
	for validatorID := range previousValidators {
		_, ok := defaultValidators[validatorID]
		if ok {
			included++
		}
	}

	// Now, we calculate how many additional default validators we can remove at
	// this epoch.
	remove := 0
	for {

		// If we have reached the maximum number of default validators we can remove
		// for this epoch, stop the loop.
		if remove >= int(v.cfg.StepSize) {
			break
		}

		// If the number of remove default validators we can remove this epoch
		// has reached the remaining number of default validators, stop loop too.
		if remove >= included {
			break
		}

		// If the number of available FTSO validators is insufficient to remove
		// remove default validators, stop as well.
		if len(ftsoValidators) <= len(defaultValidators)-included+remove {
			break
		}

		remove++
	}

	// If there were no default validators included in the previous epoch at all,
	// we have already completed the transition from default validators to FTSO
	// validators, and we simply return the FTSO validators as retrieveActive validators.
	if remove == included {
		v.log.Debug("returning FTSO validators on completed transition (%d)", len(ftsoValidators))
		return ftsoValidators, nil
	}

	// We then select the given number of included default validators. In order to
	// make the selection deterministic, we sort the validator IDs for all default
	// validators and then cut it off at the number of included ones.
	cutoff := included - remove
	defaultIDs := make([]ids.ShortID, 0, len(defaultValidators))
	for defaultID := range defaultValidators {
		defaultIDs = append(defaultIDs, defaultID)
	}
	sort.Slice(defaultIDs, func(i int, j int) bool {
		return bytes.Compare(defaultIDs[i][:], defaultIDs[j][:]) < 0
	})
	defaultIDs = defaultIDs[:cutoff]

	v.log.Debug("reducing default validators (previous: %d, rnext: %d)", included, cutoff)

	// Once we fix FTSO validators and default validators, we can no longer use the
	// default configured weight for default validators. Instead we use a proportional
	// average. For example, if 3/10 default validators have been phased out, meaning
	// 7/10 default validators are still in the set, then the default validators should
	// have 70% of the total weight, and the FTSO validators should have 30% of the
	// total weight.
	proportionalWeight := uint64(0)
	for _, weight := range ftsoValidators {
		proportionalWeight += weight
	}
	proportionalWeight /= uint64(len(defaultValidators) - cutoff)

	// We then add all available FTSO validators to the retrieveActive validators first,
	// followed by the remaining default validators with the calculated proportional
	// average weight.
	transitionValidators := make(map[ids.ShortID]uint64, len(ftsoValidators)+len(defaultIDs))
	for ftsoValidator, weight := range ftsoValidators {
		transitionValidators[ftsoValidator] = weight
		v.log.Debug("adding provider validator %s (weight: %d)", ftsoValidator.PrefixedString(constants.NodeIDPrefix), weight)
	}
	for _, defaultID := range defaultIDs {
		transitionValidators[defaultID] = proportionalWeight
		v.log.Debug("adding default validator %s (weight: %d)", defaultID.PrefixedString(constants.NodeIDPrefix), proportionalWeight)
	}

	err = v.store.Persist(epoch, transitionValidators)
	if err != nil {
		return nil, fmt.Errorf("could not persist retrieveActive validators after transition: %w", err)
	}

	return transitionValidators, nil
}
