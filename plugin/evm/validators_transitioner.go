// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/flare-foundation/flare/ids"
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

// ValidatorsTransitioner transitions validators from a static set of validators
// to a growing set of dynamic validators over a number of smooth steps.
type ValidatorsTransitioner struct {
	validators ValidatorsRetriever
	providers  ValidatorsRetriever
	cfg        TransitionConfig
}

// NewValidatorsTransitioner creates a transition from the given default validators
// to the validators retrieved from the given FTSO validators retriever.
func NewValidatorsTransitioner(validators ValidatorsRetriever, providers ValidatorsRetriever, options ...TransitionOption) *ValidatorsTransitioner {

	cfg := DefaultTransitionConfig
	for _, opt := range options {
		opt(&cfg)
	}

	v := ValidatorsTransitioner{
		validators: validators,
		providers:  providers,
		cfg:        cfg,
	}

	return &v

}

// ByEpoch returns a set of validators that is given by a smooth transition from an
// original static set of validators, to a new dynamic set of validators that is
// growing and continuously takes up more of the total set until the original static
// validators have been entirely phased out.
func (v *ValidatorsTransitioner) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	// Get the default validators for the requested epoch.
	validators, err := v.validators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get default validators: %w", err)
	}

	// We need to get the FTSO providers for the previous epoch, so we need to
	// check we are not at epoch zero. For reward epoch zero, we always return
	// the default validators.
	if epoch == 0 {
		return validators, nil
	}

	// Now that we know we are not at epoch zero, we can get the FTSO validators
	// from the previous epoch. We have to use the previous epoch's FTSO validators
	// because we need access to the full distributed rewards for the epoch, and
	// they have not yet been determined for the currently active epoch.
	providers, err := v.providers.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve FTSO validators for previous epoch: %w", err)
	}

	// If there are no FTSO validators for the previous epoch, we return the default
	// validators, as none of them can currently be phased out.
	if len(providers) == 0 {
		return validators, nil
	}

	// At this point, we have some FTSO validators available, and we have some default
	// validators available. In order to determine how many default validators to
	// phase out, we first retrieve the active validators from the previous epoch.
	previous, err := v.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve active validators for previous epoch: %w", err)
	}

	// We can then count the number of default validators that were included in the
	// active validators from the previous epoch.
	included := 0
	for validator := range previous {
		_, ok := validators[validator]
		if ok {
			included++
		}
	}

	// If there were no default validators included in the previous epoch at all,
	// we have already completed the transition from default validators to FTSO
	// validators, and we simply return the FTSO validators as active validators.
	if included == 0 {
		return providers, nil
	}

	// Now, we calculate how many additional default validators we can remove at
	// this epoch.
	additional := 0
	for {

		// If we have reached the maximum number of default validators we can remove
		// for this epoch, stop the loop.
		if additional >= int(v.cfg.StepSize) {
			break
		}

		// If the number of additional default validators we can remove this epoch
		// has reached the remaining number of default validators, stop loop too.
		if additional >= included {
			break
		}

		// If the number of available FTSO validators is insufficient to remove
		// additional default validators, stop as well.
		if len(providers) <= len(validators)-included+additional {
			break
		}

		additional++
	}

	// We then select the given number of included default validators. In order to
	// make the selection deterministic, we sort the validator IDs for all default
	// validators and then cut it off at the number of included ones.
	cutoff := included - additional
	validatorIDs := make([]ids.ShortID, 0, len(validators))
	for validatorID := range validators {
		validatorIDs = append(validatorIDs, validatorID)
	}
	sort.Slice(validatorIDs, func(i int, j int) bool {
		return bytes.Compare(validatorIDs[i][:], validatorIDs[j][:]) < 0
	})
	validatorIDs = validatorIDs[:cutoff]

	// Once we fix FTSO validators and default validators, we can no longer use the
	// default configured weight for default validators. Instead we use a proportional
	// average. For example, if 3/10 default validators have been phased out, meaning
	// 7/10 default validators are still in the set, then the default validators should
	// have 70% of the total weight, and the FTSO validators should have 30% of the
	// total weight.
	providerWeight := uint64(0)
	for _, weight := range providers {
		providerWeight += weight
	}
	providerWeight /= uint64(len(validators) - cutoff)

	// We then add all available FTSO validators to the active validators first,
	// followed by the remaining default validators with the calculated proportional
	// average weight.
	active := make(map[ids.ShortID]uint64, len(providers)+len(validators))
	for provider, weight := range providers {
		active[provider] = weight
	}
	for _, validatorID := range validatorIDs {
		active[validatorID] = providerWeight
	}

	return active, nil
}
