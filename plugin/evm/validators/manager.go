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
	Active() (map[common.Address]ids.ShortID, error)
	Weights() (map[ids.ShortID]uint64, error)
	Lookup(provider common.Address) (ids.ShortID, error)

	SetPending(provider common.Address, nodeID ids.ShortID) error
	SetEpoch(epoch uint64) error
	SetActive(provider common.Address, nodeID ids.ShortID) error
	SetWeight(nodeID ids.ShortID, weight uint64) error

	UnsetPending() error
	UnsetActive(address common.Address) error
}

type Manager struct {
	log  zerolog.Logger
	repo ValidatorRepository
	ftso *FTSO
}

func (m *Manager) SetValidatorNodeID(provider common.Address, nodeID ids.ShortID) error {
	// TODO: add to pending validator map stored under pending key in underlying DB
	return fmt.Errorf("not implemented")
}

func (m *Manager) UpdateActiveValidators() error {

	// TODO: do proper error handling on missing repository entries or undeployed
	// FTSO contracts

	epoch, err := m.repo.Epoch()
	if err != nil {
		return fmt.Errorf("could not get last epoch: %w", err)
	}

	current, err := m.ftso.Current()
	if err != nil {
		return fmt.Errorf("could not get current epoch: %w", err)
	}

	if current < epoch {
		m.log.Warn().
			Uint64("epoch", epoch).
			Uint64("current", current).
			Msg("skipping active validators update (current epoch below active epoch)")
		return nil
	}

	if current == epoch {
		m.log.Debug().
			Uint64("epoch", epoch).
			Msg("skipping active validators update (epoch unchanged)")
		return nil
	}

	pending, err := m.repo.Pending()
	if err != nil {
		return fmt.Errorf("could not get pending validators: %w", err)
	}

	for provider, nodeID := range pending {

		if nodeID == ids.ShortEmpty {
			err = m.repo.UnsetActive(provider)
			if err != nil {
				return fmt.Errorf("could not unset active (provider: %s)", provider)
			}
			continue
		}

		err = m.repo.SetActive(provider, nodeID)
		if err != nil {
			return fmt.Errorf("could not set active (provider: %s, node: %s)", provider, nodeID)
		}
	}

	err = m.repo.UnsetPending()
	if err != nil {
		return fmt.Errorf("could not unset pending: %w", err)
	}

	active, err := m.repo.Active()
	if err != nil {
		return fmt.Errorf("could not get active validators: %w", err)
	}

	supply, err := m.ftso.Supply()
	if err != nil {
		return fmt.Errorf("could not get FTSO supply: %w", err)
	}

	fraction, err := m.ftso.Fraction()
	if err != nil {
		return fmt.Errorf("could not get votepower cap fraction: %w", err)
	}

	cap := supply / float64(fraction)

	for provider, nodeID := range active {

		votepower, err := m.ftso.Votepower(provider)
		if err != nil {
			return fmt.Errorf("could not get votepower (provider: %s): %w", provider, err)
		}
		if votepower == 0 {
			continue
		}

		if votepower > cap {
			votepower = cap
		}

		rewards, err := m.ftso.Rewards(provider, current)
		if err != nil {
			return fmt.Errorf("could not get rewards (provider: %s): %w", provider, err)
		}
		if rewards == 0 {
			continue
		}

		weight := uint64(math.Pow(votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * rewards / votepower))

		err = m.repo.SetWeight(nodeID, weight)
		if err != nil {
			return fmt.Errorf("could not set validator weight (node: %s): %w", nodeID, err)
		}
	}

	err = m.repo.SetEpoch(current)
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	return nil
}

func (m *Manager) GetActiveValidators() (map[ids.ShortID]uint64, error) {
	// TODO: return the active validator map stored under the active key in underlying DB
	return nil, fmt.Errorf("not implemented")
}

func newValidatorRootHash(validators map[ids.ShortID]uint64, h common.Hash) []byte {
	panic("not implemented")
	return []byte{}
}
