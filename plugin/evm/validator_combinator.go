package evm

import (
	"github.com/flare-foundation/coreth/core/state/validatordb"
)

// ValidatorCombinator is a validator transformer that combines multiple validator
// transformers, applying them to a validator set in sequence.
type ValidatorCombinator struct {
	transforms []ValidatorTransformer
}

// NewValidatorCombinator creates a new validator combinator that applies the given list
// of validator transformers in order.
func NewValidatorCombinator(transforms ...ValidatorTransformer) *ValidatorCombinator {

	v := ValidatorCombinator{
		transforms: transforms,
	}

	return &v
}

// Transform takes a validator set and applies the configured list of validator transformers
// to the given set in order, returning the resulting validator set after all transformers
// were applied.
func (v *ValidatorCombinator) Transform(validators []*validatordb.Validator) []*validatordb.Validator {

	for _, transform := range v.transforms {
		validators = transform.Transform(validators)
	}

	return validators
}
