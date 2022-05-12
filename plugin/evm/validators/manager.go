package validators

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type Manager struct {
	db        database.Database
	evm       *vm.EVM
	contracts Contracts
}

type Contracts struct {
	Registry  evmContract
	Manager   evmContract
	Rewards   evmContract
	Whitelist evmContract
	WNAT      evmContract
	Votepower evmContract
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

func (m *Manager) validatorAddress(nodeID ids.ShortID) (common.Address, error) {
	return common.Address{}, fmt.Errorf("not implemented")
}

func newValidatorRootHash(validators map[ids.ShortID]uint64, h common.Hash) []byte {
	panic("not implemented")
	return []byte{}
}

func (m *Manager) votepowerCap() (float64, error) {

	supply := &big.Int{}
	err := newContractCall(m.evm, m.contracts.WNAT).execute(TotalSupply).decode(&supply)
	if err != nil {
		return 0, fmt.Errorf("could not get total supply: %w", err)
	}

	fraction := &big.Int{}
	err = newContractCall(m.evm, m.contracts.Manager).
		execute(Settings).
		decode(&fraction, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get votepower threshold fraction: %w", err)
	}

	capInt := big.NewInt(0).Div(supply, fraction)
	capFloat := big.NewFloat(0).SetInt(capInt)
	cap, _ := capFloat.Float64()

	return cap, nil
}

func (m *Manager) votepower(provider common.Address, epoch uint64) (float64, error) {
	vpInt := &big.Int{}
	err := newContractCall(m.evm, m.contracts.Votepower).
		execute(ProviderVotepower, provider).
		decode(&vpInt)
	if err != nil {
		return 0, fmt.Errorf("could not get provider votepower: %w", err)
	}

	vpFloat := big.NewFloat(0).SetInt(vpInt)
	votepower, _ := vpFloat.Float64()

	return votepower, nil
}

func (m *Manager) rewards(provider common.Address, epoch uint64) (float64, error) {

	rwInt := &big.Int{}
	err := newContractCall(m.evm, m.contracts.Rewards).
		execute(ProviderRewards, big.NewInt(0).SetUint64(epoch), provider).
		decode(&rwInt, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get provider rewards: %w", err)
	}

	rwFloat := big.NewFloat(0).SetInt(rwInt)
	rewards, _ := rwFloat.Float64()

	return rewards, nil
}

func (m *Manager) currentEpoch() (uint64, error) {

	epoch := &big.Int{}
	err := newContractCall(m.evm, m.contracts.Rewards).
		execute(CurrentEpoch).
		decode(&epoch)
	if err != nil {
		return 0, fmt.Errorf("could not execute current epoch retrieval: %w", err)
	}

	return epoch.Uint64(), nil
}
