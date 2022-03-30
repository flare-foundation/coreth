// (c) 2019-2020, Flare Networks Limited.

package misc

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/flare-foundation/coreth/core/state"
)

type CoreContractUpdate struct {
	Address common.Address
	Old     []byte
	New     []byte
}

func updateApplied(statedb *state.StateDB, update CoreContractUpdate) bool {

	code := statedb.GetCode(update.Address)

	log.Debug("applied check: %s\n", update.Address)
	log.Debug("new code:\n%x\n", update.Old)
	log.Debug("actual code:\n%x\n", code)

	return bytes.Equal(statedb.GetCode(update.Address), update.New)
}

func applyUpdate(statedb *state.StateDB, update CoreContractUpdate) {

	code := statedb.GetCode(update.Address)

	log.Debug("apply update: %s\n", update.Address)
	log.Debug("current code:\n%x\n", update.Old)
	log.Debug("old code:\n%x\n", update.New)
	log.Debug("new code:\n%x\n", code)

	if bytes.Equal(statedb.GetCode(update.Address), update.Old) {
		log.Debug("applying update!")
		statedb.SetCode(update.Address, update.New)
	}
}
