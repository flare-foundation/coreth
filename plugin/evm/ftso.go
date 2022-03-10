// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	lru "github.com/hashicorp/golang-lru"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/internal/ethapi"
	"github.com/flare-foundation/flare/ids"
)

const (
	epochsCacheSize = 16
)

type FTSO struct {
	blockchain *core.BlockChain
	last       uint64
	seconds    uint64
	epochs     *lru.ARCCache
}

func NewFTSO(blockchain *core.BlockChain) (*FTSO, error) {

	epochs, _ := lru.NewARC(epochsCacheSize)
	f := FTSO{
		blockchain: blockchain,
		seconds:    0,
		last:       0,
		epochs:     epochs,
	}

	seconds, err := f.RewardEpochDurationSeconds()
	if err != nil {
		return nil, fmt.Errorf("could not bootstrap reward epoch duration seconds: %w", err)
	}

	f.seconds = seconds

	return &f, nil
}

func (f *FTSO) EpochForTimestamp(timestamp uint64) (uint64, error) {

	epoch := f.last
	for {

		info, err := f.RewardEpochs(epoch)
		if err != nil {
			return 0, fmt.Errorf("could not get reward epochs info: %w", err)
		}

		if timestamp > info.end {
			epoch++
			continue
		}

		if timestamp < info.start {
			epoch--
			continue
		}

		f.last = epoch

		return epoch, nil
	}
}

func (f *FTSO) ProvidersForEpoch(epoch uint64) ([]common.Address, error) {
	return nil, nil
}

func (f *FTSO) ValidatorForProviderAtEpoch(epoch uint64, provider common.Address) (ids.ShortID, error) {
	return ids.ShortID{}, nil
}

func (f *FTSO) VotepowerForProviderAtEpoch(epoch uint64, provider common.Address) (uint64, error) {
	return 0, nil
}

func (f *FTSO) RewardsForProviderAtEpoch(epoch uint64, provider common.Address) (uint64, error) {
	return 0, nil
}

func (f *FTSO) RewardEpochDurationSeconds() (uint64, error) {

	values, err := f.call(f.blockchain.CurrentHeader().Hash(), ftsoManager, "rewardEpochDurationSeconds")
	if err != nil {
		return 0, fmt.Errorf("could not get reward epoch duration seconds: %w", err)
	}
	if len(values) != 1 {
		return 0, fmt.Errorf("wrong number of return values (have: %d, want: %d)", len(values), 1)
	}

	seconds := values[0].(*big.Int).Uint64()

	return seconds, nil
}

type rewardEpochs struct {
	start uint64
	end   uint64
}

func (f *FTSO) RewardEpochs(epoch uint64) (rewardEpochs, error) {

	entry, ok := f.epochs.Get(epoch)
	if ok {
		return entry.(rewardEpochs), nil
	}

	values, err := f.call(f.blockchain.CurrentHeader().Hash(), ftsoManager, "rewardEpochs", big.NewInt(0).SetUint64(epoch))
	if err != nil {
		return rewardEpochs{}, fmt.Errorf("could not get reward epochs: %w", err)
	}
	if len(values) != 3 {
		return rewardEpochs{}, fmt.Errorf("wrong number of return values  (have: %d, want: %d)", len(values), 3)
	}

	start := values[2].(*big.Int).Uint64()

	info := rewardEpochs{
		start: start,
		end:   start + f.seconds,
	}

	f.epochs.Add(epoch, info)

	return info, nil
}

func (f *FTSO) call(hash common.Hash, abi abi.ABI, method string, params ...interface{}) ([]interface{}, error) {

	data, err := abi.Pack(method, params...)
	if err != nil {
		return nil, fmt.Errorf("could not pack call data: %w", err)
	}

	header := f.blockchain.GetHeaderByHash(hash)
	if header == nil {
		return nil, fmt.Errorf("block hash unknown")
	}

	state, err := f.blockchain.StateAt(header.Root)
	if err != nil {
		return nil, fmt.Errorf("could not get blockchain state: %w", err)
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return nil, fmt.Errorf("could not convert arguments to message: %w", err)
	}

	vmConfig := f.blockchain.GetVMConfig()
	chainConfig := f.blockchain.Config()
	txContext := core.NewEVMTxContext(msg)
	blkContext := core.NewEVMBlockContext(header, f.blockchain, nil)
	evm := vm.NewEVM(blkContext, txContext, state, chainConfig, *vmConfig)
	defer evm.Cancel()

	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(evm, msg, gp)
	if err != nil {
		return nil, fmt.Errorf("could not apply message: %w", err)
	}
	if result.Err != nil {
		return nil, fmt.Errorf("could not execute transaction: %w", err)
	}

	values, err := abi.Unpack(method, result.ReturnData)
	if err != nil {
		return nil, fmt.Errorf("could not unpack return data: %w", err)
	}

	return values, nil
}
