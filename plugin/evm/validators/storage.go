package validators

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/flare/database"

	"github.com/flare-foundation/coreth/core/vm"
)

var (
	errNoPriceSubmitter = errors.New("no price submitter")
	errFTSONotDeployed  = errors.New("FTSO not deployed")
	errFTSONotActive    = errors.New("FTSO not active")
)

type Storage struct {
	db database.Database
}

func NewStorage(db database.Database) *Storage {

	s := Storage{
		db: db,
	}

	return &s
}

func (s *Storage) WithEVM(evm *vm.EVM) (vm.ValidatorManager, error) {

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

	submitter := evmContract{
		address: params.SubmitterAddress,
		abi:     abiSubmitter,
	}

	var managerAddress common.Address
	err = newContractCall(evm, submitter).execute(ManagerAddress).decode(&managerAddress)
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
		abi:     abiManager,
	}

	height := &big.Int{}
	err = newContractCall(evm, manager).execute(RewardEpoch, big.NewInt(0)).decode(nil, &height, nil)
	if errors.Is(err, vm.ErrExecutionReverted) || height.Uint64() == 0 {
		return nil, errFTSONotActive
	}
	if err != nil {
		return nil, fmt.Errorf("could not get first epoch: %w", err)
	}

	var rewardsAddress common.Address
	err = newContractCall(evm, manager).execute(RewardsAddress).decode(&rewardsAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get rewards address: %w", err)
	}

	rewards := evmContract{
		address: rewardsAddress,
		abi:     abiRewards,
	}

	var registryAddress common.Address
	err = newContractCall(evm, submitter).execute(RegistryAddress).decode(&registryAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get registry address: %w", err)
	}

	registry := evmContract{
		address: registryAddress,
		abi:     abiRegistry,
	}

	var whitelistAddress common.Address
	err = newContractCall(evm, submitter).execute(WhitelistAddress).decode(&whitelistAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get whitelist address: %w", err)
	}

	whitelist := evmContract{
		address: whitelistAddress,
		abi:     abiWhitelist,
	}

	var wnatAddress common.Address
	err = newContractCall(evm, rewards).execute(WNATAddress).decode(&wnatAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get WNAT address: %w", err)
	}

	wnat := evmContract{
		address: wnatAddress,
		abi:     abiWNAT,
	}

	var votepowerAddress common.Address
	err = newContractCall(evm, wnat).execute(VotepowerAddress).decode(&votepowerAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get votepower address: %w", err)
	}

	votepower := evmContract{
		address: votepowerAddress,
		abi:     abiVotepower,
	}

	contracts := Contracts{
		Registry:  registry,
		Manager:   manager,
		Rewards:   rewards,
		Whitelist: whitelist,
		WNAT:      wnat,
		Votepower: votepower,
	}

	m := Manager{
		db:        s.db,
		evm:       evm,
		contracts: contracts,
	}

	return &m, nil
}
