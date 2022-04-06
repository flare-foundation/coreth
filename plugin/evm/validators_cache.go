// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"

	"github.com/flare-foundation/flare/ids"
)

var DefaultCacheConfig = CacheConfig{
	CacheSize: 8,
}

type CacheConfig struct {
	CacheSize uint
}

type CacheOption func(*CacheConfig)

func WithCacheSize(slots uint) CacheOption {
	return func(cfg *CacheConfig) {
		cfg.CacheSize = slots
	}
}

// ValidatorsCache wraps around a validator retriever and caches the results in
// order to improve retrieval performance.
type ValidatorsCache struct {
	validators ValidatorsRetriever
	cache      *lru.Cache
}

// NewValidatorCache creates a new LRU cache for validator retrieval with the
// configured cache size.
func NewValidatorsCache(validators ValidatorsRetriever, opts ...CacheOption) *ValidatorsCache {

	cfg := DefaultCacheConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	cache, _ := lru.New(int(cfg.CacheSize))
	v := ValidatorsCache{
		validators: validators,
		cache:      cache,
	}

	return &v
}

func (v *ValidatorsCache) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	entry, ok := v.cache.Get(epoch)
	if ok {
		return entry.(map[ids.ShortID]uint64), nil
	}

	validators, err := v.validators.ByEpoch(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve validators before caching: %w", err)
	}

	v.cache.Add(epoch, validators)

	return validators, nil
}
