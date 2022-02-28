// (c) 2022, Flare Networks Limited. All rights reserved.

package misc

import (
	"bytes"

	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/params"
)

// ApplyFlareFork1Upgrades updates the Flare genesis smart contracts with updated
// versions when triggered at a given block with the given state DB. We make sure
// that the old byte code at each given address corresponds to what we expect
// before replacing it with the new byte code.
func ApplyFlareFork1Upgrades(statedb *state.StateDB) {
	for _, update := range params.FlareContractUpdates {
		byteCode := statedb.GetCode(update.Address)
		if !bytes.Equal(byteCode, update.OldByteCode) {
			continue
		}
		statedb.SetCode(update.Address, update.NewByteCode)
	}
}
