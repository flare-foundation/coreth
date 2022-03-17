// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
)

type FTSOSystem struct {
	blockchain *core.BlockChain
	submitter  EVMContract
	validation EVMContract
	abis       FTSOABIs
}

type FTSOContracts struct {
	Submitter  EVMContract
	Registry   EVMContract
	Series     []EVMContract
	Manager    EVMContract
	Rewards    EVMContract
	Whitelist  EVMContract
	Votepower  EVMContract
	WNAT       EVMContract
	Validation EVMContract
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

	abiSeries, err := abi.JSON(strings.NewReader(jsonSeries))
	if err != nil {
		return nil, fmt.Errorf("could not parse series ABI: %w", err)
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
		Submitter:  abiSubmitter,
		Registry:   abiRegistry,
		Manager:    abiManager,
		Series:     abiSeries,
		Rewards:    abiRewards,
		WNAT:       abiWNAT,
		Whitelist:  abiWhitelist,
		Votepower:  abiVotepower,
		Validation: abiValidation,
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

	var registryAddress common.Address
	err := snap.OnContract(f.submitter).Execute(RegistryAddress).Decode(&registryAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get registry address: %w", err)
	}

	registry := EVMContract{
		address: registryAddress,
		abi:     f.abis.Registry,
	}

	var addressesSeries []common.Address
	err = snap.OnContract(registry).Execute(SeriesAddresses).Decode(&addressesSeries)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get series addresses: %w", err)
	}

	series := make([]EVMContract, 0, len(addressesSeries))
	for _, addressSeries := range addressesSeries {
		serie := EVMContract{
			address: addressSeries,
			abi:     f.abis.Series,
		}
		series = append(series, serie)
	}

	var managerAddress common.Address
	err = snap.OnContract(f.submitter).Execute(ManagerAddress).Decode(&managerAddress)
	if err != nil {
		return FTSOContracts{}, fmt.Errorf("could not get manager address: %w", err)
	}

	manager := EVMContract{
		address: managerAddress,
		abi:     f.abis.Manager,
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
		Submitter:  f.submitter,
		Registry:   registry,
		Series:     series,
		Manager:    manager,
		Rewards:    rewards,
		Whitelist:  whitelist,
		WNAT:       wnat,
		Votepower:  votepower,
		Validation: f.validation,
	}

	return contracts, nil
}

func (f *FTSOSystem) EpochInfo(epoch uint64) (EpochInfo, error) {

	header := f.blockchain.CurrentHeader()
	if header == nil {
		return EpochInfo{}, fmt.Errorf("no current header")
	}

	hash := header.Hash()
	contracts, err := f.Contracts(hash)
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get contracts (hash: %x): %w", hash, err)
	}

	call := BindEVM(f.blockchain).AtBlock(hash).OnContract(contracts.Manager)

	var seconds big.Int
	err = call.Execute(EpochSeconds).Decode(&seconds)
	if err != nil {
		return EpochInfo{}, fmt.Errorf("could not get epoch seconds: %w", err)
	}

	var startHeight, startTime *big.Int
	err = call.Execute(RewardEpoch).Decode(nil, &startHeight, &startTime)
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

func (f *FTSOSystem) Snapshot(epoch uint64) (*FTSOSnapshot, error) {

	info, err := f.EpochInfo(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get epoch info: %w", err)
	}

	header := f.blockchain.GetHeaderByNumber(info.StartHeight)
	if header == nil {
		return nil, fmt.Errorf("unknown block (height: %d)", info.StartHeight)
	}

	hash := header.Hash()
	contracts, err := f.Contracts(hash)
	if err != nil {
		return nil, fmt.Errorf("could not get contracts (hash: %x): %w", hash, err)
	}

	snap := FTSOSnapshot{
		system:    f,
		hash:      hash,
		contracts: contracts,
	}

	return &snap, nil
}
