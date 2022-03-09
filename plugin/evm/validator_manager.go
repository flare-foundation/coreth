// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"

	"github.com/flare-foundation/flare/snow/validators"
)

const (
	validatorCacheSize = 2
)

type ValidatorManager struct {
	ftso *FTSO
	sets *lru.Cache
}

func NewValidatorManager(ftso *FTSO) (*ValidatorManager, error) {

	sets, _ := lru.New(validatorCacheSize)
	v := ValidatorManager{
		ftso: ftso,
		sets: sets,
	}

	return &v, nil
}

func (v *ValidatorManager) ValidatorSet(time uint64) (validators.Set, error) {

	epoch, err := v.ftso.Epoch(time)
	if err != nil {
		return nil, fmt.Errorf("could not get epoch for time: %w", err)
	}

	entry, ok := v.sets.Get(epoch)
	if ok {
		return entry.(validators.Set), nil
	}

	validatorIDs, err := v.ftso.Validators(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get validators (epoch: %d): %w", epoch, err)
	}

	set := validators.NewSet()
	for _, validatorID := range validatorIDs {
		weight, err := v.ftso.Rewards(validatorID, epoch)
		if err != nil {
			return nil, fmt.Errorf("could not get validator weight (validator: %x): %w", validatorID, err)
		}
		err = set.AddWeight(validatorID, weight)
		if err != nil {
			return nil, fmt.Errorf("could not add weight for validator: %w", err)
		}
	}

	_ = v.sets.Add(epoch, set)

	return set, nil
}
