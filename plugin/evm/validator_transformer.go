package evm

import (
	"github.com/flare-foundation/coreth/core/state/validatordb"
)

// ValidatorTransformer describes a component that transforms a validator set
// in-place by applying a transformation function to it.
type ValidatorTransformer interface {
	Transform(validators []*validatordb.Validator) []*validatordb.Validator
}
