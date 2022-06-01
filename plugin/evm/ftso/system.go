package ftso

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
)

// System represents a convenience wrapper around the EVM state that allows us to
// easily interact with the FTSO smart contracts.
type System struct {
	evm       *vm.EVM
	contracts Contracts
}

// Contracts is a convenience wrapper that groups together the smart contract addresses
// and ABIs for all FTSO smart contracts.
type Contracts struct {
	Registry  evmContract
	Manager   evmContract
	Rewards   evmContract
	Whitelist evmContract
	WNAT      evmContract
	Votepower evmContract
}

// NewSystem binds our convenience wrapper around FTSO smart contracts to the given
// EVM state.
func NewSystem(evm *vm.EVM) (*System, error) {

	// The price submitter smart contract is the bootstrapping point of origin. All
	// other FTSO smart contract addresses can be derived from it through various
	// paths.
	submitter := evmContract{
		address: params.SubmitterAddress,
		abi:     abis.submitter,
	}

	// The next relevant smart contract is the FTSO manager. If there is no return
	// value, the price submitter genesis contract has not been deployed, so the
	// FTSO system as a whole is unavailable and non-functional.
	var managerAddress common.Address
	err := newContractCall(evm, submitter).execute(getAddressManager).decode(&managerAddress)
	if errors.Is(err, errNoReturnData) {
		return nil, ErrNoPriceSubmitter
	}
	if err != nil {
		return nil, fmt.Errorf("could not get manager address: %w", err)
	}

	// If the manager address is empty, it means the non-genesis FTSO contracts have
	// not yet been deployed, and the FTSO system as a whole has thus not been deployed.
	empty := common.Address{}
	if managerAddress == empty {
		return nil, ErrFTSONotDeployed
	}

	// At this point, we have the FTSO manager and can initialize its contract wrapper.
	// We can use it to check on the status of the FTSO system deployment.
	manager := evmContract{
		address: managerAddress,
		abi:     abis.manager,
	}

	// Next, we want to get the epoch information from the FTSO system. If we can not
	// get the epoch information for the zero reward epoch, or if its start height is
	// zero, it means the FTSO system has not yet been started / activated.
	height := big.NewInt(0)
	err = newContractCall(evm, manager).execute(getEpochInfo, big.NewInt(0)).decode(nil, &height, nil)
	if errors.Is(err, vm.ErrExecutionReverted) || height.Uint64() == 0 {
		return nil, ErrFTSONotActive
	}
	if err != nil {
		return nil, fmt.Errorf("could not get first epoch info: %w", err)
	}

	// At this point, the FTSO system is deployed and active. We should be able to get
	// the FTSO rewards manager address from the FTSO manager.
	var rewardsAddress common.Address
	err = newContractCall(evm, manager).execute(getAddressRewards).decode(&rewardsAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get rewards address: %w", err)
	}

	// Create the contract wrapper for the FTSO rewards manager. It will allow us to
	// deduce the address of the WNat address responsible for delegations, which in
	// turn allows us to get the contract address that holds votepower. It will also
	// allow us to get the rewards for FTSO data providers, which in turn is part of
	// the validator weight computation.
	rewards := evmContract{
		address: rewardsAddress,
		abi:     abis.rewards,
	}

	// We should be able to get the FTSO registry from the price submitter.
	var registryAddress common.Address
	err = newContractCall(evm, submitter).execute(getAddressRegistry).decode(&registryAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get registry address: %w", err)
	}

	// Create the contract wrapper for the FTSO registry. It will allow us to get the
	// list of active FTSO assets later, as we want to get the whitelist of data
	// providers for each asset.
	registry := evmContract{
		address: registryAddress,
		abi:     abis.registry,
	}

	// We should be able to get the whitelister contract from the price submitter.
	var whitelistAddress common.Address
	err = newContractCall(evm, submitter).execute(getAddressWhitelist).decode(&whitelistAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get whitelist address: %w", err)
	}

	// Create the whitelister contract wrapper. It will allow us to retrieve the list
	// of whitelisted FTSO providers later.
	whitelist := evmContract{
		address: whitelistAddress,
		abi:     abis.whitelist,
	}

	// We should be able to get the address of the WNat address that allows people to
	// wrap native tokens as ERC20 and delegate/stake them from the rewards manager.
	var wnatAddress common.Address
	err = newContractCall(evm, rewards).execute(getAddressWNAT).decode(&wnatAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get WNAT address: %w", err)
	}

	// Create the WNat contract wrapper, which allows us to find out the votepower
	// contract address.
	wnat := evmContract{
		address: wnatAddress,
		abi:     abis.wnat,
	}

	// We should be able to get the votepower contract address from the WNat.
	var votepowerAddress common.Address
	err = newContractCall(evm, wnat).execute(getAddressVotepower).decode(&votepowerAddress)
	if err != nil {
		return nil, fmt.Errorf("could not get votepower address: %w", err)
	}

	// Create the votepower contract wrapper. It will allow us to get the votepower
	// of FTSO data providers to deduce validator weight.
	votepower := evmContract{
		address: votepowerAddress,
		abi:     abis.votepower,
	}

	// Create a simple convenience wrapper around the smart contracts.
	contracts := Contracts{
		Registry:  registry,
		Manager:   manager,
		Rewards:   rewards,
		Whitelist: whitelist,
		WNAT:      wnat,
		Votepower: votepower,
	}

	s := System{
		evm:       evm,
		contracts: contracts,
	}

	return &s, nil
}

