// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare/ids"
)

type FTSOSnapshot struct {
	system    *FTSOSystem
	hash      common.Hash
	contracts FTSOContracts
}

func (f *FTSOSnapshot) Providers() ([]common.Address, error) {

	providerMap := make(map[common.Address]struct{})
	for _, serie := range f.contracts.Series {
		var addresses []common.Address
		err := BindEVM(f.system.blockchain).
			AtBlock(f.hash).
			OnContract(serie).
			Execute(DataProviders).
			Decode(&addresses)
		if err != nil {
			return nil, fmt.Errorf("could not get provider addresses (serie: %x): %w", serie.address, err)
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
	return ids.ShortEmpty, fmt.Errorf("not implemented")
}

func (f *FTSOSnapshot) Votepower(provider common.Address) (float64, error) {

	return 0, fmt.Errorf("not implemented")
}

func (f *FTSOSnapshot) Rewards(provider common.Address) (float64, error) {
	return 0, fmt.Errorf("not implemented")
}
