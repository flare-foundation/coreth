package evm

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/flare-foundation/coreth/core/vm"
	"math"
	"math/big"
)

func GetFTSOManagerContract(evm *vm.EVM, blockTime *big.Int, chainID *big.Int) (common.Address, error) { //"0xbfA12e4E1411B62EdA8B035d71735667422A6A9e"
	FTSOManagerContractBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		common.HexToAddress(GetPriceSubmitterContract()),
		GetFtsoManagerSelector(chainID, blockTime),
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0))
	if err != nil {
		return common.Address{}, err
	}

	FTSOManagerContract := common.BytesToAddress(FTSOManagerContractBytes)

	return FTSOManagerContract, nil
}

func GetValidatorsWithWeight(evm *vm.EVM) (map[common.Address]float64, error) { //  return map of ftso price providers address and float64

	ftsoManagerContractAddress, err := GetFTSOManagerContract(evm, nil, nil)
	if err != nil {
		return nil, err
	}

	getCurrentRewardEpochBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		ftsoManagerContractAddress,
		GetCurrentRewardEpochSelector(),
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	//currentRewardEpoch := int64(binary.BigEndian.Uint64(getCurrentRewardEpochBytes))

	//instructions := append(getAttestationSelector[:], currentRoundNumber[:]...)
	input := append(GetRewardEpochVotePowerBlockSelector(), getCurrentRewardEpochBytes...)
	getRewardEpochVotePowerBlockBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		ftsoManagerContractAddress,
		input,
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	//rewardEpochVotePowerBlock := int64(binary.BigEndian.Uint64(getRewardEpochVotePowerBlockBytes))
	//fmt.Println(rewardEpochVotePowerBlock)
	ftsosBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		ftsoManagerContractAddress,
		GetFtsosSelector(),
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return nil, err
	}
	NUM_FTSOS := len(ftsosBytes) / 32
	var ftsos []common.Address
	for i := 0; i < NUM_FTSOS; i++ {
		ftsos = append(ftsos, common.BytesToAddress(ftsosBytes[i*32:(i+1)*32]))
	}

	ftsoContract := ftsos[0]
	// Now from this ftso contract retrieve wNat
	wNatBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		ftsoContract,
		GetWnatSelector(), //todo check if this is the correct way. If not, now to get public variables?
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return nil, err
	}

	wNatContract := common.BytesToAddress(wNatBytes)

	priceProviderAddresses, err := GetFtsoPriceProviderAddresses(evm)
	if err != nil {
		return nil, err
	}

	validatorsWithWeights := make(map[common.Address]float64)

	for _, priceProviderAddress := range priceProviderAddresses {

		inputDelegate := append(VotePowerOfAtSelector(), priceProviderAddress[:]...)
		inputDelegate = append(inputDelegate, getRewardEpochVotePowerBlockBytes...)

		delegateWeiBytes, _, err := evm.Call(
			vm.AccountRef(common.Address{}),
			wNatContract,
			inputDelegate,
			GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
			big.NewInt(0)) // uint256

		if err != nil {
			return nil, err
		}

		delegatedAmount := int64(binary.BigEndian.Uint64(delegateWeiBytes)) / 1000000000000000000.0

		totalVotePowerBytes, _, err := evm.Call(
			vm.AccountRef(common.Address{}),
			wNatContract,
			TotalVotePowerSelector(),
			GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
			big.NewInt(0)) // uint256

		if err != nil {
			return nil, err
		}

		totalVotePowerAmount := int64(binary.BigEndian.Uint64(totalVotePowerBytes)) / 1000000000000000000.0

		votePower := float64(delegatedAmount) / float64(totalVotePowerAmount) * 100.0

		// get rewardManager contract
		FTSORewardManagerContractBytes, _, err := evm.Call(
			vm.AccountRef(common.Address{}),
			ftsoManagerContractAddress,
			GetRewardManagerSelector(),
			GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
			big.NewInt(0)) // uint256

		if err != nil {
			return nil, err
		}

		FTSORewardManagerContract := common.BytesToAddress(FTSORewardManagerContractBytes)

		dataProviderCurrentFeePercentageBytes, _, err := evm.Call(
			vm.AccountRef(common.Address{}),
			FTSORewardManagerContract,
			GetDataProviderCurrentFeePercentageSelector(),
			GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
			big.NewInt(0)) // uint256

		if err != nil {
			return nil, err
		}

		dataProviderCurrentFeePercentage := float64(binary.BigEndian.Uint64(dataProviderCurrentFeePercentageBytes))
		fee := float64(dataProviderCurrentFeePercentage) / 100.0

		inputUnclaimedReward := append(GetUnclaimedRewardSelector(), getCurrentRewardEpochBytes...)
		inputUnclaimedReward = append(inputUnclaimedReward, priceProviderAddress[:]...) //getUnclaimedReward(currRewardEpoch, FTSO)

		rewardsBytes, _, err := evm.Call(
			vm.AccountRef(common.Address{}),
			FTSORewardManagerContract,
			inputUnclaimedReward,
			GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
			big.NewInt(0)) // uint256

		if err != nil {
			return nil, err
		}

		rewards := int64(binary.BigEndian.Uint64(rewardsBytes))

		reward_rate := float64(rewards) / float64(delegatedAmount) * (1 - fee/100)

		validationWeight := math.Log(1+votePower) * reward_rate // log(1+votePower) * reward_rate

		validatorsWithWeights[priceProviderAddress] = validationWeight

	}
	return validatorsWithWeights, nil
}

func GetFtsoPriceProviderAddresses(evm *vm.EVM) ([]common.Address, error) {
	log.Info("GetDefaultAttestors called...2")
	// Get VoterWhitelister contract
	voterWhitelisterContractBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		common.HexToAddress(GetPriceSubmitterContract()),
		GetVoterWhitelisterSelector(),
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0))
	g := GetKeeperGasMultiplier(evm.Context.BlockNumber) * evm.Context.GasLimit
	log.Info("Gas in evm call: ", "gas", g)
	log.Info("Gas in evm call: ", "GetKeeperGasMultiplier", GetKeeperGasMultiplier(evm.Context.BlockNumber))
	log.Info("Gas in evm call: ", "GasLimit", evm.Context.GasLimit)
	if err != nil {
		return []common.Address{}, err
	}
	// Get FTSO prive providers
	voterWhitelisterContract := common.BytesToAddress(voterWhitelisterContractBytes)
	priceProvidersBytes, _, err := evm.Call(
		vm.AccountRef(common.Address{}),
		voterWhitelisterContract,
		GetFtsoWhitelistedPriceProvidersSelector(),
		GetKeeperGasMultiplier(evm.Context.BlockNumber)*evm.Context.GasLimit,
		big.NewInt(0))
	if err != nil {
		return []common.Address{}, err
	}
	NUM_WHITELISTED_PRICE_PROVIDERS := len(priceProvidersBytes) / 32
	var ftsosAddresses []common.Address
	for i := 0; i < NUM_WHITELISTED_PRICE_PROVIDERS; i++ {
		ftsosAddresses = append(ftsosAddresses, common.BytesToAddress(priceProvidersBytes[i*32:(i+1)*32]))
	}
	return ftsosAddresses, nil
}
