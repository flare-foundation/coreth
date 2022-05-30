package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/core/state/validatordb"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
)

const (
	RootDegree      = 4     // which root to take of the votepower for validator weight (quadratic root)
	RatioMultiplier = 100.0 // how much to multiply votepower by to have more reasonable numbers
)

// FTSOSystem represents an abstraction layer on top of our code that interacts with
// the EVM state in order to retrieve information from the FTSO smart contracts. It
// allows us to use the relevant parameters and computations to calculate validator
// weights.
type FTSOSystem interface {
	Current() (uint64, error)
	Cap() (float64, error)
	Whitelist() ([]common.Address, error)
	Votepower(provider common.Address) (float64, error)
	Rewards(provider common.Address) (float64, error)
}

// EVM describes the EVM in terms of a single function to set code at a given location.
// It's there just to simplify testing of the validator snapshot.
type EVM interface {
	SetCode(address common.Address, code []byte)
}

// ValidationManager is responsible for initializating validator state snapshots on
// top of the injected validator database repository. The injected transform is applied
// against the validator set when calculating the active validators for a new epoch.
type ValidatorManager struct {
	log       logging.Logger
	repo      *validatordb.Repository
	transform ValidatorTransformer
}

// NewValidatorManager initializes a new validator manager on top of the given validator
// database repository and with the given validator transform.
func NewValidatorManager(log logging.Logger, repo *validatordb.Repository, transform ValidatorTransformer) *ValidatorManager {

	m := ValidatorManager{
		log:       log,
		repo:      repo,
		transform: transform,
	}

	return &m
}

// WithEVM initializes a new validator state snapshot on top of the given EVM state.
// It does so by initializing the FTSO smart contract wrappers on top of the EVM state,
// and by using the validator state root in the EVM state to initialize the validator
// state on top of the validator repository at the right point in time.
func (m *ValidatorManager) WithEVM(evm *vm.EVM) (vm.ValidatorSnapshot, error) {

	// We initialize the FTSO smart contract wrappers on top of the EVM for use by
	// the validator snapshot when calculating validator weights and related parameters.
	ftso, err := ftso.NewSystem(evm)
	if err != nil {
		return nil, fmt.Errorf("could not initialize FTSO system: %w", err)
	}

	// The validator state root is stored as the "code" located at the validator registry
	// address. The precompiled smart contract system circumvents any loading of contracts
	// stored at the addresses of precompiled contracts by directly calling into the
	// related Go code when these contracts are invoked. This means that the code under
	// this address is unused and can't be reached. We can thus reclaim it for this
	// purpose, allowing us to root the validator state in its corresponding EVM state.
	code := evm.StateDB.GetCode(params.ValidationAddress)
	if len(code) == 0 {
		err = m.bootstrap(evm.StateDB)
		code = evm.StateDB.GetCode(params.ValidationAddress)
	}
	if err != nil {
		return nil, fmt.Errorf("could not bootstrap validator state: %w", err)
	}

	root := common.BytesToHash(code)
	state, err := m.repo.WithRoot(root)
	if err != nil {
		return nil, fmt.Errorf("could not initialize validators state (root: %s): %w", root, err)
	}

	s := ValidatorSnapshot{
		log:       m.log,
		evm:       evm.StateDB,
		ftso:      ftso,
		root:      root,
		transform: m.transform,
		state:     state,
	}

	return &s, nil
}

func (m *ValidatorManager) bootstrap(evm vm.StateDB) error {

	state, err := m.repo.WithRoot(common.Hash{})
	if err != nil {
		return fmt.Errorf("could not open empty validator state: %w", err)
	}

	err = state.SetEpoch(0)
	if err != nil {
		return fmt.Errorf("could not bootstrap epoch: %w", err)
	}

	err = state.SetCandidates([]*validatordb.Candidate{})
	if err != nil {
		return fmt.Errorf("could not bootstrap candidates: %w", err)
	}

	err = state.SetValidators([]*validatordb.Validator{})
	if err != nil {
		return fmt.Errorf("could not bootstrap validators: %w", err)
	}

	root, err := state.RootHash()
	if err != nil {
		return fmt.Errorf("could not compute bootstrap hash: %w", err)
	}

	evm.SetCode(params.ValidationAddress, root[:])

	return nil
}
