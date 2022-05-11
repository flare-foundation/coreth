package ftso

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/flare/ids"
)

var (
	errNoPriceSubmitter    = errors.New("no price submitter")
	errFTSONotDeployed     = errors.New("FTSO not deployed")
	errFTSONotActive       = errors.New("FTSO not active")
	errRegistryNotDeployed = errors.New("validators registry not deployed")
)

type Storer interface {
	// Get retrieves the given key if it's present in the key-value data store.
	Get(key []byte) ([]byte, error)

	// Put inserts the given value into the key-value data store.
	Put(key []byte, value []byte) error
}

type Validator struct {
	evm       *vm.EVM
	contracts FTSOContracts
}

type FTSOContracts struct {
	Registry   EVMContract
	Manager    EVMContract
	Rewards    EVMContract
	Whitelist  EVMContract
	WNAT       EVMContract
	Votepower  EVMContract
	Validation EVMContract
}

var ValidatorDB Storer

func NewValidator(evm *vm.EVM) (*Validator, error) {
	abiRegistry, err := abi.JSON(strings.NewReader(jsonRegistry))
	if err != nil {
		return nil, fmt.Errorf("could not parse registry ABI: %w", err)
	}

	abiManager, err := abi.JSON(strings.NewReader(jsonManager))
	if err != nil {
		return nil, fmt.Errorf("could not parse manager ABI: %w", err)
	}

	abiRewards, err := abi.JSON(strings.NewReader(jsonRewards))
	if err != nil {
		return nil, fmt.Errorf("could not parse rewards ABI: %w", err)
	}

	abiWhitelist, err := abi.JSON(strings.NewReader(jsonWhitelist))
	if err != nil {
		return nil, fmt.Errorf("could not parse whitelist ABI: %w", err)
	}

	abiWNAT, err := abi.JSON(strings.NewReader(jsonWNAT))
	if err != nil {
		return nil, fmt.Errorf("could not parse WNAT ABI: %w", err)
	}

	abiVotepower, err := abi.JSON(strings.NewReader(jsonVotepower))
	if err != nil {
		return nil, fmt.Errorf("could not parse votepower ABI: %w", err)
	}

	abiSubmitter, err := abi.JSON(strings.NewReader(jsonSubmitter))
	if err != nil {
		return nil, fmt.Errorf("could not parse submitter ABI: %w", err)
	}

	submitter := EVMContract{
		address: params.SubmitterAddress,
		abi:     abiSubmitter,
	}

	var managerAddress common.Address
	err = NewContractCall(evm, submitter).Execute(ManagerAddress).Decode(&managerAddress)
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

	manager := EVMContract{
		address: managerAddress,
		abi:     abiManager,
	}

	height := &big.Int{}
	err = NewContractCall(evm, manager).Execute(RewardEpoch, big.NewInt(0)).Decode(nil, &height, nil)
	if errors.Is(err, vm.ErrExecutionReverted) || height.Uint64() == 0 {
		return nil, errFTSONotActive
	}
	if err != nil {
		return nil, fmt.Errorf("could not get first epoch: %w", err)
	}

	var rewardsAddress common.Address
	err = NewContractCall(evm, manager).Execute(RewardsAddress).Decode(&rewardsAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get rewards address: %w", err)
	}

	rewards := EVMContract{
		address: rewardsAddress,
		abi:     abiRewards,
	}

	var registryAddress common.Address
	err = NewContractCall(evm, submitter).Execute(RegistryAddress).Decode(&registryAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get registry address: %w", err)
	}

	registry := EVMContract{
		address: registryAddress,
		abi:     abiRegistry,
	}

	var whitelistAddress common.Address
	err = NewContractCall(evm, submitter).Execute(WhitelistAddress).Decode(&whitelistAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get whitelist address: %w", err)
	}

	whitelist := EVMContract{
		address: whitelistAddress,
		abi:     abiWhitelist,
	}

	var wnatAddress common.Address
	err = NewContractCall(evm, rewards).Execute(WNATAddress).Decode(&wnatAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get WNAT address: %w", err)
	}

	wnat := EVMContract{
		address: wnatAddress,
		abi:     abiWNAT,
	}

	var votepowerAddress common.Address
	err = NewContractCall(evm, wnat).Execute(VotepowerAddress).Decode(&votepowerAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get votepower address: %w", err)
	}

	votepower := EVMContract{
		address: votepowerAddress,
		abi:     abiVotepower,
	}

	contracts := FTSOContracts{
		Registry:  registry,
		Manager:   manager,
		Rewards:   rewards,
		Whitelist: whitelist,
		WNAT:      wnat,
		Votepower: votepower,
		Validation: EVMContract{
			address: params.ValidationAddress,
			abi:     abiManager,
		},
	}

	return &Validator{
		evm:       evm,
		contracts: contracts,
	}, nil
}

func (v *Validator) LastRewardEpoch() (uint64, error) {
	ValidatorDB.Get(nil)

	return 0, nil
}

func (v *Validator) UpdateLastRewardEpoch(epoch uint64) error {
	ValidatorDB.Put(nil, nil)

	return nil
}

func (v *Validator) Cap() (float64, error) {

	supply := &big.Int{}
	err := NewContractCall(v.evm, v.contracts.WNAT).Execute(TotalSupply).Decode(&supply)
	if err != nil {
		return 0, fmt.Errorf("could not get total supply: %w", err)
	}

	fraction := &big.Int{}
	err = NewContractCall(v.evm, v.contracts.Manager).
		Execute(Settings).
		Decode(&fraction, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get votepower threshold fraction: %w", err)
	}

	capInt := big.NewInt(0).Div(supply, fraction)
	capFloat := big.NewFloat(0).SetInt(capInt)
	cap, _ := capFloat.Float64()

	return cap, nil
}

func (v *Validator) Validator(provider common.Address) (ids.ShortID, error) {
	// TODO this is precompiled now
	// TODO either move this to precompiled contract with static list or put the static list here

	return [20]byte{}, nil
}

func (v *Validator) Votepower(provider common.Address, epoch uint64) (float64, error) {
	powerHeight := &big.Int{}
	startHeight := &big.Int{}
	startTime := &big.Int{}

	err := NewContractCall(v.evm, v.contracts.Manager).
		Execute(RewardEpoch, big.NewInt(0).SetUint64(epoch)).
		Decode(&powerHeight, &startHeight, &startTime)
	if err != nil {
		return 0, fmt.Errorf("could not execute epoch info retrieval: %w", err)
	}

	vpInt := &big.Int{}
	err = NewContractCall(v.evm, v.contracts.Votepower).
		Execute(ProviderVotepower, provider).
		Decode(&vpInt)
	if err != nil {
		return 0, fmt.Errorf("could not get provider votepower: %w", err)
	}

	vpFloat := big.NewFloat(0).SetInt(vpInt)
	votepower, _ := vpFloat.Float64()

	return votepower, nil
}

func (v *Validator) Rewards(provider common.Address, epoch uint64) (float64, error) {

	rwInt := &big.Int{}
	err := NewContractCall(v.evm, v.contracts.Rewards).
		Execute(ProviderRewards, big.NewInt(0).SetUint64(epoch), provider).
		Decode(&rwInt, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get provider rewards: %w", err)
	}

	rwFloat := big.NewFloat(0).SetInt(rwInt)
	rewards, _ := rwFloat.Float64()

	return rewards, nil
}
