package validators

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type ValidatorRepository interface {
	Epoch() (uint64, error)
	Pending() (map[common.Address]ids.ShortID, error)
	Active() (map[common.Address]ids.ShortID, error)
	Weights(epoch uint64) (map[ids.ShortID]uint64, error)
	Lookup(provider common.Address) (ids.ShortID, error)

	SetPending(provider common.Address, nodeID ids.ShortID) error
	SetEpoch(epoch uint64) error
	SetActive(provider common.Address, nodeID ids.ShortID) error
	SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error

	UnsetPending() error
	UnsetActive(address common.Address) error
}

type Manager struct {
	log  logging.Logger
	repo ValidatorRepository
	ftso *FTSO
}

func (m *Manager) GetActiveNodeID(provider common.Address) (ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (m *Manager) GetPendingNodeID(provider common.Address) (ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (m *Manager) GetActiveValidator(nodeID ids.ShortID) (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (m *Manager) GetPendingValidator(nodeID ids.ShortID) (common.Address, error) {
	// TODO implement me
	panic("implement me")
}

func NewManager(log logging.Logger, repo ValidatorRepository, ftso *FTSO) *Manager {

	m := Manager{
		log:  log,
		repo: repo,
		ftso: ftso,
	}

	return &m
}

func (m *Manager) SetValidatorNodeID(provider common.Address, nodeID ids.ShortID) error {

	err := m.repo.SetPending(provider, nodeID)
	if err != nil {
		return fmt.Errorf("could not set pending validator: %w", err)
	}

	// TODO: calculate and update root hash for the pending validator set

	return nil
}

func (m *Manager) UpdateActiveValidators() error {

	// TODO: do proper error handling on missing repository entries or undeployed
	// FTSO contracts

	active, err := m.repo.Epoch()
	if err != nil {
		return fmt.Errorf("could not get last epoch: %w", err)
	}

	current, err := m.ftso.Current()
	if err != nil {
		return fmt.Errorf("could not get current epoch: %w", err)
	}

	if current < active {
		m.log.Warn("skipping active validators update (current epoch below active epoch, active: %d, current: %d", current, active)
		return nil
	}

	if current == active {
		m.log.Debug("skipping active validators update (active epoch unchanged, active: %d)", active)
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

	validators, err := m.repo.Active()
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

	for provider, nodeID := range validators {

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

		err = m.repo.SetWeight(current, nodeID, weight)
		if err != nil {
			return fmt.Errorf("could not set validator weight (node: %s): %w", nodeID, err)
		}
	}

	err = m.repo.SetEpoch(current)
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	// TODO:
	// 1) get previous root hash in byte code on validator registry address from state DB;
	// 2) calculate new validator hash from all active validators with new weights;
	// 3) calculate new root hash as concat hash of old root hash and new validator hash;
	// 4) set the new root hash as byte code at the validator registry address in state DB.

	return nil
}

func (m *Manager) GetActiveValidators() (map[ids.ShortID]uint64, error) {

	current, err := m.ftso.Current()
	if err != nil {
		return nil, fmt.Errorf("could not get current epoch: %w", err)
	}

	validators, err := m.repo.Weights(current)
	if err != nil {
		return nil, fmt.Errorf("could not get active validators: %w", err)
	}

	return validators, nil
}
