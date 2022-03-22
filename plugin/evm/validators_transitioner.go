// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/flare-foundation/flare/ids"
)

// ValidatorsTransitioner transitions validators from a static set of validators
// to a growing set of dynamic validators over a number of smooth steps.
type ValidatorsTransitioner struct {
	validators map[ids.ShortID]uint64
	providers  ValidatorRetriever
}

// NewValidatorsTransitioner creates a transition from the given default validators
// to the validators retrieved from the given FTSO validators retriever.
func NewValidatorsTransitioner(validators map[ids.ShortID]uint64, providers ValidatorRetriever) *ValidatorsTransitioner {

	v := ValidatorsTransitioner{
		validators: validators,
		providers:  providers,
	}

	return &v

}

// ByEpoch returns a set of validators that is given by a smooth transition from an
// original static set of validators, to a new dynamic set of validators that is
// growing and continuously takes up more of the total set until the original static
// validators have been entirely phased out.
func (v *ValidatorsTransitioner) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	// We need to get the FTSO providers for the previous epoch, so we need to
	// check we are not at epoch zero.
	if epoch < 1 {
		return v.validators, nil
	}

	// We start by getting the FTSO validators from the previous epoch. Since that
	// epoch is over, votepower and rewards are available and we can compute the
	// weights for them.
	providers, err := v.providers.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve FTSO validators for previous epoch: %w", err)
	}

	// If there are non, we return the default validator set. This is an important
	// point, as this is where we leave the recursion, where we decide how many
	// default validators to keep.
	if len(providers) == 0 {
		return v.validators, nil
	}

	// At this point, some providers are available. We have to decide how many default
	// validators to keep. First, we check how many default validators were included
	// in the last epoch.
	previous, err := v.ByEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve active validators for previous epoch: %w", err)
	}
	include := 0
	for validator := range previous {
		_, ok := v.validators[validator]
		if ok {
			include++
		}
	}

	// If there were no default validators in the previous set, we don't need any
	// now either.
	if include == 0 {
		return providers, nil
	}

	// If we have enough FTSO validators to fill one more spot, diminish by one
	// the default validators we will include.
	if len(providers) > len(v.validators)-include {
		include--
	}

	// We have to sort the default validators deterministically, so that we pick
	// the same ones across nodes.
	validators := make([]ids.ShortID, 0, len(v.validators))
	for validator := range v.validators {
		validators = append(validators, validator)
	}
	sort.Slice(validators, func(i int, j int) bool {
		return bytes.Compare(validators[i][:], validators[j][:]) < 0
	})

	// Then we can limit the number to what was determined earlier.
	validators = validators[:include]

	// Next, we try to make sure that the default validators have proportionally
	// the same average weighting as the FTSO validators.
	providerWeight := uint64(0)
	for _, weight := range providers {
		providerWeight += weight
	}
	providerWeight /= uint64((len(v.validators) - include))

	// Finally, we add the selected default validators to the set of the validators
	// with the weights we have calculated.
	active := make(map[ids.ShortID]uint64, len(providers)+len(validators))
	for provider, weight := range providers {
		active[provider] = weight
	}
	for _, validator := range validators {
		active[validator] = providerWeight
	}

	return active, nil
}
