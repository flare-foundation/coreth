// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
)

var DefaultTransitionConfig = TransitionConfig{
	MinSteps: 4,
}

type TransitionConfig struct {
	MinSteps uint
}

type TransitionOption func(*TransitionConfig)

func WithMinSteps(steps uint) TransitionOption {
	return func(cfg *TransitionConfig) {
		cfg.MinSteps = steps
	}
}

// ValidatorsTransitioner transitions a set of validators from one retriever to
// the second retriever in a smooth deterministic manner over a pre-configured
// minimum number of steps.
type ValidatorsTransitioner struct {
	ids        []ids.ShortID
	validators ValidatorRetriever
	cfg        TransitionConfig
}

func NewValidatorsTransitioner(ids []ids.ShortID, validators ValidatorRetriever, opts ...TransitionOption) *ValidatorsTransitioner {

	cfg := DefaultTransitionConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	v := ValidatorsTransitioner{
		ids:        ids,
		validators: validators,
		cfg:        cfg,
	}

	return &v

}

func (v *ValidatorsTransitioner) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	set, err := v.validators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve first set: %w", err)
	}

	return set, nil
}

// var defaultValidators []ids.ShortID

// // If we have less than 4 reward epochs, we always return the default
// // validators.
// if epoch < 3 {
// 	return defaultValidators, nil
// }

// // We get all raw validator sets for the four last reward epochs, from the
// // oldest to the newest set.
// sets := make([]validators.Set, 0, 4)
// for e := epoch - 3; e <= epoch; e++ {
// 	set, err := v.FTSOValidators(e)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not get FTSO validators (epoch: %d): %w", e, err)
// 	}
// 	sets = append(sets, set)
// }

// // The default set is always all of the available FTSO validators as a basis.
// var set validators.Set
// err = set.Set(sets[0].List())
// if err != nil {
// 	return nil, fmt.Errorf("could not set validator set: %w", err)
// }

// // We use the maximum weight for the weight of the default validators.
// var max uint64
// for _, validator := range set.List() {
// 	if validator.Weight() > max {
// 		max = validator.Weight()
// 	}
// }

// // Then, we will add a number of default validators to that set that are
// // equivalent to losing only 25% of default validators per epoch.
// OuterLoop:
// for r := 4; r >= 1; r-- {
// 	for e := 0; e > r; e++ {
// 		threshold := (4 - e)
// 		if sets[e].Len() < (threshold * len(defaultValidators) / 4) {
// 			continue OuterLoop
// 		}
// 	}
// 	for _, validator := range defaultValidators[:r*len(defaultValidators)/4] {
// 		err = set.AddWeight(validator, max)
// 		if err != nil {
// 			return nil, fmt.Errorf("could not add weight (validator: %x, weight: %d): %w", validator, max, err)
// 		}
// 	}
// 	break
// }

// set := validators.NewSet()
// normalizer := float64(math.MaxUint64) / total
// for validator, weight := range weights {
// 	err := set.AddWeight(validator, uint64(weight*normalizer))
// 	if err != nil {
// 		return nil, fmt.Errorf("could not add validator weight: %w", err)
// 	}
// }
