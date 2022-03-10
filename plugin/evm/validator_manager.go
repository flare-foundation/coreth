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

	var totalVotepower float64
	votepowers := make(map[ids.ShortID]float64, len(providers))
	rewards := make(map[ids.ShortID]float64, len(providers))
	for _, provider := range providers {

		validator, err := v.ftso.ValidatorForProviderAtEpoch(epoch, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get validator (epoch: %d, provider: %x): %w", epoch, provider, err)
		}

		votepower, err := v.ftso.VotepowerForProviderAtEpoch(epoch, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get votepower (epoch: %d, provider: %x): %w", epoch, provider, err)
		}
		totalVotepower += votepower
		votepowers[validator] = votepower

		reward, err := v.ftso.RewardForProviderAtEpoch(epoch, provider)
		if err != nil {
			return nil, fmt.Errorf("could not get rewards (epoch: %d, provider: %x): %w", epoch, provider, err)
		}
		rewards[validator] = reward
	}

	var totalAbsolute float64
	absolutes := make(map[ids.ShortID]float64)
	for validator, votepower := range votepowers {
		reward := rewards[validator]
		votepower /= totalVotepower
		absolute := math.Log(votepower) * reward
		absolutes[validator] = absolute
		totalAbsolute += absolute
	}

	set := validators.NewSet()
	ratio := float64(math.MaxUint64) / totalAbsolute
	for validator, absolute := range absolutes {
		relative := uint64(absolute * ratio)
		err := set.AddWeight(validator, relative)
		if err != nil {
			return nil, fmt.Errorf("could not add validator weight: %w", err)
		}
	}

	v.sets.Add(epoch, set)

	return set, nil
}
