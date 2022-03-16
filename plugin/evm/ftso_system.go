// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core"
)

type FTSOSystem struct {
	blockchain *core.BlockChain
	addresses  FTSOAddresses
	abis       FTSOABIs
	methods    FTSOMethods
}

func NewFTSOSystem(blockchain *core.BlockChain) *FTSOSystem {

	f := FTSOSystem{
		blockchain: blockchain,
		addresses:  DefaultFTSOAddresses,
		abis:       DefaultFTSOABIs,
		methods:    DefaultFTSOMethods,
	}

	return &f
}

func (f *FTSOSystem) EpochInfo(epoch uint64) (EpochInfo, error) {

	header := f.blockchain.CurrentHeader()

	call := NewFTSOCaller(f.blockchain, header.Hash())

	submitter := FTSOContract{
		address: f.addresses.Submitter,
		abi:     f.abis.Submitter,
	}

	var managerAddress common.Address
	err := call.
		OnContract(submitter).
		Execute(f.methods.ManagerAddress).
		Decode(&managerAddress)
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get manager address: %w", err)
	}

	var seconds big.Int
	manager := FTSOContract{
		address: managerAddress,
		abi:     f.abis.Manager,
	}
	err = call.OnContract(manager).
		Execute(f.methods.EpochSeconds).
		Decode(&seconds)
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get epoch seconds: %w", err)
	}

	var startHeight, startTime big.Int
	err = call.OnContract(manager).
		Execute(f.methods.EpochInfo).
		Decode(nil, &startHeight, &startTime)
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get epoch info: %w", err)
	}

	info := EpochInfo{
		StartHeight: startHeight.Uint64(),
		StartTime:   startTime.Uint64(),
		EndTime:     startTime.Uint64() + seconds.Uint64(),
	}

	return info, nil
}

func (f *FTSOSystem) Snapshot(hash common.Hash) (*FTSOSnapshot, error) {

	s := FTSOSnapshot{
		system: f,
		hash:   hash,
	}

	return &s, nil

}
