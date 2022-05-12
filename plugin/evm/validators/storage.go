package validators

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/database"

	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
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

	m := Manager{
		db:        s.db,
		evm:       evm,
		contracts: contracts,
	}

	return &m, nil
}
