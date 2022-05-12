package validators

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	"github.com/flare-foundation/flare/ids"
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type ValidatorRepository interface {
	Epoch() (uint64, error)
	Pending() (map[common.Address]ids.ShortID, error)
	Active() (map[ids.ShortID]uint64, error)
	Lookup(address common.Address) (ids.ShortID, error)

	SetEpoch(epoch uint64) error
	SetNodeID(address common.Address, nodeID ids.ShortID) error
	SetWeight(nodeID ids.ShortID, weight uint64) error
}

type Manager struct {
	log  zerolog.Logger
	repo ValidatorRepository
	ftso *FTSO
}

func (m *Manager) SetValidatorNodeID(address common.Address, nodeID ids.ShortID) error {
	// TODO: add to pending validator map stored under pending key in underlying DB
	return fmt.Errorf("not implemented")
}

func (m *Manager) UpdateActiveValidators() error {

	last, err := m.repo.LastEpoch()
	// TODO: check for bootstrapping state
	if err != nil {
		return fmt.Errorf("could not get last epoch: %w", err)
	}

	current, err := m.ftso.Current()
	if err != nil {
		return fmt.Errorf("could not get current epoch: %w", err)
	}

	if current < last {
		m.log.Warn().
			Uint64("current", current).
			Uint64("last", last).
			Msg("skipping active validators update (current bigger than last")
		return nil
	}

	if current == last {
		m.log.Debug().
			Uint64("epoch", current).
			Msg("skipping active validators update (epoch unchanged)")
		return nil
	}

	pending, err := m.repo.ActiveValidators()
	if err != nil {
		return fmt.Errorf("could not get active validators: %w", err)
	}

	for 
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

func (m *Manager) recalculateWeight(validators map[ids.ShortID]uint64, lastRewardEpoch uint64) error {
	votepowerCap, err := m.votepowerCap()
	if err != nil {
		return fmt.Errorf("could not get votepower cap: %w", err)
	}

	for validatorID := range validators {

		provider, err := m.validatorAddress(validatorID)
		if err != nil {
			return fmt.Errorf("could not get FTSO validator (provider: %s): %w", provider, err)
		}

		votepower, err := m.votepower(provider, lastRewardEpoch)
		if err != nil {
			return fmt.Errorf("could not get vote power (provider: %s): %w", provider, err)
		}
		if votepower == 0 {
			continue
		}

		if votepower > votepowerCap {
			votepower = votepowerCap
		}

		rewards, err := m.rewards(provider, lastRewardEpoch)
		if err != nil {
			return fmt.Errorf("could not get rewards (provider: %s): %w", provider, err)
		}
		if rewards == 0 {
			continue
		}

		weight := uint64(math.Pow(votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * rewards / votepower))

		validators[validatorID] = weight
	}

	return nil
}

func newValidatorRootHash(validators map[ids.ShortID]uint64, h common.Hash) []byte {
	panic("not implemented")
	return []byte{}
}
