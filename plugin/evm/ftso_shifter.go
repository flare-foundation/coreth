// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ethereum/go-ethereum/common"
)

type FTSOShifter struct {
}

func NewFTSOShifter() *FTSOShifter {

	f := FTSOShifter{}

	return &f
}

func (f *FTSOShifter) ToBlock(hash common.Hash) FTSO {
	return nil
}
