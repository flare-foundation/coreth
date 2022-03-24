// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/accounts/abi"
)

type EVMContract struct {
	address common.Address
	abi     abi.ABI
}
