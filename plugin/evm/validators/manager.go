package validators

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/core/vm"
)

type Manager struct {
	db  database.Database
	evm *vm.EVM
}

func (m *Manager) SetValidatorNodeID(address common.Address, nodeID ids.ShortID) error {
	// TODO: add to pending validator map stored under pending key in underlying DB
	return fmt.Errorf("not implemented")
}

func (m *Manager) UpdateActiveValidators() error {
	// TODO: check that we actually switched to the next reward epoch, otherwise we
	// should either fail hard, or just do nothing (check with Ilan); make sure to
	// also store the new epoch if we changed
	// TODO: for each validator node ID in the pending validator map, move it to the
	// same key in the active validator map; if the node ID is empty, delete the entry
	// from the active validator map instead
	// TODO: for each validator in the active map, recalculate its weight using the
	// the unclaimed rewards and the votepower (reuse the code we had)
	// TODO: compute the root hash of all new active validators, and hash together
	// with the hash stored as code in the EVM under the validator registry address,
	// and replace the previous hash with that new hash - this will ensure that the
	// full validator set and history is part of the consensus
	return fmt.Errorf("not implemented")
}

func (m *Manager) GetActiveValidators() (map[ids.ShortID]uint64, error) {
	// TODO: return the active validator map stored under the active key in underlying DB
	return nil, fmt.Errorf("not implemented")
}
