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
	validators map[ids.ShortID]uint64
	providers  ValidatorRetriever
	cfg        TransitionConfig
}

// NewValidatorsTransitioner creates a transition from the given default validators
// to the validators retrieved from the given FTSO validators retriever.
func NewValidatorsTransitioner(validators map[ids.ShortID]uint64, providers ValidatorRetriever, options ...TransitionOption) *ValidatorsTransitioner {

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

	// We need to get the FTSO providers for the previous epoch, so we need to
	// check we are not at epoch zero.
	if epoch == 0 {
		return v.validators, nil
	}

	// We start by getting the FTSO validators from the previous epoch. Since that
	// epoch is over, votepower and rewards are available and we can compute the
	// weights for them.
	providers, err := v.providers.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve FTSO validators for previous epoch: %w", err)
	}

	// If there are none, we return the default validator set. This is an important
	// point, as this is where we leave the recursion, where we decide how many
	// default validators to keep.
	if len(providers) == 0 {
		return v.validators, nil
	}

	// At this point, we should start including some FTSO validators in the active
	// validator set. This depends on how many we included in the previous set, so
	// we recurse. The cache is there to avoid recursing all the way back to the
	// first transition on later requests.
	previous, err := v.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve active validators for previous epoch: %w", err)
	}

	// First, we compute the number of default validators included in the previous
	// active validators by increasing the number by one for each one found in the
	// previous set.
	numDefault := 0
	for validator := range previous {
		_, ok := v.validators[validator]
		if !ok {
			continue
		}
		numDefault++
	}

	// Then, we decrease the number of default validators that will be included for
	// this epoch by one as long as there are enough FTSO validators available, but
	// at most until there are none left or we reached the maximum step size.
	for i := 0; i < int(v.cfg.StepSize) && numDefault > 0; i++ {
		if len(providers) <= len(v.validators)-numDefault {
			break
		}
		numDefault--
	}

	// In order to always select the same default validators, we sort their IDs
	// deterministically, and then cut off at the number we should still include.
	validators := make([]ids.ShortID, 0, len(v.validators))
	for validator := range v.validators {
		validators = append(validators, validator)
	}
	sort.Slice(validators, func(i int, j int) bool {
		return bytes.Compare(validators[i][:], validators[j][:]) < 0
	})
	validators = validators[:numDefault]

	// Next, we try to make sure that the default validators have proportionally
	// the same average weighting as the FTSO validators.
	providerWeight := uint64(0)
	for _, weight := range providers {
		providerWeight += weight
	}
	providerWeight /= uint64(len(v.validators) - numDefault)

	// Finally, we add the selected default validators to the set of the validators
	// with the average weight we have calculated.
	active := make(map[ids.ShortID]uint64, len(providers)+len(validators))
	for provider, weight := range providers {
		active[provider] = weight
	}
	for _, validator := range validators {
		active[validator] = providerWeight
	}

	return active, nil
}
