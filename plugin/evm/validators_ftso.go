// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/flare/ids"
)

var DefaultFTSOConfig = FTSOConfig{
	RootDegree: 4,
}

type FTSOConfig struct {
	RootDegree uint
}

type FTSOOption func(*FTSOConfig)

func WithRootDegree(degree uint) FTSOOption {
	return func(cfg *FTSOConfig) {
		cfg.RootDegree = degree
	}
}

type EpochHandler interface {
	StartHeight(epoch uint64) (uint64, error)
}

type StateShifter interface {
	ToBlock(hash common.Hash) (FTSO, error)
}

type FTSO interface {
	Indices() ([]uint64, error)
	Providers(index uint64) ([]common.Address, error)
	Validator(provider common.Address) (ids.ShortID, error)
	Votepower(provider common.Address) (float64, error)
	Rewards(provider common.Address) (float64, error)
}

// ValidatorsFTSO is responsible for retrieving the set of validators for the FTSO
// data providers, in accordance with the defined formula and configured root degree.
type ValidatorsFTSO struct {
	blockchain *core.BlockChain
	epochs     EpochHandler
	shift      StateShifter
	cfg        FTSOConfig
}

func NewValidatorsFTSO(blockchain *core.BlockChain, epochs EpochHandler, shift StateShifter, opts ...FTSOOption) *ValidatorsFTSO {

	cfg := DefaultFTSOConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	v := ValidatorsFTSO{
		blockchain: blockchain,
		epochs:     epochs,
		shift:      shift,
		cfg:        cfg,
	}

	return &v
}

func (v *ValidatorsFTSO) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	height, err := v.epochs.StartHeight(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get start time: %w", err)
	}

	header := v.blockchain.GetHeaderByNumber(height)
	if header == nil {
		return nil, fmt.Errorf("unknown header (height: %d)", height)
	}

	hash := header.Hash()
	ftso, err := v.shift.ToBlock(hash)
	if err != nil {
		return nil, fmt.Errorf("could not shift state (block: %x): %w", hash, err)
	}

	indices, err := ftso.Indices()
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO indices: %w", err)
	}

	providerMap := make(map[common.Address]struct{})
	for _, index := range indices {
		providers, err := ftso.Providers(index)
		if err != nil {
			return nil, fmt.Errorf("could not get  providers (index: %d): %w", index, err)
		}
		for _, provider := range providers {
			providerMap[provider] = struct{}{}
		}
	}

	providers := make([]common.Address, 0, len(providerMap))
	for provider := range providerMap {
		providers = append(providers, provider)
	}

	validators := make(map[ids.ShortID]uint64, len(providers))
	for _, provider := range providers {

		id, err := ftso.Validator(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get validator ID (provider: %x): %w", provider, err)
		}

		votepower, err := ftso.Votepower(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get vote power (provider: %x): %w", provider, err)
		}

		rewards, err := ftso.Rewards(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get rewards (provider: %x): %w", provider, err)
		}

		weight := uint64(math.Pow(votepower, 1.0/float64(v.cfg.RootDegree)) * (rewards / votepower))

		validators[id] = weight
	}

	return validators, nil
}
