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
	RootDegree      = 4
	RatioMultiplier = 100.0
)

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
	log       logging.Logger
	state     *validatordb.State
	transform ValidatorTransformer
}

func NewValidatorManager(log logging.Logger, state *validatordb.State, transform ValidatorTransformer) *ValidatorManager {

	m := ValidatorManager{
		log:       log,
		state:     state,
		transform: transform,
	}

	return &m
}

func (m *ValidatorManager) WithEVM(evm *vm.EVM) (vm.ValidatorSet, error) {

	ftso, err := ftso.NewSystem(evm)
	if err != nil {
		return nil, fmt.Errorf("could not initialize FTSO system: %w", err)
	}

	code := evm.StateDB.GetCode(params.ValidationAddress)
	root := common.BytesToHash(code)
	snapshot, err := m.state.WithRoot(root)
	if err != nil {
		return nil, fmt.Errorf("could not initialize validators state snapshot: %w", err)
	}

	s := ValidatorSet{
		log:       m.log,
		state:     evm.StateDB,
		ftso:      ftso,
		root:      root,
		transform: m.transform,
		snapshot:  snapshot,
	}

	return &s, nil
}
