// (c) 2019-2020, Flare Networks Limited.

package misc

import (
	"github.com/flare-foundation/coreth/core/state"
)

// ApplyPotatoHardFork updates the Flare genesis smart contracts with updated
// versions when triggered at a given block with the given state DB.
func ApplyPotatoHardFork(statedb *state.StateDB) {
	// update the byte code for the Flare genesis contracts
}
