// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

var DefaultFTSOMethods = FTSOMethods{
	EpochSeconds:   "rewardEpochDurationSeconds",
	ManagerAddress: "getFtsoManager",
}

type FTSOMethods struct {
	EpochSeconds   string
	ManagerAddress string
}
