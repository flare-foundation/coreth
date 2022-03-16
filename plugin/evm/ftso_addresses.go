// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ethereum/go-ethereum/common"
)

var DefaultFTSOAddresses = FTSOAddresses{
	Submitter:  common.HexToAddress(""),
	Validators: common.HexToAddress(""),
}

type FTSOAddresses struct {
	Submitter  common.Address
	Validators common.Address
}
