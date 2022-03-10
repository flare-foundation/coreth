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
	blockchain  *core.BlockChain
	last        uint64
	seconds     uint64
	epochs      *lru.ARCCache
	manager     common.Address
	registry    common.Address
	whitelister common.Address
	votepower   common.Address
}

func NewFTSO(blockchain *core.BlockChain) (*FTSO, error) {

	epochs, _ := lru.NewARC(epochsCacheSize)
	f := FTSO{
		blockchain: blockchain,
		seconds:    0,
		last:       0,
		epochs:     epochs,
	}

	manager, err := f.FTSOManager()
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO manager address: %w", err)
	}

	f.manager = manager

	registry, err := f.FTSORegistry()
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO registry: %w", err)
	}

	f.registry = registry

	whitelister, err := f.VoterWhitelister()
	if err != nil {
		return nil, fmt.Errorf("could not get voter whitelister: %w", err)
	}

	f.whitelister = whitelister

	votepower, err := f.ReadVotePowerContract()
	if err != nil {
		return nil, fmt.Errorf("could not get read vote power contract address: %w", err)
	}

	f.votepower = votepower

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

	info, err := f.RewardEpochs(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get rewards epoch: %w", err)
	}

	header := f.blockchain.GetHeaderByNumber(info.block)
	if header == nil {
		return nil, fmt.Errorf("unknown header (number: %d)", info.block)
	}

	hash := header.Hash()
	indices, err := f.GetSupportedIndices(hash)
	if err != nil {
		return nil, fmt.Errorf("could not get FTSOs: %w", err)
	}

	providerSet := make(map[common.Address]struct{})
	for _, index := range indices {
		addresses, err := f.GetFTSOWhitelistedPriceProviders(hash, index)
		if err != nil {
			return nil, fmt.Errorf("could not get whitelisted price providers (index: %d): %w", index, err)
		}
		for _, address := range addresses {
			providerSet[address] = struct{}{}
		}
	}

	providers := make([]common.Address, 0, len(providerSet))
	for provider := range providerSet {
		providers = append(providers, provider)
	}

	return providers, nil
}

func (f *FTSO) ValidatorForProviderAtEpoch(epoch uint64, provider common.Address) (ids.ShortID, error) {

	// TODO: call get vote power block function directly
	info, err := f.RewardEpochs(epoch)
	if err != nil {
		return ids.ShortEmpty, fmt.Errorf("could not get rewards epoch: %w", err)
	}

	header := f.blockchain.GetHeaderByNumber(info.block)
	if header == nil {
		return ids.ShortEmpty, fmt.Errorf("unknown header (number: %d)", info.block)
	}

	hash := header.Hash()
	validator, err := f.GetNodIDForDataProvider(hash, provider)
	if err != nil {
		return ids.ShortEmpty, fmt.Errorf("could not get validator for provider: %w", err)
	}

	return validator, nil
}

func (f *FTSO) VotepowerForProviderAtEpoch(epoch uint64, provider common.Address) (float64, error) {

	info, err := f.RewardEpochs(epoch)
	if err != nil {
		return 0, fmt.Errorf("could not get rewards epoch: %w", err)
	}

	header := f.blockchain.GetHeaderByNumber(info.block)
	if header == nil {
		return 0, fmt.Errorf("unknown header (number: %d)", info.block)
	}

	hash := header.Hash()
	votepower, err := f.VotePowerOf(hash, provider)
	if err != nil {
		return 0, fmt.Errorf("could not get vote power for provider: %w", err)
	}

	return votepower, nil
}

func (f *FTSO) RewardForProviderAtEpoch(epoch uint64, provider common.Address) (float64, error) {
	return 0, nil
}

func (f *FTSO) FTSOManager() (common.Address, error) {
	return common.Address{}, nil
}

func (f *FTSO) FTSORegistry() (common.Address, error) {
	return common.Address{}, nil
}

func (f *FTSO) VoterWhitelister() (common.Address, error) {
	return common.Address{}, nil
}

func (f *FTSO) ReadVotePowerContract() (common.Address, error) {
	return common.Address{}, nil
}

func (f *FTSO) RewardEpochDurationSeconds() (uint64, error) {

	hash := f.blockchain.CurrentHeader().Hash()

	values, err := f.call(hash, f.manager, ftsoManagerABI, "rewardEpochDurationSeconds")
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
	block uint64
	start uint64
	end   uint64
}

func (f *FTSO) RewardEpochs(epoch uint64) (rewardEpochs, error) {

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

	block := values[0].(*big.Int).Uint64()
	start := values[2].(*big.Int).Uint64()

	info := rewardEpochs{
		block: block,
		start: start,
		end:   start + f.seconds,
	}

	f.epochs.Add(epoch, info)

	return info, nil
}

func (f *FTSO) GetSupportedIndices(hash common.Hash) ([]*big.Int, error) {

	values, err := f.call(hash, f.registry, ftsoRegistryABI, "getSupportedIndices")
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO indices: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of return values (have %d, want: %d)", len(values), 1)
	}

	indices := values[0].([]*big.Int)

	return indices, nil
}

func (f *FTSO) GetFTSOWhitelistedPriceProviders(hash common.Hash, index *big.Int) ([]common.Address, error) {

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

func (f *FTSO) GetNodIDForDataProvider(hash common.Hash, provider common.Address) (ids.ShortID, error) {

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

func (f *FTSO) VotePowerOf(hash common.Hash, provider common.Address) (float64, error) {

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

func (f *FTSO) call(hash common.Hash, to common.Address, abi abi.ABI, method string, params ...interface{}) ([]interface{}, error) {

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
	args := ethapi.TransactionArgs{To: &to, Input: &input}
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
