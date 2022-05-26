package evm

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/core/state/validators"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
	"github.com/flare-foundation/coreth/trie"
)

type ValidatorSet struct {
	log      logging.Logger
	state    vm.StateDB
	ftso     *ftso.System
	root     common.Hash
	snapshot *validators.Snapshot
}

func (v *ValidatorSet) SetValidator(provider common.Address, nodeID ids.ShortID) error {
	return v.snapshot.SetPending(provider, nodeID)
}

func (v *ValidatorSet) UpdateValidators() error {

	// Get the reward epoch for which validators are currently active.
	active, err := v.snapshot.GetEpoch()
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
		v.log.Warn("skipping active validators update (current epoch below active epoch, active: %d, current: %d", current, active)
		return nil
	}

	// If the current FTSO reward epoch is the same as the active epoch, we do not
	// need to do anything to update the validators.
	if current == active {
		v.log.Debug("skipping active validators update (active epoch unchanged, active: %d)", active)
		return nil
	}

	// Drop the current validator weights.
	err = v.snapshot.DropWeights()
	if err != nil {
		return fmt.Errorf("could not drop validator weights: %w", err)
	}

	// Get the list of whitelisted FTSO data providers with their votepower, as it
	// was stored in the previous reward epoch switchover.
	entries, err := v.snapshot.GetEntries()
	if err != nil {
		return fmt.Errorf("could not get providers: %w", err)
	}

	v.log.Info("processing %d validator entries for epoch %d", len(entries), current)

	// For each of these providers, we now get the rewards that they accumulated over
	// the previous epoch and calculate its validator weight.
	for _, entry := range entries {

		rewards, err := v.ftso.Rewards(entry.Provider)
		if err != nil {
			return fmt.Errorf("could not get rewards (provider: %s, node: %s): %w", entry.Provider, entry.NodeID, err)
		}

		if rewards == 0 {
			v.log.Debug("validator entry skipped (provider: %s, node: %s): no rewards", entry.Provider, entry.NodeID)
			continue
		}

		weight := uint64(math.Pow(entry.Votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * rewards / entry.Votepower))
		err = v.snapshot.SetWeight(entry.NodeID, weight)
		if err != nil {
			return fmt.Errorf("could not set validator weight (provider: %s, node: %s): %w", entry.Provider, entry.NodeID, err)
		}

		v.log.Debug("validator entry processed (provider: %s, node: %s, weight; %d)", entry.Provider, entry.NodeID, weight)
	}

	// Get the votepower cap for FTSO data providers. If they have more votepower
	// than the cap, they won't accumulate rewards for the excess votepower, which
	// means their performance would be deflated. So we use the cap in those cases.
	cap, err := v.ftso.Cap()
	if err != nil {
		return fmt.Errorf("could not get votepower cap: %w", err)
	}

	// Get the pending node IDs for data providers.
	pending, err := v.snapshot.AllPending()
	if err != nil {
		return fmt.Errorf("could not get pending validators: %w", err)
	}

	// Update the active node IDs for the data providers with the pending ones.
	for provider, nodeID := range pending {

		if nodeID == ids.ShortEmpty {
			err = v.snapshot.UnsetActive(provider)
			if err != nil {
				return fmt.Errorf("could not unset active validator (provider: %s, node: %s)", provider, nodeID)
			}
			continue
		}

		err = v.snapshot.SetActive(provider, nodeID)
		if err != nil {
			return fmt.Errorf("could not set active validator (provider: %s, node: %s)", provider, nodeID)
		}
	}

	// Unset the pending node IDs for FTSO data providers, as they have all been used now.
	err = v.snapshot.DropPending()
	if err != nil {
		return fmt.Errorf("could not drop pending validators: %w", err)
	}

	// Retrieve the whitelist of FTSO data providers as they are now, at the first
	// block of a new reward epoch. We will store this to be used on the next epoch
	// switchover, as we can only ever retrieve the whitelist at the current block.
	whitelist, err := v.ftso.Whitelist()
	if err != nil {
		return fmt.Errorf("could not get whitelist: %w", err)
	}

	v.log.Info("processing %d whitelisted providers for epoch %d", current+1)

	// For each provider, retrieve the votepower as it is now, at the first block
	// of the new reward epoch. Votepower can go down as the reward epoch goes on,
	// which could lead to an inflated performance rating if we would use the votepower
	// at a later block.
	entries = make([]validators.Entry, 0, len(whitelist))
	for _, provider := range whitelist {

		nodeID, err := v.snapshot.OneActive(provider)
		var missErr *trie.MissingNodeError
		if errors.As(err, &missErr) {
			v.log.Debug("whitelisted provider skipped (provider: %s, no node set)", provider)
			continue
		}
		if err != nil {
			return fmt.Errorf("could not get node (provider: %s): %w", provider, err)
		}

		votepower, err := v.ftso.Votepower(provider)
		if err != nil {
			return fmt.Errorf("could not get votepower (provider: %s, node: %s): %w", provider, nodeID, err)
		}

		if votepower == 0 {
			v.log.Debug("whitelisted provider skipped (provider: %s, node: %s, no vote power)", provider, nodeID)
			continue
		}

		if votepower > cap {
			v.log.Debug("whitelisted provider capped (provider: %s, node: %s, votepower: %d, cap: %d)", provider, nodeID, votepower, cap)
			votepower = cap
		}

		entry := validators.Entry{
			Provider:  provider,
			NodeID:    nodeID,
			Votepower: votepower,
		}

		entries = append(entries, entry)

		v.log.Debug("whitelisted provider processed (provider: %s, node: %s, votepower: %d)", provider, nodeID, votepower)
	}

	v.log.Info("prepared %d validator entries for epoch %d", current+1)

	// Sort the entries by node ID to always have the same order and avoid mismatches.
	sort.Slice(entries, func(i int, j int) bool {
		return bytes.Compare(entries[i].NodeID[:], entries[j].NodeID[:]) < 0
	})

	// Store the mapping of FTSO data providers to votepower; this will be used when
	// we calculate the validator weighting for the new reward epoch on the switchover
	// to the next reward epoch
	err = v.snapshot.SetEntries(entries)
	if err != nil {
		return fmt.Errorf("could not set providers: %w", err)
	}

	// Set the active epoch for validators to the new current epoch
	err = v.snapshot.SetEpoch(current)
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	return nil
}

func (v *ValidatorSet) GetValidators() (map[ids.ShortID]uint64, error) {
	return v.snapshot.AllWeights()
}

func (v *ValidatorSet) GetActiveNodeID(provider common.Address) (ids.ShortID, error) {
	return v.snapshot.OneActive(provider)
}

func (v *ValidatorSet) GetPendingNodeID(provider common.Address) (ids.ShortID, error) {
	return v.snapshot.OnePending(provider)
}

func (v *ValidatorSet) GetActiveProvider(nodeID ids.ShortID) (common.Address, error) {
	return v.snapshot.LookupActive(nodeID)
}

func (v *ValidatorSet) GetPendingProvider(nodeID ids.ShortID) (common.Address, error) {
	return v.snapshot.LookupPending(nodeID)
}

func (v *ValidatorSet) Close() error {

	root, err := v.snapshot.RootHash()
	if err != nil {
		return fmt.Errorf("could not get validator state root: %w", err)
	}

	if root == v.root {
		return nil
	}

	v.state.SetCode(params.ValidationAddress, root[:])

	return nil
}
