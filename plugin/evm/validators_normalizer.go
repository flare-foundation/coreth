// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
	"github.com/flare-foundation/flare/utils/logging"
)

// ValidatorsNormalizer is responsible for normalizing validators, so that their
// weights come out to the same total, irrespective of their original weights or
// how many validators are in a set.
type ValidatorsNormalizer struct {
	log      logging.Logger
	retrieve ValidatorsRetriever
}

// NewValidatorsNormalizer wraps a new validators retriever in the normalizer, making
// sure that all sets retrieved from the wrapper retriever have the same total weight.
func NewValidatorsNormalizer(log logging.Logger, retrieve ValidatorsRetriever) *ValidatorsNormalizer {

	v := ValidatorsNormalizer{
		log:      log,
		retrieve: retrieve,
	}

	return &v
}

// ByEpoch retrieves the validators from the underlying retriever and normalizes
// their weights so that they always have the same total approximate weight.
func (v *ValidatorsNormalizer) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	validators, err := v.retrieve.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve validators for normalizing: %w", err)
	}

	if len(validators) == 0 {
		return validators, nil
	}

	var totalWeight uint64
	for _, weight := range validators {
		totalWeight += weight
	}

	ratio := math.MaxInt64 / totalWeight
	normalized := make(map[ids.ShortID]uint64, len(validators))
	for validator, weight := range validators {
		v.log.Debug("normalizing weight for %s from %d to %d", validator.PrefixedString(constants.NodeIDPrefix), weight, weight*ratio)
		normalized[validator] = weight * ratio
	}

	return normalized, nil
}