// Current returns the active rewards epoch of the FTSO system.
func (s *System) Current() (uint64, error) {

	epoch := big.NewInt(0)
	err := newContractCall(s.evm, s.contracts.Manager).
		execute(getEpochCurrent).
		decode(&epoch)
	if err != nil {
		return 0, fmt.Errorf("could not execute current epoch retrieval: %w", err)
	}

	return epoch.Uint64(), nil
}

// Cap returns the current votepower cap for FTSO data providers.
func (s *System) Cap() (float64, error) {

	supply := big.NewInt(0)
	err := newContractCall(s.evm, s.contracts.WNAT).
		execute(getFTSOSupply).
		decode(&supply)
	if err != nil {
		return 0, fmt.Errorf("could not get total supply: %w", err)
	}

	fraction := big.NewInt(0)
	err = newContractCall(s.evm, s.contracts.Manager).
		execute(getFTSOSettings).
		decode(&fraction, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get votepower threshold fraction: %w", err)
	}

	capInt := big.NewInt(0).Div(supply, fraction)
	capFloat := big.NewFloat(0).SetInt(capInt)
	cap, _ := capFloat.Float64()

	return cap, nil
}

// Whitelist returns the complete list of addresses for all FTSO data providers.
func (s *System) Whitelist() ([]common.Address, error) {

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

// Votepower returns the current votepower for the FTSO data provider with the given
// address.
func (s *System) Votepower(provider common.Address) (float64, error) {

	current, err := s.Current()
	if err != nil {
		return 0, fmt.Errorf("could not get current epoch: %w", err)
	}
	epoch := big.NewInt(0).SetUint64(current - 1)

	height := big.NewInt(0)
	err = newContractCall(s.evm, s.contracts.Manager).
		execute(getEpochInfo, epoch).
		decode(&height, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch info (epoch: %s): %w", epoch, err)
	}

	votepowerInt := big.NewInt(0)
	err = newContractCall(s.evm, s.contracts.Votepower).
		execute(getProviderVotepower, provider, height).
		decode(&votepowerInt)
	if err != nil {
		return 0, fmt.Errorf("could not get provider votepower: %w", err)
	}

	votepowerFloat := big.NewFloat(0).SetInt(votepowerInt)
	votepower, _ := votepowerFloat.Float64()

	return votepower, nil
}

// Rewards returns the current unclaimed rewards for the FTSO data provider with the
// given address.
func (s *System) Rewards(provider common.Address) (float64, error) {

	current, err := s.Current()
	if err != nil {
		return 0, fmt.Errorf("could not get current epoch: %w", err)
	}
	epoch := big.NewInt(0).SetUint64(current)

	rewardsInt := big.NewInt(0)
	err = newContractCall(s.evm, s.contracts.Rewards).
		execute(getProviderRewards, epoch, provider).
		decode(&rewardsInt, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get provider rewards: %w", err)
	}

	rewardsFloat := big.NewFloat(0).SetInt(rewardsInt)
	rewards, _ := rewardsFloat.Float64()

	return rewards, nil
}
