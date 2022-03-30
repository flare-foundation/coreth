// (c) 2019-2020, Flare Networks Limited.

package misc

import (
	"bytes"
	"encoding/hex"

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

	log.Debug("update check", "address", update.Address.Hex(), "current code", hex.EncodeToString(code), "new code", hex.EncodeToString(update.New))

	return bytes.Equal(code, update.New)
}

func applyUpdate(statedb *state.StateDB, update CoreContractUpdate) {

	code := statedb.GetCode(update.Address)

	log.Debug("update apply", "address", update.Address.Hex(), "current code", hex.EncodeToString(code), "old code", hex.EncodeToString(update.Old), "new code", hex.EncodeToString(update.New))

	if bytes.Equal(statedb.GetCode(update.Address), update.Old) {
		log.Debug("applying update!")
		statedb.SetCode(update.Address, update.New)
	}
}
