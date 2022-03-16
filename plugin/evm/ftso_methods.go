// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

var DefaultFTSOMethods = FTSOMethods{
	EpochSeconds:     "rewardEpochDurationSeconds",
	ManagerAddress:   "getFtsoManager",
	RegistryAddress:  "getFtsoRegistry",
	EpochInfo:        "rewardEpochs",
	AssetIndices:     "getSupportedIndices",
	DataProviders:    "getFtsoWhitelistedPriceProviders",
	ValidatorNode:    "getNodeIdForDataProvider",
	VotePower:        "votePowerOf",
	UnclaimedRewards: "getUnclaimedReward",
}

type FTSOMethods struct {
	EpochSeconds     string
	ManagerAddress   string
	RegistryAddress  string
	EpochInfo        string
	AssetIndices     string
	DataProviders    string
	ValidatorNode    string
	VotePower        string
	UnclaimedRewards string
}
