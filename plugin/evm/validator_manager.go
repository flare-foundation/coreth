package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/core/state/validators"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type ValidatorSnapshot interface {
	SetEntries(epoch uint64, entries []validators.Entry) error
	GetEntries(epoch uint64) ([]validators.Entry, error)

	SetPending(provider common.Address, nodeID ids.ShortID) error
	GetPending(provider common.Address) (ids.ShortID, error)

	SetActive(provider common.Address, nodeID ids.ShortID) error
	GetActive(provider common.Address) (ids.ShortID, error)

	Epoch() (uint64, error)
	Pending() (map[common.Address]ids.ShortID, error)
	Active() (map[common.Address]ids.ShortID, error)
	Weights(epoch uint64) (map[ids.ShortID]uint64, error)

	LookupPending(nodeID ids.ShortID) (common.Address, error)
	LookupActive(nodeID ids.ShortID) (common.Address, error)

	SetEpoch(epoch uint64) error
	SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error

	UnsetPending() error
	UnsetActive(address common.Address) error

	RootHash() (common.Hash, error)
}

type FTSOSystem interface {
	Current() (uint64, error)
	Cap() (float64, error)
	Whitelist() ([]common.Address, error)
	Votepower(provider common.Address) (float64, error)
	Rewards(provider common.Address) (float64, error)
}

type Consensus interface {
	SetCode(address common.Address, code []byte)
}

type ValidatorManager struct {
	log   logging.Logger
	state *validators.State
}

func NewValidatorManager(log logging.Logger, state *validators.State) *ValidatorManager {

	m := ValidatorManager{
		log:   log,
		state: state,
	}

	return &m
}

func (m *ValidatorManager) WithEVM(evm *vm.EVM) (vm.ValidatorSet, error) {

	ftso, err := ftso.NewSystem(evm)
	if err != nil {
		return nil, fmt.Errorf("could not initialize FTSO system: %w", err)
	}

	root := evm.StateDB.GetCode(params.ValidationAddress)
	snapshot, err := m.state.WithRoot(common.BytesToHash(root))
	if err != nil {
		return nil, fmt.Errorf("could not initialize validators state snapshot: %w", err)
	}

	s := ValidatorSet{
		ftso:     ftso,
		snapshot: snapshot,
	}

	return &s, nil
}
