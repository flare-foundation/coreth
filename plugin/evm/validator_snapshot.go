package evm

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/core/state/valstate"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
)

type ValidatorSnapshot struct {
	mgr   *ValidatorManager
	ftso  *ftso.System
	state vm.StateDB
}

func (v *ValidatorSnapshot) SetValidatorNodeID(provider common.Address, nodeID ids.ShortID) error {

	err := v.mgr.db.SetPending(provider, nodeID)
	if err != nil {
		return fmt.Errorf("could not set pending validator: %w", err)
	}

	hash, err := v.mgr.db.RootHash()
	if err != nil {
		return fmt.Errorf("could not get validator state root hash: %w", err)
	}

	v.state.SetCode(params.ValidationAddress, hash[:])

	return nil
}

func (v *ValidatorSnapshot) UpdateActiveValidators() error {

	// Get the reward epoch for which validators are currently active.
	active, err := v.mgr.db.Epoch()
	if err != nil {
		return fmt.Errorf("could not get last epoch: %w", err)
	}

	// Get the current reward epoch of the FTSO system.
	current, err := v.ftso.Current()
	if err != nil {
		return fmt.Errorf("could not get current epoch: %w", err)
	}

	// The current FTSO reward epoch should never be smaller than the active epoch
	// of validators, otherwise a fatal error happened.
	if current < active {
		v.mgr.log.Warn("skipping active validators update (current epoch below active epoch, active: %d, current: %d", current, active)
		return nil
	}

	// If the current FTSO reward epoch is the same as the active epoch, we do not
	// need to do anything to update the validators.
	if current == active {
		v.mgr.log.Debug("skipping active validators update (active epoch unchanged, active: %d)", active)
		return nil
	}

	// Retrieve the whitelist of FTSO data providers as they are now, at the first
	// block of a new reward epoch. We will store this to be used on the next epoch
	// switchover, as we can only ever retrieve the whitelist at the current block.
	whitelist, err := v.ftso.Whitelist()
	if err != nil {
		return fmt.Errorf("could not get whitelist: %w", err)
	}

	// Get the list of whitelisted FTSO data providers with their votepower, as it
	// was stored in the previous reward epoch switchover.
	entries, err := v.mgr.db.GetEntries(active)
	if errors.Is(err, valstate.ErrNoEntries) {

		// TODO check if this doesn't exist in the DB yet; if so, we should bootstrap the system

	}
	if err != nil {
		return fmt.Errorf("could not get providers: %w", err)
	}

	var averageWeight uint64

	// For each of these providers, we now get the rewards that they accumulated over
	// the previous epoch and calculate its validator weight.
	for _, entry := range entries {

		rewards, err := v.ftso.Rewards(entry.Provider)
		if err != nil {
			return fmt.Errorf("could not get rewards (provider: %s): %w", entry.Provider, err)
		}

		if rewards == 0 {
			continue
		}

		weight := uint64(math.Pow(entry.Votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * rewards / entry.Votepower))

		if averageWeight == 0 {
			averageWeight += weight
		} else {
			averageWeight += (averageWeight + weight) / 2
		}

		err = v.mgr.db.SetWeight(current, entry.NodeID, weight)
		if err != nil {
			return fmt.Errorf("could not set validator weight (node: %s): %w", entry.NodeID, err)
		}
	}

	// TODO add the default validators with the average weight of the FTSO validators.

	_ = averageWeight

	// Get the votepower cap for FTSO data providers. If they have more votepower
	// than the cap, they won't accumulate rewards for the excess votepower, which
	// means their performance would be deflated. So we use the cap in those cases.
	cap, err := v.ftso.Cap()
	if err != nil {
		return fmt.Errorf("could not get votepower cap: %w", err)
	}

	// Get the pending node IDs for data providers.
	pending, err := v.mgr.db.Pending()
	if err != nil {
		return fmt.Errorf("could not get pending validators: %w", err)
	}

	// Update the active node IDs for the data providers with the pending ones.
	for provider, nodeID := range pending {

		if nodeID == ids.ShortEmpty {
			err = v.mgr.db.UnsetActive(provider)
			if err != nil {
				return fmt.Errorf("could not unset active (provider: %s)", provider)
			}
			continue
		}

		err = v.mgr.db.SetActive(provider, nodeID)
		if err != nil {
			return fmt.Errorf("could not set active (provider: %s, node: %s)", provider, nodeID)
		}
	}

	// Unset the pending node IDs for FTSO data providers, as they have all been used now.
	err = v.mgr.db.UnsetPending()
	if err != nil {
		return fmt.Errorf("could not unset pending: %w", err)
	}

	// For each provider, retrieve the votepower as it is now, at the first block
	// of the new reward epoch. Votepower can go down as the reward epoch goes on,
	// which could lead to an inflated performance rating if we would use the votepower
	// at a later block.
	entries = make([]valstate.Entry, 0, len(whitelist))
	for _, provider := range whitelist {

		nodeID, err := v.mgr.db.GetActive(provider)
		// TODO: check if it doesn't exist and simply continue in that case
		if err != nil {
			return fmt.Errorf("could not get node (provider: %s): %w", provider, err)
		}

		votepower, err := v.ftso.Votepower(provider)
		if err != nil {
			return fmt.Errorf("could not get votepower (provider: %s): %w", provider, err)
		}

		if votepower == 0 {
			continue
		}

		if votepower > cap {
			votepower = cap
		}

		entry := valstate.Entry{
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
	err = v.mgr.db.SetEntries(current, entries)
	if err != nil {
		return fmt.Errorf("could not set providers: %w", err)
	}

	// Set the active epoch for validators to the new current epoch
	err = v.mgr.db.SetEpoch(current)
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	// calculate new validator hash from all active validators with new weights;
	hash, err := v.mgr.db.RootHash()
	if err != nil {
		return fmt.Errorf("could not get root hash: %w", err)
	}

	// set the new root hash as byte code at the validator registry address in state DB.
	v.state.SetCode(params.ValidationAddress, hash[:])

	return nil
}

func (v *ValidatorSnapshot) GetActiveValidators() (map[ids.ShortID]uint64, error) {

	current, err := v.ftso.Current()
	if err != nil {
		return nil, fmt.Errorf("could not get current epoch: %w", err)
	}

	validators, err := v.mgr.db.Weights(current)
	if err != nil {
		return nil, fmt.Errorf("could not get active validators: %w", err)
	}

	return validators, nil
}

func (v *ValidatorSnapshot) GetActiveNodeID(provider common.Address) (ids.ShortID, error) {
	return v.mgr.db.GetActive(provider)
}

func (v *ValidatorSnapshot) GetPendingNodeID(provider common.Address) (ids.ShortID, error) {
	return v.mgr.db.GetPending(provider)
}

func (v *ValidatorSnapshot) GetActiveValidator(nodeID ids.ShortID) (common.Address, error) {
	return v.mgr.db.LookupActive(nodeID)
}

func (v *ValidatorSnapshot) GetPendingValidator(nodeID ids.ShortID) (common.Address, error) {
	return v.mgr.db.LookupPending(nodeID)
}

func entriesHash(s []valstate.Entry) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(s)
	if err != nil {
		return nil, fmt.Errorf("could not encode entries: %w", err)
	}
	return b.Bytes(), nil
}
