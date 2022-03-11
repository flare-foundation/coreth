// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	lru "github.com/hashicorp/golang-lru"

	"github.com/flare-foundation/coreth/core/types"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
)

const (
	validatorSetsCacheSize = 4
)

type ValidatorManager struct {
	ftso *FTSO
	sets *lru.Cache
}

func NewValidatorManager(ftso *FTSO) (*ValidatorManager, error) {

	sets, _ := lru.New(validatorSetsCacheSize)
	v := ValidatorManager{
		ftso: ftso,
		sets: sets,
	}

	return &v, nil
}

func (v *ValidatorManager) ValidatorSet(header *types.Header) (validators.Set, error) {

	epoch, err := v.ftso.EpochForTimestamp(header.Time)
	if err != nil {
		return nil, fmt.Errorf("could not get epoch for timestamp: %w", err)
	}

	entry, ok := v.sets.Get(epoch)
	if ok {
		return entry.(validators.Set), nil
	}

	providers, err := v.ftso.ProvidersForEpoch(epoch - 1)
	if err != nil {
		return nil, fmt.Errorf("could not get validators (epoch: %d): %w", epoch, err)
	}

	var total float64
	weights := make(map[ids.ShortID]float64, len(providers))
	for _, provider := range providers {

		validator, err := v.ftso.ValidatorForProviderAtEpoch(epoch-1, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get validator (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		votepower, err := v.ftso.VotepowerForProviderAtEpoch(epoch-1, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get votepower (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		reward, err := v.ftso.RewardForProviderAtEpoch(epoch-1, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get rewards (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		weight := math.Pow(votepower, 0.25) * (reward / votepower)
		weights[validator] = weight
		total += weight
	}

	set := validators.NewSet()
	normalizer := float64(math.MaxUint64) / total
	for validator, weight := range weights {
		err := set.AddWeight(validator, uint64(weight*normalizer))
		if err != nil {
			return nil, fmt.Errorf("could not add validator weight: %w", err)
		}
	}

	v.sets.Add(epoch, set)

	return set, nil
}
