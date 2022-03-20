// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

type FTSOSnapshot struct {
	system    *FTSOSystem
	epoch     uint64
	power     common.Hash
	start     common.Hash
	end       common.Hash
	contracts FTSOContracts
}

func (f *FTSOSnapshot) Providers() ([]common.Address, error) {

	var indices []*big.Int
	err := BindEVM(f.system.blockchain).
		AtBlock(f.start).
		OnContract(f.contracts.Registry).
		Execute(SeriesIndices).
		Decode(&indices)
	if err != nil {
		return nil, fmt.Errorf("could not get series indices: %w", err)
	}

	providerMap := make(map[common.Address]struct{})
	for _, index := range indices {
		var addresses []common.Address
		err := BindEVM(f.system.blockchain).
			AtBlock(f.start).
			OnContract(f.contracts.Whitelist).
			Execute(DataProviders, index).
			Decode(&addresses)
		if err != nil {
			return nil, fmt.Errorf("could not get provider addresses (index: %d): %w", index, err)
		}
		for _, address := range addresses {
			providerMap[address] = struct{}{}
		}
	}

	providers := make([]common.Address, 0, len(providerMap))
	for provider := range providerMap {
		providers = append(providers, provider)
	}

	return providers, nil
}

func (f *FTSOSnapshot) Validator(provider common.Address) (ids.ShortID, error) {

	var validator [20]byte
	err := BindEVM(f.system.blockchain).
		AtBlock(f.start).
		OnContract(f.contracts.Validation).
		Execute(ProviderNode, provider).
		Decode(&validator)
	if errors.Is(err, errNoReturnData) {
		return ids.ShortEmpty, errRegistryNotDeployed
	}
	if err != nil {
		return ids.ShortEmpty, fmt.Errorf("could not get provider node: %w", err)
	}

	return ids.ShortID(validator), nil
}

func (f *FTSOSnapshot) Votepower(provider common.Address) (float64, error) {

	vpInt := &big.Int{}
	err := BindEVM(f.system.blockchain).
		AtBlock(f.power).
		OnContract(f.contracts.Votepower).
		Execute(ProviderVotepower, provider).
		Decode(&vpInt)
	if err != nil {
		return 0, fmt.Errorf("could not get provider votepower: %w", err)
	}

	vpFloat := big.NewFloat(0).SetInt(vpInt)
	votepower, _ := vpFloat.Float64()

	return votepower, nil
}

func (f *FTSOSnapshot) Rewards(provider common.Address) (float64, error) {

	epoch := big.NewInt(0).SetUint64(f.epoch)

	rwInt := &big.Int{}
	err := BindEVM(f.system.blockchain).
		AtBlock(f.end).
		OnContract(f.contracts.Rewards).
		Execute(ProviderRewards, epoch, provider).
		Decode(&rwInt, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get provider rewards: %w", err)
	}

	rwFloat := big.NewFloat(0).SetInt(rwInt)
	rewards, _ := rwFloat.Float64()

	return rewards, nil
}
