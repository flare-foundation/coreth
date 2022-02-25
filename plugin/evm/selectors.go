package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func GetPriceSubmitterContract() string {
	switch {
	default:
		return "0x1000000000000000000000000000000000000003"
	}
}

func GetFtsoWhitelistedPriceProvidersSelector() []byte { //getFtsoWhitelistedPriceProviders(uint256)
	switch {
	default:
		return []byte{0x09, 0xfc, 0xb4, 0x00}
	}
} //0x 09 fc b4 00

func GetKeeperGasMultiplier(blockNumber *big.Int) uint64 {
	switch {
	default:
		return 100
	}
}

func GetCurrentRewardEpochSelector() []byte { //getCurrentRewardEpoch()
	switch {
	default:
		return []byte{0xe7, 0xc8, 0x30, 0xd4}
	}
} //e7 c8 30 d4

func GetRewardEpochVotePowerBlockSelector(currentRewardEpoch int64) []byte { //getRewardEpochVotePowerBlock(uint256)
	// todo use the currentRewardEpoch input
	switch {
	default:
		return []byte{0xf2, 0xed, 0xab, 0x5a}
	}
} // 0x f2 ed ab 5a

func VotePowerOfAtSelector() []byte { //votePowerOfAt(address,uint256)
	switch {
	default:
		return []byte{0x92, 0xbf, 0xe6, 0xd8}
	}
} // 0x 92 bf e6 d8

func TotalVotePowerSelector() []byte { //totalVotePower() //todo verify if this is the right function as it should take an input
	switch {
	default:
		return []byte{0xf5, 0xf3, 0xd4, 0xf7}
	}
} // f5 f3 d4 f7

func GetDataProviderCurrentFeePercentageSelector() []byte { //getDataProviderCurrentFeePercentage(address)
	switch {
	default:
		return []byte{0xcf, 0xbc, 0xd2, 0x5f}
	}
} //0x cf bc d2 5f

func GetUnclaimedRewardSelector() []byte { //getUnclaimedReward(uint256,address)
	switch {
	default:
		return []byte{0x65, 0x7d, 0x96, 0x95}
	}
} //0x 65 7d 96 95

func GetFtsoManagerSelector(blockTime *big.Int, chainID *big.Int) []byte { //getFtsoManager()
	switch {
	default:
		return []byte{0xb3, 0x9c, 0x68, 0x58}
	}
} //b3 9c 68 58

func GetVoterWhitelisterSelector() []byte { //getVoterWhitelister()
	switch {
	default:
		return []byte{0x71, 0xe1, 0xfa, 0xd9}
	}
} //0x 71 e1 fa d9

func GetWnatSelector() []byte { //wNat
	switch {
	default:
		return []byte{0x21, 0x18, 0xd5, 0xd0}
	}
} //21 18 d5 d0

func GetRewardManagerSelector() []byte { //rewardManager
	switch {
	default:
		return []byte{0xb7, 0xf4, 0x1e, 0x2d}
	}
} //b7 f4 1e 2d

func GetFtsosSelector() []byte { //getFtsos()
	switch {
	default:
		return []byte{0xce, 0x69, 0xf8, 0x33}
	}
} //0x ce 69 f8 33

func GetPriceSubmittedContract() common.Address { //PriceSubmitter contract
	switch {
	default:
		return common.HexToAddress("0x1000000000000000000000000000000000000003")
	}
}

func getCreatorsContractAddress() common.Address {
	return common.HexToAddress("0x1000000000000000000000000000000000000004")
}

func getCreatorsContractFunction4Bytes() []byte {
	switch {
	default:
		return []byte{0xe6, 0xad, 0xc1, 0xee} //getCreators()
	}
}

func getValidatorsContractFunction4Bytes() []byte {
	switch {
	default:
		return []byte{0xb7, 0xab, 0x4d, 0xb5} //getValidators()
	}
}

func GetFtsoManagerContract() []byte { //getFtsoManager()
	switch {
	default:
		return []byte{0xb3, 0x9c, 0x68, 0x58}
	}
} //0x b3 9c 68 58

func GetWnatContract(blockTime *big.Int) string {
	switch {
	default:
		return "0x02f0826ef6aD107Cfc861152B32B52fD11BaB9ED"
	}
}
