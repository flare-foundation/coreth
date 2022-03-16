// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
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

type FTSOCaller interface {
	EpochInfo(epoch uint64) (EpochInfo, error)
}

type EpochInfo struct {
	StartHeight uint64
	StartTime   uint64
	EndTime     uint64
}

type FTSOEpochs struct {
	call  FTSOCaller
	cache *lru.Cache
}

func NewFTSOEpochs(call FTSOCaller, opts ...EpochsOption) *FTSOEpochs {

	cfg := DefaultEpochsConfig

	cache, _ := lru.New(int(cfg.CacheSize))
	f := FTSOEpochs{
		call:  call,
		cache: cache,
	}

	return &f

}

func (f *FTSOEpochs) StartHeight(epoch uint64) (uint64, error) {

	entry, ok := f.cache.Get(epoch)
	if ok {
		return entry.(EpochInfo).StartHeight, nil
	}

	info, err := f.call.EpochInfo(epoch)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch info: %w", err)
	}

	f.cache.Add(epoch, info)

	return info.StartHeight, nil
}

func (f *FTSOEpochs) StartTime(epoch uint64) (uint64, error) {

	entry, ok := f.cache.Get(epoch)
	if ok {
		return entry.(EpochInfo).StartTime, nil
	}

	info, err := f.call.EpochInfo(epoch)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch info: %w", err)
	}

	f.cache.Add(epoch, info)

	return info.StartTime, nil
}

func (f *FTSOEpochs) EndTime(epoch uint64) (uint64, error) {

	entry, ok := f.cache.Get(epoch)
	if ok {
		return entry.(EpochInfo).EndTime, nil
	}

	info, err := f.call.EpochInfo(epoch)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch info: %w", err)
	}

	f.cache.Add(epoch, info)

	return info.EndTime, nil
}
