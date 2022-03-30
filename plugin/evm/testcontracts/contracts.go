//go:build integration
// +build integration

//go:generate solc --abi --bin manager.sol --overwrite -o manager
//go:generate solc --abi --bin registry.sol --overwrite -o registry
//go:generate solc --abi --bin reward.sol --overwrite -o reward
//go:generate solc --abi --bin store.sol --overwrite -o store
//go:generate solc --abi --bin submitter.sol --overwrite -o submitter
//go:generate solc --abi --bin validator.sol --overwrite -o validator
//go:generate solc --abi --bin votepower.sol --overwrite -o votepower
//go:generate solc --abi --bin whitelist.sol --overwrite -o whitelist
//go:generate solc --abi --bin wnat.sol --overwrite -o wnat

//go:generate go-bindata -nometadata -o data.go -pkg testcontracts -ignore .+\.sol$ -ignore contracts.go ./...

// Package testcontracts contains smart contracts that are used for integration tests.
package testcontracts

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/accounts/abi/bind"
	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
)

// Slices epochs, rewardEpochPowerHeight, rewardEpochStartHeight and rewardEpochStartTime
// should all have the same length.
//
// When calling RewardEpoch method using 'epochs[i]' values 'rewardEpochPowerHeight[i]',
// 'rewardEpochStartHeight[i]', 'rewardEpochStartTime[i]' will be returned
//
func DeployManager(auth *bind.TransactOpts, be *backends.SimulatedBackend,
	rewardManager common.Address, rewardEpochDurationSeconds *big.Int, rewardEpochsStartTs *big.Int, currentRewardEpoch *big.Int,
	fraction *big.Int, epochs []*big.Int, rewardEpochPowerHeight []*big.Int, rewardEpochStartHeight []*big.Int, rewardEpochStartTime []*big.Int,
) common.Address {

	return deployTestContract(auth, be, managerManagerAbiBytes, managerManagerBinBytes, rewardManager, rewardEpochDurationSeconds,
		rewardEpochsStartTs, currentRewardEpoch, fraction, epochs, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime)
}

func ManagerABI() abi.ABI {
	return contractAbi(managerManagerAbiBytes)
}

func DeployFTSORegistry(auth *bind.TransactOpts, be *backends.SimulatedBackend, supportedIndices []*big.Int) common.Address {

	return deployTestContract(auth, be, registryRegistryAbiBytes, registryRegistryBinBytes, supportedIndices)
}

func FTSORegistryABI() abi.ABI {
	return contractAbi(registryRegistryAbiBytes)
}

// Slices of epochs, providers and unclaimedRewards should all have the same length.
//
// When calling ProviderRewards method using 'epochs[i]' and 'providers[i]'
// value 'unclaimedRewards[i]' will be returned.
//
func DeployReward(auth *bind.TransactOpts, be *backends.SimulatedBackend,
	wNat common.Address, epochs []*big.Int, providers []common.Address, unclaimedRewards []*big.Int,
) common.Address {

	return deployTestContract(auth, be, rewardRewardAbiBytes, rewardRewardBinBytes, wNat, epochs, providers, unclaimedRewards)
}

func RewardABI() abi.ABI {
	return contractAbi(rewardRewardAbiBytes)
}

func DeploySubmitter(auth *bind.TransactOpts, be *backends.SimulatedBackend,
	voterWhitelister common.Address, ftsoRegistry common.Address, ftsoManager common.Address,
) common.Address {

	return deployTestContract(auth, be, submitterSubmitterAbiBytes, submitterSubmitterBinBytes, voterWhitelister, ftsoRegistry, ftsoManager)
}

func SubmitterABI() abi.ABI {
	return contractAbi(submitterSubmitterAbiBytes)
}

// Slice of dataProvidersAddresses should have the same length as nodes outer slice
//
// When calling ProviderNode method using 'dataProvidersAddresses[i]'
// value 'nodes[i]' will be returned.
//
func DeployValidator(auth *bind.TransactOpts, be *backends.SimulatedBackend,
	dataProvidersAddresses []common.Address, nodes [][20]byte,
) common.Address {

	return deployTestContract(auth, be, validatorValidatorAbiBytes, validatorValidatorBinBytes, dataProvidersAddresses, nodes)
}

func ValidatorABI() abi.ABI {
	return contractAbi(validatorValidatorAbiBytes)
}

// Slices of providers and vps should all have the same length.
//
// When calling ProviderVotepower method using 'providers[i]'
// value 'vps[i]' will be returned.
//
func DeployVotepower(auth *bind.TransactOpts, be *backends.SimulatedBackend,
	providers []common.Address, vps []*big.Int,
) common.Address {

	return deployTestContract(auth, be, votepowerVotepowerAbiBytes, votepowerVotepowerBinBytes, providers, vps)
}

func VotepowerABI() abi.ABI {
	return contractAbi(votepowerVotepowerAbiBytes)
}

func DeployWhitelist(auth *bind.TransactOpts, be *backends.SimulatedBackend,
	ftsoIndices []*big.Int, priceProvidersAddresses []common.Address,
) common.Address {

	return deployTestContract(auth, be, whitelistWhitelistAbiBytes, whitelistWhitelistBinBytes, ftsoIndices, priceProvidersAddresses)
}

func WhitelistABI() abi.ABI {
	return contractAbi(whitelistWhitelistAbiBytes)
}

func DeployWNAT(auth *bind.TransactOpts, be *backends.SimulatedBackend, wNat common.Address, totalSupply *big.Int) common.Address {

	return deployTestContract(auth, be, wnatWnatAbiBytes, wnatWnatBinBytes, wNat, totalSupply)
}

func WnatABI() abi.ABI {
	return contractAbi(wnatWnatAbiBytes)
}

func DeployStore(auth *bind.TransactOpts, be *backends.SimulatedBackend) common.Address {

	return deployTestContract(auth, be, storeStoreAbiBytes, storeStoreBinBytes)
}

func StoreABI() abi.ABI {
	return contractAbi(storeStoreAbiBytes)
}

func deployTestContract(auth *bind.TransactOpts, be *backends.SimulatedBackend, abiBytes func() ([]byte, error), abiBinBytes func() ([]byte, error), params ...interface{}) common.Address {

	abi := contractAbi(abiBytes)

	abiBin, err := abiBinBytes()
	checkErr(err)

	addr, _, _, err := bind.DeployContract(auth, abi, common.FromHex(string(abiBin)), be, params...)
	checkErr(err)

	return addr
}

func contractAbi(abiBytes func() ([]byte, error)) abi.ABI {
	rawAbi, err := abiBytes()
	checkErr(err)

	abiJSON, err := abi.JSON(bytes.NewReader(rawAbi))
	checkErr(err)

	return abiJSON
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
