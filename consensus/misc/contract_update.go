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
	code := statedb.GetCode(update.Address)

	return bytes.Equal(code, update.New)
}

func applyUpdate(statedb *state.StateDB, update CoreContractUpdate) {
	code := statedb.GetCode(update.Address)

	if bytes.Equal(code, update.Old) {
		statedb.SetCode(update.Address, update.New)
	}
}
