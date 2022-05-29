package evm

import (
	"bytes"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/core/state/validatordb"
	"github.com/flare-foundation/flare/ids"
)

// AggregateFunc is an aggregation function that will return an aggregate weight resulting
// from a list of provided meights.
type AggregateFunc func([]uint64) uint64

// AggregateMax will pick the maximum weight from the input weights and return it.
func AggregateMax(weights []uint64) uint64 {
	max := uint64(0)
	for _, weight := range weights {
		if weight > max {
			max = weight
		}
	}
	return max
}

// AggregateAverage will calculate the average weight of the given weights and return
// it.
func AggregateAverage(weights []uint64) uint64 {
	total := uint64(0)
	for _, weight := range weights {
		total += weight
	}
	average := total / uint64(len(weights))
	return average
}

// ValidatorAdditioner is a validator transformer that adds a list of validators to
// a validator set, which each of the added validators receiving a resulting weight
// according to a configured aggregation function.
type ValidatorAdditioner struct {
	additional []ids.ShortID
	aggregate  AggregateFunc
}

// NewValidatorAdditioner creates a new validator additioner that adds the configured
// list of validators using the given aggregation function to determine the weights of
// the new validators.
func NewValidatorAdditioner(additional []ids.ShortID, aggregate AggregateFunc) *ValidatorAdditioner {

	v := ValidatorAdditioner{
		additional: additional,
		aggregate:  aggregate,
	}

	return &v
}

// Transform takes the given validator set and adds the configured list of validators to
// the resulting set, using the configured aggregation function to determine the weight
// of the added validators.
// NOTE: If the given validator set already contains one of the configured additional
// validators, the new weight calculated by the aggregate function will be added to
// the existing weight.
func (v *ValidatorAdditioner) Transform(validators []*validatordb.Validator) []*validatordb.Validator {

	// We use a provider address of all ones as placeholder for default validators.
	placeholder := common.BytesToAddress(bytes.Repeat([]byte{0xff}, 20))

	// Collect all of the weights and input them into the aggregation function to
	// get the weight to be added to the validators.
	weights := make([]uint64, 0, len(validators))
	for _, validator := range validators {
		weights = append(weights, validator.Weight)
	}
	added := v.aggregate(weights)

	// Create a lookups of deltas that should be applied against certain node IDs.
	deltas := make(map[ids.ShortID]uint64)
	for _, additionalID := range v.additional {
		deltas[additionalID] = added
	}

	// Go through the validators and apply the deltas to where the node IDs overlap.
	for _, validator := range validators {
		added, ok := deltas[validator.NodeID]
		if ok {
			validator.Providers = append(validator.Providers, placeholder)
			validator.Weight += added
			delete(deltas, validator.NodeID)
		}
	}

	// For every remaining delta, create a new validator and add it to the list,
	// then sort to keep a deterministic order.
	for nodeID, weight := range deltas {
		validator := validatordb.Validator{
			Providers: []common.Address{placeholder},
			NodeID:    nodeID,
			Weight:    weight,
		}
		validators = append(validators, &validator)
	}
	sort.Slice(validators, func(i int, j int) bool {
		return bytes.Compare(validators[i].NodeID[:], validators[j].NodeID[:]) < 0
	})

	return validators
}
