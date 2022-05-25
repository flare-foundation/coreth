package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/core/state/valstate"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type ValidatorDB interface {
	SetEntries(epoch uint64, entries []valstate.Entry) error
	GetEntries(epoch uint64) ([]valstate.Entry, error)

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
	log logging.Logger
	db  ValidatorDB
}

func NewValidatorManager(log logging.Logger, db ValidatorDB) *ValidatorManager {

	m := ValidatorManager{
		log: log,
		db:  db,
	}

	return &m
}

func (m *ValidatorManager) WithEVM(evm *vm.EVM) (vm.ValidatorSnapshot, error) {

	ftso, err := ftso.NewSystem(evm)
	if err != nil {
		return nil, fmt.Errorf("could not initialize FTSO system: %w", err)
	}

	s := ValidatorSnapshot{
		mgr:   m,
		ftso:  ftso,
		state: evm.StateDB,
	}

	return &s, nil
}
