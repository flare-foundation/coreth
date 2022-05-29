package evm

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/core/state/validatordb"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
)

type ValidatorSnapshot struct {
	log       logging.Logger
	evm       vm.StateDB
	ftso      *ftso.System
	root      common.Hash
	transform ValidatorTransformer
	state     *validatordb.State
}

func (v *ValidatorSnapshot) SetValidator(provider common.Address, nodeID ids.ShortID) error {
	return v.state.SetMapping(provider, nodeID)
}

// UpdateValidators executes the logic for a validator epoch switchover. It will only
// execute if the current epoch in the FTSO system has increased and is higher than
// the epoch of the active validator set.
//
// If this is the case, it will execute the following steps:
//
// 1) get the validator candidates from the previous switchover, calculate their
// respective validator weights and store them as the new active validators.
//
// 2) get the whitelist of FTSO data providers and retrieve all the data to create
// the new list of validator candidates, and store them as the new candidates.
//
// 3) set the active epoch for the validator state to the current FTSO epoch.
func (v *ValidatorSnapshot) UpdateValidators() error {

	// Get the active epoch from the validator snapshot.
	active, err := v.state.GetEpoch()
	if err != nil {
		return fmt.Errorf("could not get active epoch: %w", err)
	}

	// Get the current epoch from the FTSO system.
	ftso, err := v.ftso.Current()
	if err != nil {
		return fmt.Errorf("could not get current FTSO epoch: %w", err)
	}

	// If the epoch on the FTSO system is below the active epoch on the validator
	// snapshot, we have a fatal error; this should _never_ happen.
	if active > ftso {
		return fmt.Errorf("snapshot epoch ahead of FTSO epoch (active: %d, ftso: %d)", active, ftso)
	}

	// If the current FTSO reward epoch is the same as the active epoch, we do not
	// need to update validators, as there are no changes.
	if active == ftso {
		v.log.Debug("skipping validator update (epoch: %d, no epoch change)", active)
		return nil
	}

	// In all other cases, the epoch on the FTSO is ahead of the active epoch on our
	// validator state, so we should execute an epoch switchover of the active validators.
	// As a first step, we should get the prepared validator entries. They contain the
	// information from the last epoch switchover that allows us to infer the  validator
	// weighting for the epoch switchover we are executing now.
	candidates, err := v.state.GetCandidates()
	if err != nil {
		return fmt.Errorf("could not get candidates: %w", err)
	}

	v.log.Info("processing %d candidates for starting epoch %d", len(candidates), ftso)

	// For each of the prepared validator entries, we get their unclaimed rewards at
	// the current FTSO system state. As this corresponds to the block where the rewards
	// are first released, nobody was able to claim them yet, and they serve as proxy for
	// the performance of the data provider over the last epoch. We then use the prepared
	// information (votepower and node ID) to calculate each data provider's validator weight.
	validators := make([]*validatordb.Validator, 0, len(candidates))
	for _, candidate := range candidates {

		totalRewards := float64(0)
		for _, provider := range candidate.Providers {

			rewards, err := v.ftso.Rewards(provider)
			if err != nil {
				return fmt.Errorf("could not get rewards (node: %s, provider: %s): %w", candidate.NodeID, provider, err)
			}

			totalRewards += rewards
		}

		if totalRewards == 0 {
			v.log.Debug("candidate skipped (node: %s): no rewards", candidate.NodeID)
			continue
		}

		weight := uint64(math.Pow(candidate.Votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * totalRewards / candidate.Votepower))

		validator := validatordb.Validator{
			Providers: candidate.Providers,
			NodeID:    candidate.NodeID,
			Weight:    weight,
		}
		validators = append(validators, &validator)

		v.log.Debug("candidate processed (node: %s, providers: %d, weight: %d)", candidate.NodeID, len(candidate.Providers), weight)
	}

	v.log.Info("obtained %d validators from %d candidates", len(validators), len(candidates))

	// Next, we apply the transform for the validators. This will in general do two things:
	// 1) add the default validators to the validator set with average FTSO validator weight; and
	// 2) normalize the total validator weight to a higher number for better sampling.
	// In the future, it can be used to phase out default validators and other fork-related changes.
	validators = v.transform.Transform(validators)

	// At this point, we are ready to set the new active validators to the computed set.
	err = v.state.SetValidators(validators)
	if err != nil {
		return fmt.Errorf("could not set validators: %w", err)
	}

	// Next, we update the current entries by adding and removing pending node IDs as
	// set by the data providers during the previous epoch. As a first step, we simply
	// retrieve all the pending node IDs, as set by the data providers.
	mappings, err := v.state.AllMappings()
	if err != nil {
		return fmt.Errorf("could not get mapings: %w", err)
	}

	// Next, we retrieve the whitelist of FTSO data providers as they are now,
	// at the first block of the starting reward epoch. Only data providers on this
	// list are eligible to become validators next epoch.
	whitelist, err := v.ftso.Whitelist()
	if err != nil {
		return fmt.Errorf("could not get whitelist: %w", err)
	}

	v.log.Info("processing %d providers for upcoming epoch %d", ftso+1)

	// We also need to retrieve the current cap on votepower from the FTSO system;
	// we need to apply it against each data provider's votepower to avoid a skew
	// in validation weight for data providers which have more votepower than this,
	// as it would mess with the performance calculation.
	cap, err := v.ftso.Cap()
	if err != nil {
		return fmt.Errorf("could not get votepower cap: %w", err)
	}

	// Now, for each data provider on the list, we see if it has a valid node ID
	// set in the mappings, and if it has any votepower. If it does, we cap the
	// votepower and add it to the candidate for the given node ID, and create that
	// candidate if it is the first provider mapping to that node ID.
	lookup := make(map[ids.ShortID]*validatordb.Candidate)
	for _, provider := range whitelist {

		nodeID, ok := mappings[provider]
		if !ok {
			v.log.Debug("provider skipped (address: %s): no mapping")
			continue
		}

		votepower, err := v.ftso.Votepower(provider)
		if err != nil {
			return fmt.Errorf("could not get votepower (address: %s, node: %s): %w", provider, nodeID, err)
		}

		if votepower == 0 {
			v.log.Debug("provider skipped (address: %s, node: %s): no votepower", provider, nodeID)
			continue
		}

		if votepower > cap {
			v.log.Debug("provider capped (address: %s, node: %s, votepower: %d, cap: %d)", provider, nodeID, votepower, cap)
			votepower = cap
		}

		candidate, ok := lookup[nodeID]
		if !ok {
			candidate = &validatordb.Candidate{
				Providers: []common.Address{},
				NodeID:    nodeID,
				Votepower: 0,
			}
			lookup[nodeID] = candidate
		}

		candidate.Providers = append(candidate.Providers, provider)
		candidate.Votepower += votepower

		v.log.Debug("provider processed (provider: %s, node: %s, votepower: %d)", provider, nodeID, votepower)
	}

	// We need to make sure to put candidates in a list and sort them deterministically
	// so that the storage hash remains the same across all nodes.
	candidates = make([]*validatordb.Candidate, 0, len(lookup))
	for _, candidate := range lookup {
		candidates = append(candidates, candidate)
	}
	sort.Slice(candidates, func(i int, j int) bool {
		return bytes.Compare(candidates[i].NodeID[:], candidates[j].NodeID[:]) < 0
	})

	v.log.Info("obtained %d candidates from %d providers", len(candidates), len(whitelist))

	// Store the mapping of FTSO data providers to votepower; this will be used when
	// we calculate the validator weighting for the new reward epoch on the switchover
	// to the next reward epoch.
	err = v.state.SetCandidates(candidates)
	if err != nil {
		return fmt.Errorf("could not set providers: %w", err)
	}

	// Set the active epoch for the validator snapshot to the current FTSO system epoch.
	err = v.state.SetEpoch(ftso)
	if err != nil {
		return fmt.Errorf("could not set active epoch: %w", err)
	}

	return nil
}

