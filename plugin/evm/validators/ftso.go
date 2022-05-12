package validators

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
)

type FTSO struct {
	evm       *vm.EVM
	contracts Contracts
}

func NewFTSO(evm *vm.EVM) (*FTSO, error) {

	submitter := evmContract{
		address: params.SubmitterAddress,
		abi:     abis.submitter,
	}

	var managerAddress common.Address
	err := newContractCall(evm, submitter).execute(getAddressManager).decode(&managerAddress)
	if errors.Is(err, errNoReturnData) {
		return nil, errNoPriceSubmitter
	}
	if err != nil {
		return nil, fmt.Errorf("could not get manager address: %w", err)
	}

	empty := common.Address{}
	if managerAddress == empty {
		return nil, errFTSONotDeployed
	}

	manager := evmContract{
		address: managerAddress,
		abi:     abis.manager,
	}

	height := &big.Int{}
	err = newContractCall(evm, manager).execute(getEpochInfo, big.NewInt(0)).decode(nil, &height, nil)
	if errors.Is(err, vm.ErrExecutionReverted) || height.Uint64() == 0 {
		return nil, errFTSONotActive
	}
	if err != nil {
		return nil, fmt.Errorf("could not get first epoch info: %w", err)
	}

	var rewardsAddress common.Address
	err = newContractCall(evm, manager).execute(getAddressRewards).decode(&rewardsAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get rewards address: %w", err)
	}

	rewards := evmContract{
		address: rewardsAddress,
		abi:     abis.rewards,
	}

	var registryAddress common.Address
	err = newContractCall(evm, submitter).execute(getAddressRegistry).decode(&registryAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get registry address: %w", err)
	}

	registry := evmContract{
		address: registryAddress,
		abi:     abis.registry,
	}

	var whitelistAddress common.Address
	err = newContractCall(evm, submitter).execute(getAddressWhitelist).decode(&whitelistAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get whitelist address: %w", err)
	}

	whitelist := evmContract{
		address: whitelistAddress,
		abi:     abis.whitelist,
	}

	var wnatAddress common.Address
	err = newContractCall(evm, rewards).execute(getAddressWNAT).decode(&wnatAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get WNAT address: %w", err)
	}

	wnat := evmContract{
		address: wnatAddress,
		abi:     abis.wnat,
	}

	var votepowerAddress common.Address
	err = newContractCall(evm, wnat).execute(getAddressVotepower).decode(&votepowerAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get votepower address: %w", err)
	}

	votepower := evmContract{
		address: votepowerAddress,
		abi:     abis.votepower,
	}

	contracts := Contracts{
		Registry:  registry,
		Manager:   manager,
		Rewards:   rewards,
		Whitelist: whitelist,
		WNAT:      wnat,
		Votepower: votepower,
	}

	f := FTSO{
		evm:       evm,
		contracts: contracts,
	}

	return &f, nil
}

func (s *FTSO) Supply() (float64, error) {

	var supplyInt big.Int
	err := newContractCall(s.evm, s.contracts.WNAT).
		execute(getFTSOSupply).
		decode(&supplyInt)
	if err != nil {
		return 0, fmt.Errorf("could not get total supply: %w", err)
	}

	supplyFloat := big.NewFloat(0).SetInt(&supplyInt)
	supply, _ := supplyFloat.Float64()

	return supply, nil

}

func (s *FTSO) Fraction() (uint64, error) {

	var fraction big.Int
	err := newContractCall(s.evm, s.contracts.Manager).
		execute(getFTSOSettings).
		decode(&fraction, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get votepower threshold fraction: %w", err)
	}

	return fraction.Uint64(), nil
}

func (s *FTSO) Votepower(provider common.Address, epoch uint64) (float64, error) {

	var votepowerInt big.Int
	err := newContractCall(s.evm, s.contracts.Votepower).
		execute(getProviderVotepower, provider).
		decode(&votepowerInt)
	if err != nil {
		return 0, fmt.Errorf("could not get provider votepower: %w", err)
	}

	votepowerFloat := big.NewFloat(0).SetInt(&votepowerInt)
	votepower, _ := votepowerFloat.Float64()

	return votepower, nil
}

func (s *FTSO) Rewards(provider common.Address, epoch uint64) (float64, error) {

	var rewardsInt big.Int
	err := newContractCall(s.evm, s.contracts.Rewards).
		execute(getProviderRewards, big.NewInt(0).SetUint64(epoch), provider).
		decode(&rewardsInt, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get provider rewards: %w", err)
	}

	rewardsFloat := big.NewFloat(0).SetInt(&rewardsInt)
	rewards, _ := rewardsFloat.Float64()

	return rewards, nil
}

func (s *FTSO) Providers() ([]common.Address, error) {

	var indices []*big.Int
	err := newContractCall(s.evm, s.contracts.Registry).
		execute(getFTSOIndices).
		decode(&indices)
	if err != nil {
		return nil, fmt.Errorf("could not get series indices: %w", err)
	}

	providerMap := make(map[common.Address]struct{})
	for _, index := range indices {
		var addresses []common.Address
		err := newContractCall(s.evm, s.contracts.Whitelist).
			execute(getFTSOProviders, index).
			decode(&addresses)
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

func (s *FTSO) Current() (uint64, error) {

	var epoch big.Int
	err := newContractCall(s.evm, s.contracts.Rewards).
		execute(getEpochCurrent).
		decode(&epoch)
	if err != nil {
		return 0, fmt.Errorf("could not execute current epoch retrieval: %w", err)
	}

	return epoch.Uint64(), nil
}
