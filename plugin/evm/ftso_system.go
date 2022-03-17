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
	submitter  common.Address
	validation common.Address
}

type FTSOContracts struct {
	Submitter  EVMContract
	Registry   EVMContract
	Manager    EVMContract
	Rewards    EVMContract
	FTSO       EVMContract
	Whitelist  EVMContract
	Validation EVMContract
	Votepower  EVMContract
}

func NewFTSOSystem(blockchain *core.BlockChain, submitter common.Address, validation common.Address) *FTSOSystem {

	f := FTSOSystem{
		blockchain: blockchain,
		submitter:  submitter,
		validation: validation,
	}

	return &f
}

func (f *FTSOSystem) Contracts(epoch uint64) (FTSOContracts, error) {

	info, err := f.EpochInfo(epoch)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get epoch info: %w", err)
	}

	header := f.blockchain.GetHeaderByNumber(info.StartHeight)
	if header == nil {
		return FTSOContracts{}, fmt.Errorf("unknown header (height: %d)", info.StartHeight)
	}

	contracts := FTSOContracts{}

	return contracts, nil
}

func (f *FTSOSystem) EpochInfo(epoch uint64) (EpochInfo, error) {

	contracts := FTSOContracts{}

	var seconds big.Int
	err := BindEVM(f.blockchain).
		OnContract(contracts.Manager).
		Execute(EpochSeconds).
		Decode(&seconds)
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get epoch seconds: %w", err)
	}

	var startHeight, startTime *big.Int
	err = BindEVM(f.blockchain).
		OnContract(contracts.Manager).
		Execute(RewardEpoch).
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
