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
	MinSteps:    4,
	Placeholder: 50_000,
}

type TransitionConfig struct {
	MinSteps    uint
	Placeholder uint64
}

type TransitionOption func(*TransitionConfig)

func WithMinSteps(steps uint) TransitionOption {
	return func(cfg *TransitionConfig) {
		cfg.MinSteps = steps
	}
}

// ValidatorsTransitioner transitions validators from a static set of validators
// to a growing set of dynamic validators over a number of smooth steps.
type ValidatorsTransitioner struct {
	validators map[ids.ShortID]uint64
	providers  ValidatorRetriever
	cfg        TransitionConfig
}

// NewValidatorsTransitioner creates a transition from the given validators to the
// validators retrieved from the given providers.
func NewValidatorsTransitioner(validators map[ids.ShortID]uint64, providers ValidatorRetriever, opts ...TransitionOption) *ValidatorsTransitioner {

	cfg := DefaultTransitionConfig
	for _, opt := range opts {
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

	// First, we get the dynamic set of validators from the underlying retriever.
	providers, err := v.providers.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get provider validators: %w", err)
	}

	// If there are non, we always return the full set of static validators.
	if len(providers) == 0 {
		return v.validators, nil
	}

	// Otherwise, we try to identify the number of the step we are at in our transition
	// from the static to the dynamic set. It can range from zero to the minimum number
	// of steps; at zero, we still return the static set; at minimum steps reached, we
	// return only the dynamic set.
	size := uint(len(v.validators))
	steps := uint(0)
Loop:
	for try := uint(1); try <= v.cfg.MinSteps; try++ {
		thresholds := findThresholds(size, try, v.cfg.MinSteps)
		for i, threshold := range thresholds {
			e := epoch - uint64(i)
			if e > epoch {
				break Loop
			}
			providers, err := v.providers.ByEpoch(e)
			if err != nil {
				return nil, fmt.Errorf("could not get providers: %w", err)
			}
			if uint(len(providers)) < threshold {
				break Loop
			}
			steps = try
		}
	}

	// If we are not ready to take any steps yet, we stick with the default
	// validator set still.
	if steps == 0 {
		return v.validators, nil
	}

	// If we have reached the minimum number of steps, we can return the dynamic
	// set.
	if steps == v.cfg.MinSteps {
		return providers, nil
	}

	// If we are somewhere in-between, we need to balance the list of providers
	// with the list of static default validators. First, we choose a deterministic
	// list of validator IDs from the default validators, cut off at the appropriate
	// percentage depending on steps
	validators := make([]ids.ShortID, 0, len(v.validators))
	for validator := range v.validators {
		validators = append(validators, validator)
	}
	sort.Slice(validators, func(i int, j int) bool {
		return bytes.Compare(validators[i][:], validators[j][:]) < 0
	})
	cutoff := (size - steps*size/v.cfg.MinSteps)
	validators = validators[:cutoff]

	// Then, we try to balance the weights between the default validators and the
	// validators from the FTSO. In order to do so, we calculate the total weight
	// of providers. From that, we derive the total weigth we should have based on
	// the step we are on and the proportion the providers should have. Then we
	// derive the weight default validators should have based on the total.
	providerWeight := uint64(0)
	for _, weight := range providers {
		providerWeight += weight
	}
	totalWeight := (providerWeight) * uint64(v.cfg.MinSteps) / uint64(steps)
	validatorWeight := totalWeight * uint64(v.cfg.MinSteps-steps) / uint64(v.cfg.MinSteps)

	// Finally, we add the selected default validators to the set of the providers
	// with the weights we have calculated.
	for _, validator := range validators {
		providers[validator] = validatorWeight
	}

	return providers, nil
}

func findThresholds(size uint, steps uint, min uint) []uint {
	thresholds := make([]uint, 0, steps)
	for i := uint(0); i < steps; i++ {
		threshold := (steps - i) * size / min
		thresholds = append(thresholds, threshold)
	}
	return thresholds
}
