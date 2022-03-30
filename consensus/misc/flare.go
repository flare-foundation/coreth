// (c) 2019-2020, Flare Networks Limited.

package misc

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/state"
)

type CoreContractUpdate struct {
	Address common.Address
	Old     []byte
	New     []byte
}

func updateApplied(statedb *state.StateDB, update CoreContractUpdate) bool {
	return bytes.Equal(statedb.GetCode(update.Address), update.New)
}

func applyUpdate(statedb *state.StateDB, update CoreContractUpdate) {
	if bytes.Equal(statedb.GetCode(update.Address), update.Old) {
		statedb.SetCode(update.Address, update.New)
	}
}
