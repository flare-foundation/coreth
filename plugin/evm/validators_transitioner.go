// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

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

// ValidatorsTransitioner transitions a set of validators from one retriever to
// the second retriever in a smooth deterministic manner over a pre-configured
// minimum number of steps.
type ValidatorsTransitioner struct {
	validators []ids.ShortID
	providers  ValidatorRetriever
	cfg        TransitionConfig
}

func NewValidatorsTransitioner(validators []ids.ShortID, providers ValidatorRetriever, opts ...TransitionOption) *ValidatorsTransitioner {

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

func (v *ValidatorsTransitioner) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	size := uint(len(v.validators))
	steps := uint(0)
Loop:
	for try := uint(1); try <= v.cfg.MinSteps; try++ {
		thresholds := findThresholds(size, try, v.cfg.MinSteps)
		for i, threshold := range thresholds {
			e := epoch - uint64(i)
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

	providers, err := v.providers.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get provider: %w", err)
	}

	cutoff := (size - steps*size/v.cfg.MinSteps)
	validators := v.validators[:cutoff]

	totalWeight := uint64(0)
	for _, weight := range providers {
		totalWeight += weight
	}
	averageWeight := totalWeight / uint64(size)
	for _, validator := range validators {
		providers[validator] = averageWeight
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
