package vm

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils/logging"
)

const (
	jsonValidator = `[{"inputs":[],"name":"getActiveValidators","outputs":[{"components":[{"internalType":"address[]","name":"providers","type":"address[]"},{"internalType":"bytes20","name":"nodeId","type":"bytes20"},{"internalType":"uint64","name":"weight","type":"uint64"}],"internalType":"structIValidatorRegistry.Validator[]","name":"_validators","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_provider","type":"address"}],"name":"getProviderNodeId","outputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_provider","type":"address"}],"name":"setProviderNodeId","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"updateActiveValidators","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
)

const (
	setMapping   = "setProviderNodeID"
	getMapping   = "getProviderNodeID"
	updateActive = "updateActiveValidators"
	getActive    = "getActiveValidators"
)

func init() {
	var err error
	validatorABI, err = abi.JSON(strings.NewReader(jsonValidator))
	if err != nil {
		panic(fmt.Sprintf("invalid ABI JSON: %s", err))
	}
}

var (
	registry     *validatorRegistry
	validatorABI abi.ABI
)

type ValidatorManager interface {
	WithEVM(evm *EVM) (ValidatorSet, error)
}

type ValidatorSet interface {
	SetMapping(provider common.Address, nodeID ids.ShortID) error
	GetMapping(provider common.Address) (ids.ShortID, error)

	UpdateValidators() error
	GetValidators() (validation.Set, error)

	Close() error
}

type validatorRegistry struct {
	log logging.Logger
	mgr ValidatorManager
}

func InjectDependencies(log logging.Logger, mgr ValidatorManager) {
	registry.log = log
	registry.mgr = mgr
}

func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, gas uint64, read bool) ([]byte, uint64, error) {

	snapshot, err := v.mgr.WithEVM(evm)
	if err != nil {
		return nil, 0, fmt.Errorf("could not initialize validator snapshot: %w", err)
	}

	method, err := validatorABI.MethodById(input[:4])
	if err != nil {
		return nil, 0, fmt.Errorf("could not get validator's method: %w", err)
	}

	var cost uint64
	switch method.Name {
	case setMapping:
		cost = 1_000_000
	case getMapping:
		cost = 200_000
	case updateActive:
		cost = 4_000_000
	case getActive:
		cost = 800_000
	default:
		cost = 100_000
	}
	if cost > gas {
		return nil, 0, ErrOutOfGas
	}
	gas = gas - cost

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, fmt.Errorf("could not unpack input: %w", err)
	}

	switch method.Name {

	case setMapping:

		provider := caller.Address()

		nodeID := args[0].(ids.ShortID)

		err := snapshot.SetMapping(provider, nodeID)
		if err != nil {
			return nil, gas, fmt.Errorf("could not set provider node: %w", err)
		}

		err = snapshot.Close()
		if err != nil {
			return nil, gas, fmt.Errorf("could not close validator snapshot: %w", err)
		}

		return nil, gas, nil

	case getMapping:

		provider := args[0].(common.Address)

		nodeID, err := snapshot.GetMapping(provider)
		if err != nil {
			return nil, gas, fmt.Errorf("could not get pending node: %w", err)
		}

		ret, err := validatorABI.Pack(getMapping, nodeID[:])
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack output %s: %w", method.Name, err)
		}

		return ret, gas, nil

	case updateActive:

		err = snapshot.UpdateValidators()
		if err != nil {
			return nil, gas, fmt.Errorf("could not update active validators: %w", err)
		}

		err = snapshot.Close()
		if err != nil {
			return nil, gas, fmt.Errorf("could not close validator snapshot: %w", err)
		}

		return nil, gas, nil

	case getActive:

		set, err := snapshot.GetValidators()
		if err != nil {
			return nil, gas, fmt.Errorf("could net get active validators: %w", err)
		}

		validators := make(map[ids.ShortID]uint64)
		for _, validator := range set.List() {
			validators[validator.ID()] = validator.Weight()
		}

		ret, err := validatorABI.Pack(getActive, validators)
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack output: %w", err)
		}

		return ret, gas, nil

	default:

		return nil, gas, ErrExecutionReverted
	}

}
