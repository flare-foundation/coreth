// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	"github.com/flare-foundation/flare/ids"
)

type ValidatorsNormalizer struct {
	validators ValidatorRetriever
}

func NewValidatorsNormalizer(validators ValidatorRetriever) *ValidatorsNormalizer {

	v := ValidatorsNormalizer{
		validators: validators,
	}

	return &v
}

func (v *ValidatorsNormalizer) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	validators, err := v.validators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve validators: %w", err)
	}

	var totalWeight uint64
	for _, weight := range validators {
		totalWeight += weight
	}

	ratio := math.MaxUint64 / totalWeight
	for validator, weight := range validators {
		validators[validator] = weight * ratio
	}

	return validators, nil
}
