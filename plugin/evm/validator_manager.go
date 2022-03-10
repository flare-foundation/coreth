// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	lru "github.com/hashicorp/golang-lru"

	"github.com/flare-foundation/coreth/core/types"
	"github.com/flare-foundation/flare/snow/validators"
)

const (
	validatorSetsCacheSize = 2
)

type ValidatorManager struct {
	ftso *FTSO
	sets *lru.ARCCache
}

func NewValidatorManager(ftso *FTSO) (*ValidatorManager, error) {

	sets, _ := lru.NewARC(validatorSetsCacheSize)
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

	providers, err := v.ftso.ProvidersForEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get validators (epoch: %d): %w", epoch, err)
	}

	set := validators.NewSet()
	for _, provider := range providers {

		validator, err := v.ftso.ValidatorForProviderAtEpoch(epoch, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get validator (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		votepower, err := v.ftso.VotepowerForProviderAtEpoch(epoch, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get votepower (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		rewards, err := v.ftso.RewardsForProviderAtEpoch(epoch, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get rewards (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		weight := uint64(math.Log(float64(votepower)) * float64(rewards))
		err = set.AddWeight(validator, weight)
		if err != nil {
			return nil, fmt.Errorf("could not add weight: %w", err)
		}
	}

	v.sets.Add(epoch, set)

	return set, nil
}
