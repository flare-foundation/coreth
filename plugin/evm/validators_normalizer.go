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

	return v.calcWeightRatio(validators), nil
}

func (v *ValidatorsNormalizer) calcWeightRatio(validators map[ids.ShortID]uint64) map[ids.ShortID]uint64 {
	var totalWeight uint64
	for _, wght := range validators {
		totalWeight += wght
	}

	v.log.Debug("normalizing weight from %d to %d", totalWeight, math.MaxInt32)

	ratio := math.MaxInt64 / totalWeight
	normalized := make(map[ids.ShortID]uint64, len(validators))
	for val, wght := range validators {
		normalized[val] = wght * ratio
	}

	v.log.Debug("new normalized validator set: %#v", normalized)

	return normalized
}