// GetValidators retrieves the active validators from the validator state and converts
// them to a default validation set as used by the consensus logic.
func (v *ValidatorSnapshot) GetValidators() (validation.Set, error) {

	validators, err := v.state.GetValidators()
	if err != nil {
		return nil, fmt.Errorf("could not get validators: %w", err)
	}

	set := validation.NewSet()
	for _, validator := range validators {

		err := set.AddWeight(validator.NodeID, validator.Weight)
		if err != nil {
			return nil, fmt.Errorf("could not set weight: %w", err)
		}
	}

	return set, nil
}

// SetMapping will set the mapping between an FTSO data provider's address and the
// node ID of the validator that will receive the validator weight attributed to it.
func (v *ValidatorSnapshot) SetMapping(provider common.Address, nodeID ids.ShortID) error {
	return v.state.SetMapping(provider, nodeID)
}

// GetMapping will return the node ID of the validator which receives the validator
// weight for the FTSO data provider with the given address.
func (v *ValidatorSnapshot) GetMapping(provider common.Address) (ids.ShortID, error) {
	return v.state.GetMapping(provider)
}

// Close will close the validator snapshot and make sure its current state is properly
// rooted into the EVM state by setting the code at the precompiled contract address
// to the root hash of the validator state.
func (v *ValidatorSnapshot) Close() error {

	// Get the root hash of the validator state. This will commit all changes to the
	// validator database and calculate the root hash on top of the corresponding
	// trie.
	root, err := v.state.RootHash()
	if err != nil {
		return fmt.Errorf("could not get validator state root: %w", err)
	}

	// We do not need to change the EVM state if the validator state root has not
	// changed.
	if root == v.root {
		return nil
	}

	// Otherwise, we set the code at the validator registry address to the validator
	// state root so that the validator state is part of the consensus state.
	v.evm.SetCode(params.ValidationAddress, root[:])

	return nil
}
