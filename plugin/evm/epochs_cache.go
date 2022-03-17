// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

var DefaultEpochsConfig = EpochsConfig{
	CacheSize: 16,
}

type EpochsConfig struct {
	CacheSize uint
}

type EpochsOption func(*EpochsConfig)

func WithCacheSize(size uint) EpochsOption {
	return func(cfg *EpochsConfig) {
		cfg.CacheSize = size
	}
}

type EpochsCache struct {
	epochs Epochs
	cache  *lru.Cache
}

func NewEpochsCache(epochs Epochs, opts ...EpochsOption) *EpochsCache {

	cfg := DefaultEpochsConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	cache, _ := lru.New(int(cfg.CacheSize))
	e := EpochsCache{
		epochs: epochs,
		cache:  cache,
	}

	return &e
}

func (e *EpochsCache) Details(epoch uint64) (EpochDetails, error) {

	entry, ok := e.cache.Get(epoch)
	if ok {
		return entry.(EpochDetails), nil
	}

	details, err := e.epochs.Details(epoch)
	if err != nil {
		return EpochDetails{}, fmt.Errorf("could not get epoch details: %w", err)
	}

	e.cache.Add(epoch, details)

	return details, nil
}
