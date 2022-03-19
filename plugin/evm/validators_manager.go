// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"

	lru "github.com/hashicorp/golang-lru"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
)

var DefaultManagerCacheConfig = CacheConfig{
	CacheSize: 4,
}

type EpochMapper interface {
	ByTimestamp(timestamp uint64) (uint64, error)
}

type ValidatorRetriever interface {
	ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error)
}

type ValidatorsManager struct {
	blockchain *core.BlockChain
	epochs     EpochMapper
	validators ValidatorRetriever
	sets       *lru.Cache
}

func NewValidatorsManager(blockchain *core.BlockChain, epochs EpochMapper, validators ValidatorRetriever, opts ...CacheOption) *ValidatorsManager {

	cfg := DefaultManagerCacheConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	sets, _ := lru.New(int(cfg.CacheSize))
	v := ValidatorsManager{
		blockchain: blockchain,
		epochs:     epochs,
		validators: validators,
		sets:       sets,
	}

	return &v
}

func (v *ValidatorsManager) ByBlock(blockID common.Hash) (validators.Set, error) {

	header := v.blockchain.GetHeaderByHash(blockID)
	if header == nil {
		return nil, fmt.Errorf("unknown block (hash: %x)", hash)
	}

	epoch, err := v.epochs.ByTimestamp(header.Time)
	if err != nil && !errors.Is(err, errFTSONotDeployed) && !errors.Is(err, errFTSONotActive) {
		return nil, fmt.Errorf("could not get epoch (timestamp: %d): %w", header.Time, err)
	}

	entry, ok := v.sets.Get(epoch)
	if ok {
		return entry.(validators.Set), nil
	}

	validatorMap, err := v.validators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve validators (epoch: %d): %w", epoch, err)
	}

	set := validators.NewSet()
	for id, weight := range validatorMap {
		err = set.AddWeight(id, weight)
		if err != nil {
			return nil, fmt.Errorf("could not add weight: %w", err)
		}
	}

	v.sets.Add(epoch, set)

	return set, nil
}
