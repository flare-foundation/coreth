// (c) 2022, Flare Networks Limited. All rights reserved.

package params

import (
	"github.com/ethereum/go-ethereum/common"
)

var FlareContractUpdates = []ContractUpdate{}

type ContractUpdate struct {
	Address     common.Address
	OldByteCode []byte
	NewByteCode []byte
}
