package validators

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type Entry struct {
	Provider  common.Address
	NodeID    ids.ShortID
	Votepower float64
}

type ValidatorRepository interface {
	SetEntries(epoch uint64, entries []Entry) error
	GetEntries(epoch uint64) ([]Entry, error)

	SetPending(provider common.Address, nodeID ids.ShortID) error
	GetPending(provider common.Address) (ids.ShortID, error)

	SetActive(provider common.Address, nodeID ids.ShortID) error
	GetActive(provider common.Address) (ids.ShortID, error)

	Epoch() (uint64, error)
	Pending() (map[common.Address]ids.ShortID, error)
	Active() (map[common.Address]ids.ShortID, error)
	Weights(epoch uint64) (map[ids.ShortID]uint64, error)

	Lookup(nodeID ids.ShortID, prefix []byte) (common.Address, error)

	SetEpoch(epoch uint64) error
	SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error

	UnsetPending() error
	UnsetActive(address common.Address) error
}

type FTSOSystem interface {
	Current() (uint64, error)
	Cap() (float64, error)
	Whitelist() ([]common.Address, error)
	Votepower(provider common.Address) (float64, error)
	Rewards(provider common.Address) (float64, error)
	StateDB() vm.StateDB
}

type Manager struct {
	log  logging.Logger
	repo ValidatorRepository
	ftso FTSOSystem
}

func NewManager(log logging.Logger, repo ValidatorRepository, ftso FTSOSystem) *Manager {

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

	// Get the reward epoch for which validators are currently active.
	active, err := m.repo.Epoch()
	if err != nil {
		return fmt.Errorf("could not get last epoch: %w", err)
	}

	// Get the current reward epoch of the FTSO system.
	current, err := m.ftso.Current()
	if err != nil {
		return fmt.Errorf("could not get current epoch: %w", err)
	}

	// The current FTSO reward epoch should never be smaller than the active epoch
	// of validators, otherwise a fatal error happened.
	if current < active {
		m.log.Warn("skipping active validators update (current epoch below active epoch, active: %d, current: %d", current, active)
		return nil
	}

	// If the current FTSO reward epoch is the same as the active epoch, we do not
	// need to do anything to update the validators.
	if current == active {
		m.log.Debug("skipping active validators update (active epoch unchanged, active: %d)", active)
		return nil
	}

	// Get the list of whitelisted FTSO data providers with their votepower, as it
	// was stored in the previous reward epoch switchover.
	entries, err := m.repo.GetEntries(active)
	if errors.Is(err, errNoEntries) {
		// TODO: check if this doesn't exist in the DB yet; if so, we should bootstrap the system
	}
	if err != nil {
		return fmt.Errorf("could not get providers: %w", err)
	}

	// For each of these providers, we now get the rewards that they accumulated over
	// the previous epoch and calculate its validator weight.
	for _, entry := range entries {

		rewards, err := m.ftso.Rewards(entry.Provider)
		if err != nil {
			return fmt.Errorf("could not get rewards (provider: %s): %w", entry.Provider, err)
		}

		if rewards == 0 {
			continue
		}

		weight := uint64(math.Pow(entry.Votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * rewards / entry.Votepower))

		err = m.repo.SetWeight(current, entry.NodeID, weight)
		if err != nil {
			return fmt.Errorf("could not set validator weight (node: %s): %w", entry.NodeID, err)
		}
	}

	// TODO: add the default validators with the average weight of the FTSO validators.

	// Retrieve the whitelist of FTSO data providers as they are now, at the first
	// block of a new reward epoch. We will store this to be used on the next epoch
	// switchover, as we can only ever retrieve the whitelist at the current block.
	whitelist, err := m.ftso.Whitelist()
	if err != nil {
		return fmt.Errorf("could not get whitelist: %w", err)
	}

	// Get the votepower cap for FTSO data providers. If they have more votepower
	// than the cap, they won't accumulate rewards for the excess votepower, which
	// means their performance would be deflated. So we use the cap in those cases.
	cap, err := m.ftso.Cap()
	if err != nil {
		return fmt.Errorf("could not get votepower cap: %w", err)
	}

	// Get the pending node IDs for data providers.
	pending, err := m.repo.Pending()
	if err != nil {
		return fmt.Errorf("could not get pending validators: %w", err)
	}

	// Update the active node IDs for the data providers with the pending ones.
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

	// Unset the pending node IDs for FTSO data providers, as they have all been used now.
	err = m.repo.UnsetPending()
	if err != nil {
		return fmt.Errorf("could not unset pending: %w", err)
	}

	// For each provider, retrieve the votepower as it is now, at the first block
	// of the new reward epoch. Votepower can go down as the reward epoch goes on,
	// which could lead to an inflated performance rating if we would use the votepower
	// at a later block.
	entries = make([]Entry, 0, len(whitelist))
	for _, provider := range whitelist {

		nodeID, err := m.repo.GetActive(provider)
		// TODO: check if it doesn't exist and simply continue in that case
		if err != nil {
			return fmt.Errorf("could not get node (provider: %s): %w", provider, err)
		}

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

		entry := Entry{
			Provider:  provider,
			NodeID:    nodeID,
			Votepower: votepower,
		}

		entries = append(entries, entry)
	}

	// Sort the entries by node ID to always have the same order and avoid mismatches.
	sort.Slice(entries, func(i int, j int) bool {
		return bytes.Compare(entries[i].NodeID[:], entries[j].NodeID[:]) < 0
	})

	// Store the mapping of FTSO data providers to votepower; this will be used when
	// we calculate the validator weighting for the new reward epoch on the switchover
	// to the next reward epoch
	err = m.repo.SetEntries(current, entries)
	if err != nil {
		return fmt.Errorf("could not set providers: %w", err)
	}

	// Set the active epoch for validators to the new current epoch
	err = m.repo.SetEpoch(current)
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	stateDB := m.ftso.StateDB()

	// get previous root hash in byte code on validator registry address from state DB;
	hashByteCode := stateDB.GetCode(params.ValidationAddress)

	// calculate new validator hash from all active validators with new weights;
	valHash, err := hash(entries)
	if err != nil {
		return err
	}

	// calculate new root hash as concat hash of old root hash and new validator hash;
	newHash := crypto.Keccak256Hash(hashByteCode, valHash)

	// set the new root hash as byte code at the validator registry address in state DB.
	stateDB.SetCode(params.ValidationAddress, newHash.Bytes())

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

func (m *Manager) GetActiveNodeID(provider common.Address) (ids.ShortID, error) {
	return m.repo.GetActive(provider)
}

func (m *Manager) GetPendingNodeID(provider common.Address) (ids.ShortID, error) {
	return m.repo.GetPending(provider)
}

func (m *Manager) GetActiveValidator(nodeID ids.ShortID) (common.Address, error) {
	return m.repo.Lookup(nodeID, activePrefix)
}

func (m *Manager) GetPendingValidator(nodeID ids.ShortID) (common.Address, error) {
	return m.repo.Lookup(nodeID, pendingPrefix)
}

func hash(s []Entry) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(s)
	if err != nil {
		return nil, fmt.Errorf("could not encode entries: %w", err)
	}
	return b.Bytes(), nil
}
