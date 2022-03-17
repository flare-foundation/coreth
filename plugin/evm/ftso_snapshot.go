// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type FTSOSnapshot struct {
	system *FTSOSystem
	hash   common.Hash
}

func (f *FTSOSnapshot) Indices() ([]uint64, error) {

	// call := NewFTSOCaller(f.system.blockchain, f.hash)

	// submitter := FTSOContract{
	// 	address: f.system.addresses.Submitter,
	// 	abi:     f.system.abis.Submitter,
	// }

	// var registryAddress common.Address
	// err := call.
	// 	OnContract(submitter).
	// 	Execute(f.system.methods.RegistryAddress).
	// 	Decode(&registryAddress)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not get manager address: %w", err)
	// }

	// registry := FTSOContract{
	// 	address: registryAddress,
	// 	abi:     f.system.abis.Registry,
	// }

	// var values []*big.Int
	// err = call.OnContract(registry).
	// 	Execute(f.system.methods.AssetIndices).
	// 	Decode(&values)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not execute asset indices call: %w", err)
	// }

	// indices := make([]uint64, 0, len(values))
	// for _, index := range values {
	// 	indices = append(indices, index.Uint64())
	// }

	return nil, fmt.Errorf("not implemented")
}

func (f *FTSOSnapshot) Providers(index uint64) ([]common.Address, error) {
	return nil, fmt.Errorf("not implemented")
}

func (f *FTSOSnapshot) Votepower(provider common.Address) (float64, error) {
	return 0, fmt.Errorf("not implemented")
}

func (f *FTSOSnapshot) Rewards(provider common.Address) (float64, error) {
	return 0, fmt.Errorf("not implemented")
}
