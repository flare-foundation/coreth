// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/vm"
)

type FTSOSystem struct {
	blockchain *core.BlockChain
	submitter  EVMContract
	validation EVMContract
	abis       FTSOABIs
}

type FTSOABIs struct {
	Registry  abi.ABI
	Manager   abi.ABI
	Rewards   abi.ABI
	Whitelist abi.ABI
	WNAT      abi.ABI
	Votepower abi.ABI
}

type FTSOContracts struct {
	Registry   EVMContract
	Manager    EVMContract
	Rewards    EVMContract
	Whitelist  EVMContract
	Votepower  EVMContract
	Validation EVMContract
}

type FTSOEpoch struct {
	PowerHeight uint64
	StartHeight uint64
	StartTime   uint64
}

func NewFTSOSystem(blockchain *core.BlockChain, addressSubmitter common.Address, addressValidation common.Address) (*FTSOSystem, error) {

	abiSubmitter, err := abi.JSON(strings.NewReader(jsonSubmitter))
	if err != nil {
		return nil, fmt.Errorf("could not parse submitter ABI: %w", err)
	}

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

	abiValidation, err := abi.JSON(strings.NewReader(jsonValidation))
	if err != nil {
		return nil, fmt.Errorf("could not get validation ABI: %w", err)
	}

	submitter := EVMContract{
		address: addressSubmitter,
		abi:     abiSubmitter,
	}

	validation := EVMContract{
		address: addressValidation,
		abi:     abiValidation,
	}

	abis := FTSOABIs{
		Registry:  abiRegistry,
		Manager:   abiManager,
		Rewards:   abiRewards,
		WNAT:      abiWNAT,
		Whitelist: abiWhitelist,
		Votepower: abiVotepower,
	}

	f := FTSOSystem{
		blockchain: blockchain,
		submitter:  submitter,
		validation: validation,
		abis:       abis,
	}

	return &f, nil
}

func (f *FTSOSystem) Contracts(hash common.Hash) (FTSOContracts, error) {

	snap := BindEVM(f.blockchain).AtBlock(hash)

	var managerAddress common.Address
	err := snap.OnContract(f.submitter).Execute(ManagerAddress).Decode(&managerAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get manager address: %w", err)
	}

	empty := common.Address{}
	if managerAddress == empty {
		return FTSOContracts{}, errFTSONotDeployed
	}

	manager := EVMContract{
		address: managerAddress,
		abi:     f.abis.Manager,
	}

	height := &big.Int{}
	err = snap.OnContract(manager).Execute(RewardEpoch, big.NewInt(0)).Decode(nil, &height, nil)
	if errors.Is(err, vm.ErrExecutionReverted) || height.Uint64() == 0 {
		return FTSOContracts{}, errFTSONotActive
	}
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get first epoch: %w", err)
	}

	var rewardsAddress common.Address
	err = snap.OnContract(manager).Execute(RewardsAddress).Decode(&rewardsAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get rewards address: %w", err)
	}

	rewards := EVMContract{
		address: rewardsAddress,
		abi:     f.abis.Rewards,
	}

	var registryAddress common.Address
	err = snap.OnContract(f.submitter).Execute(RegistryAddress).Decode(&registryAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get registry address: %w", err)
	}

	registry := EVMContract{
		address: registryAddress,
		abi:     f.abis.Registry,
	}

	var whitelistAddress common.Address
	err = snap.OnContract(f.submitter).Execute(WhitelistAddress).Decode(&whitelistAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get whitelist address: %w", err)
	}

	whitelist := EVMContract{
		address: whitelistAddress,
		abi:     f.abis.Whitelist,
	}

	var wnatAddress common.Address
	err = snap.OnContract(rewards).Execute(WNATAddress).Decode(&wnatAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get WNAT address: %w", err)
	}

	wnat := EVMContract{
		address: wnatAddress,
		abi:     f.abis.WNAT,
	}

	var votepowerAddress common.Address
	err = snap.OnContract(wnat).Execute(VotepowerAddress).Decode(&votepowerAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get votepower address: %w", err)
	}

	votepower := EVMContract{
		address: votepowerAddress,
		abi:     f.abis.Votepower,
	}

	contracts := FTSOContracts{
		Registry:   registry,
		Manager:    manager,
		Rewards:    rewards,
		Whitelist:  whitelist,
		Votepower:  votepower,
		Validation: f.validation,
	}

	return contracts, nil
}

func (f *FTSOSystem) Current(hash common.Hash) (uint64, error) {

	contracts, err := f.Contracts(hash)
	if err != nil {
		return 0, fmt.Errorf("could not get contracts: %w", err)
	}

	epoch := &big.Int{}
	err = BindEVM(f.blockchain).
		AtBlock(hash).
		OnContract(contracts.Manager).
		Execute(CurrentEpoch).
		Decode(&epoch)
	if err != nil {
		return 0, fmt.Errorf("could not execute current epoch retrieval: %w", err)
	}

	return epoch.Uint64(), nil
}

func (f *FTSOSystem) Details(epoch uint64) (FTSOEpoch, error) {

	header := f.blockchain.CurrentHeader()
	if header == nil {
		return FTSOEpoch{}, fmt.Errorf("no current header")
	}

	hash := header.Hash()
	contracts, err := f.Contracts(hash)
	if err != nil {
		return FTSOEpoch{}, fmt.Errorf("could not get contracts (hash: %x): %w", hash, err)
	}

	call := BindEVM(f.blockchain).AtBlock(hash).OnContract(contracts.Manager)

	powerHeight := &big.Int{}
	startHeight := &big.Int{}
	startTime := &big.Int{}
	err = call.
		Execute(RewardEpoch, big.NewInt(0).SetUint64(epoch)).
		Decode(&powerHeight, &startHeight, &startTime)
	if err != nil {
		return FTSOEpoch{}, fmt.Errorf("could not execute epoch info retrieval (hash: %x): %w", hash, err)
	}

	info := FTSOEpoch{
		PowerHeight: powerHeight.Uint64(),
		StartHeight: startHeight.Uint64(),
		StartTime:   startTime.Uint64(),
	}

	return info, nil
}

func (f *FTSOSystem) Snapshot(epoch uint64) (Snapshot, error) {

	currentEpoch, err := f.Details(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get current epoch details: %w", err)
	}

	powerHeader := f.blockchain.GetHeaderByNumber(currentEpoch.PowerHeight)
	if powerHeader == nil {
		return nil, fmt.Errorf("unknown power block (height: %d)", currentEpoch.PowerHeight)
	}

	startHeader := f.blockchain.GetHeaderByNumber(currentEpoch.StartHeight)
	if startHeader == nil {
		return nil, fmt.Errorf("unknown current block (height: %d)", currentEpoch.StartHeight)
	}

	nextEpoch, err := f.Details(epoch + 1)
	if err != nil {
		return nil, fmt.Errorf("could not get next epoch details: %w", err)
	}

	endHeader := f.blockchain.GetHeaderByNumber(nextEpoch.StartHeight)
	if endHeader == nil {
		return nil, fmt.Errorf("unknown next block (height: %d)", nextEpoch.StartHeight)
	}

	startHash := startHeader.Hash()
	contracts, err := f.Contracts(startHash)
	if err != nil {
		return nil, fmt.Errorf("could not get contracts (hash: %x): %w", startHash, err)
	}

	snap := FTSOSnapshot{
		system:    f,
		epoch:     epoch,
		power:     powerHeader.Hash(),
		start:     startHash,
		end:       endHeader.Hash(),
		contracts: contracts,
	}

	return &snap, nil
}
