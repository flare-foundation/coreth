// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/internal/ethapi"
)

type FTSOWrapper struct {
	state     *state.StateDB
	addresses FTSOAddresses
	abis      FTSOABIs
	methods   FTSOMethods
}

func NewFTSOWrapper(state *state.StateDB) *FTSOWrapper {

	f := FTSOWrapper{
		state:     state,
		addresses: DefaultFTSOAddresses,
		abis:      DefaultFTSOABIs,
		methods:   DefaultFTSOMethods,
	}

	return &f
}

func (f *FTSOWrapper) EpochInfo(epoch uint64) (EpochInfo, error) {

	seconds, err := f.rewardSeconds()
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get reward seconds: %w", err)
	}

}

func (f *FTSOWrapper) Indices() ([]uint64, error) {
}

func (f *FTSOWrapper) Providers(index uint64) ([]common.Address, error) {
}

func (f *FTSOWrapper) Votepower(provider common.Address) (float64, error) {
}

func (f *FTSOWrapper) Rewards(provider common.Address) (float64, error) {
}

func (f *FTSOWrapper) manager() (common.Address, error) {

	var address common.Address
	err := f.call.OnContract(f.contracts.PriceSubmitter).Execute(f.methods.ManagerAddress).Decode(&address)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not execute manager address call: %w", err)
	}

	return address, nil
}

func (f *FTSOWrapper) rewardSeconds() (uint64, error) {

	manager, err := f.manager()
	if err != nil {
		return 0, fmt.Errorf("could not get manager address: %w", err)
	}

	values, err := f.call(manager, f.abis.Manager, f.methods.EpochSeconds)
	if err != nil {
		return 0, fmt.Errorf("could not get reward seconds: %w", err)
	}

	seconds := values[0].(*big.Int).Uint64()

	return seconds, nil
}

func (f *FTSOWrapper) RewardEpochs(epoch uint64) (rewardEpochs, error) {

	entry, ok := f.epochs.Get(epoch)
	if ok {
		return entry.(rewardEpochs), nil
	}

	hash := f.blockchain.CurrentHeader().Hash()

	values, err := f.call(hash, f.manager, ftsoManagerABI, "rewardEpochs", big.NewInt(0).SetUint64(epoch))
	if err != nil {
		return rewardEpochs{}, fmt.Errorf("could not get reward epochs: %w", err)
	}
	if len(values) != 3 {
		return rewardEpochs{}, fmt.Errorf("wrong number of return values (have: %d, want: %d)", len(values), 3)
	}

	powerBlock := values[0].(*big.Int).Uint64()
	startBlock := values[0].(*big.Int).Uint64()
	startTime := values[2].(*big.Int).Uint64()

	info := rewardEpochs{
		powerBlock: powerBlock,
		startBlock: startBlock,
		startTime:  startTime,
		endTime:    startTime + f.seconds,
	}

	f.epochs.Add(epoch, info)

	return info, nil
}

func (f *FTSOWrapper) GetSupportedIndices(hash common.Hash) ([]*big.Int, error) {

	values, err := f.call(hash, f.registry, ftsoRegistryABI, "getSupportedIndices")
	if err != nil {
		return nil, fmt.Errorf("could not get FTSOWrapper indices: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of return values (have %d, want: %d)", len(values), 1)
	}

	indices := values[0].([]*big.Int)

	return indices, nil
}

func (f *FTSOWrapper) GetFTSOWhitelistedPriceProviders(hash common.Hash, index *big.Int) ([]common.Address, error) {

	values, err := f.call(hash, f.whitelister, voterWhitelisterABI, "getFtsoWhitelistedPriceProviders", index)
	if err != nil {
		return nil, fmt.Errorf("could not get whitelisted price providers: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of return values (have %d, want: %d)", len(values), 1)
	}

	addresses := values[0].([]common.Address)

	return addresses, nil
}

func (f *FTSOWrapper) GetNodeIDForDataProvider(hash common.Hash, provider common.Address) (ids.ShortID, error) {

	values, err := f.call(hash, validatorRegistryAddress, validatorRegistryABI, "getNodeIdForDataProvider", provider)
	if err != nil {
		return ids.ShortEmpty, fmt.Errorf("could not get node ID for data provider: %w", err)
	}
	if len(values) != 1 {
		return ids.ShortEmpty, fmt.Errorf("wrong number of return values (have %d, want: %d)", len(values), 1)
	}

	id := values[0].([20]byte)

	return ids.ShortID(id), nil
}

func (f *FTSOWrapper) VotePowerOf(hash common.Hash, provider common.Address) (float64, error) {

	values, err := f.call(hash, f.votepower, readVotePowerContractABI, "votePowerOf", provider)
	if err != nil {
		return 0, fmt.Errorf("could not get node ID for data provider: %w", err)
	}
	if len(values) != 1 {
		return 0, fmt.Errorf("wrong number of return values (have %d, want: %d)", len(values), 1)
	}

	value := values[0].(*big.Int)
	votepower, _ := big.NewFloat(0).SetInt(value).Float64()

	return votepower, nil
}

func (f *FTSOWrapper) GetUnclaimedReward(hash common.Hash, epoch uint64, provider common.Address) (float64, error) {

	values, err := f.call(hash, f.rewards, ftsoRewardManagerABI, "getUnclaimedReward", big.NewInt(0).SetUint64(epoch), provider)
	if err != nil {
		return 0, fmt.Errorf("could not get unclaimed rewards: %w", err)
	}
	if len(values) != 2 {
		return 0, fmt.Errorf("wrong number of return values (have %d, want: %d)", len(values), 2)
	}

	value := values[0].(*big.Int)
	reward, _ := big.NewFloat(0).SetInt(value).Float64()

	return reward, nil
}

func (f *FTSOWrapper) call(to common.Address, abi abi.ABI) FTSOCall {

	data, err := abi.Pack(method, params...)
	if err != nil {
		return nil, fmt.Errorf("could not pack call data: %w", err)
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &to, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return nil, fmt.Errorf("could not convert arguments to message: %w", err)
	}

	vmConfig := *f.blockchain.GetVMConfig()
	vmConfig.NoBaseFee = true
	chainConfig := f.blockchain.Config()
	txContext := core.NewEVMTxContext(msg)
	blkContext := core.NewEVMBlockContext(header, f.blockchain, nil)
	evm := vm.NewEVM(blkContext, txContext, f.state, chainConfig, vm.Config{NoBaseFee: true})
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

	if len(values) != num {
		return nil, fmt.Errorf("wrong number of return values (want: %d, have: %d)", num, len(values))
	}

	return values, nil
}
