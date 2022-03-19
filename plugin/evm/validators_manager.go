// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"
	"math/big"

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

// ValidatorsManager is responsible for choosing the strategy for building a validator
// set depending on a block hash. It might choose a legacy static validator set, as used
// before the hard fork upgrade, or a dynamic set of validators based on a transition to
// the FTSO validator set.
type ValidatorsManager struct {
	blockchain *core.BlockChain
	epochs     EpochMapper
	validators map[ids.ShortID]uint64
	providers  ValidatorRetriever
	cache      *lru.Cache
}

// NewValidatorsManager creates a new manager of validator sets. It uses the given
// blockchain to map block hashes to block headers, the given epoch mapper no map
// block timestamps to FTSO rewards epochs, the given validators as the legacy static
// validator set, and the given retriever to get the validator set based on FTSO
// data providers.
func NewValidatorsManager(blockchain *core.BlockChain, epochs EpochMapper, validators map[ids.ShortID]uint64, providers ValidatorRetriever, opts ...CacheOption) *ValidatorsManager {

	cfg := DefaultManagerCacheConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	cache, _ := lru.New(int(cfg.CacheSize))
	v := ValidatorsManager{
		blockchain: blockchain,
		epochs:     epochs,
		validators: validators,
		providers:  providers,
		cache:      cache,
	}

	return &v
}

// ByBlock returns the validator set that was active at the block with the given hash.
func (v *ValidatorsManager) ByBlock(hash common.Hash) (validators.Set, error) {

	// First, we check whether we have the header with the given hash available.
	header := v.blockchain.GetHeaderByHash(hash)
	if header == nil {
		return nil, fmt.Errorf("unknown block (hash: %x)", hash)
	}

	// If the hard fork was not active at the given block yet, we simply return the
	// default validator set, which corresponds to what we had before the upgrade.
	if !v.blockchain.Config().IsFlareHardFork1(big.NewInt(0).SetUint64(header.Time)) {
		return toSet(v.validators)
	}

	// If the hard fork is active, we try to map the header to an FTSO rewards epoch.
	// If the FTSO is not yet deployed, or not yet active, we simply go ahead with an
	// epoch value of zero as well.
	epoch, err := v.epochs.ByTimestamp(header.Time)
	if err != nil && !errors.Is(err, errFTSONotDeployed) && !errors.Is(err, errFTSONotActive) {
		return nil, fmt.Errorf("could not get epoch (timestamp: %d): %w", header.Time, err)
	}

	// If we have already cached the set of validators for this epoch, there is no
	// reason to go through any of the computations again.
	entry, ok := v.cache.Get(epoch)
	if ok {
		return entry.(validators.Set), nil
	}

	// Otherwise, we retrieve the validators from the injected retriever, which
	// applies any kind of transformations that might be required.
	validators, err := v.providers.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve validators (epoch: %d): %w", epoch, err)
	}

	// We then convert the validators to a set...
	set, err := toSet(validators)
	if err != nil {
		return nil, fmt.Errorf("could not convert validators to set: %w", err)
	}

	// ... and cache it before returning it.
	v.cache.Add(epoch, set)

	return set, nil
}

// toSet converts a list of validators and weights to a `validators.Set`.
func toSet(validatorMap map[ids.ShortID]uint64) (validators.Set, error) {
	set := validators.NewSet()
	for validator, weight := range validatorMap {
		err := set.AddWeight(validator, weight)
		if err != nil {
			return nil, fmt.Errorf("could not add weight: %w", err)
		}
	}
	return set, nil
}
