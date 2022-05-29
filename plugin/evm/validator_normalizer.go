package evm

import (
	"github.com/flare-foundation/coreth/core/state/validatordb"
)

// ValidatorNormalizer is a validator transformer which transforms a given validator set
// by proportionally increasing all validator weights up to or just below a configured
// maximum ceiling of total weight.
type ValidatorNormalizer struct {
	ceil uint64
}

// NewValidatorNormalizer creates a new validator normalizer using the given ceil as the
// ceiling up to which the total validator weights are increased.
func NewValidatorNormalizer(ceil uint64) *ValidatorNormalizer {

	v := ValidatorNormalizer{
		ceil: ceil,
	}

	return &v
}

// Transform takes a validator set and increases the weights of the validators proportionally
// up to or just below the configured ceiling of total weight.
func (v *ValidatorNormalizer) Transform(validators []*validatordb.Validator) []*validatordb.Validator {

	// Calculate the total weight of all validators.
	totalWeight := uint64(0)
	for _, validator := range validators {
		totalWeight += validator.Weight
	}

	// Calculate the ratio of current total weight divided by ceiling of desired total weight.
	ratio := float64(v.ceil) / float64(totalWeight)

	// Create the normalized validator set by multiplying each weight by the ratio.
	for _, validator := range validators {
		validator.Weight = uint64(float64(validator.Weight) * ratio)
	}

	return validators
}
