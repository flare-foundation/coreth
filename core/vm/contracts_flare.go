package vm

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

const (
	jsonValidator = `[{"inputs":[{"internalType":"address","name":"_dataProvider","type":"address"}],"name":"getActiveNodeID","outputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"name":"getActiveValidator","outputs":[{"internalType":"address","name":"_dataProvider","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_dataProvider","type":"address"}],"name":"getPendingNodeID","outputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"name":"getPendingValidator","outputs":[{"internalType":"address","name":"_dataProvider","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"name":"setValidatorNodeID","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"updateActiveValidators","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
)

const (
	setValidator       = "setValidatorNodeID"
	updateValidators   = "updateActiveValidators"
	getActiveNodeID    = "getActiveNodeID"
	getActiveProvider  = "getActiveValidator"
	getPendingNodeID   = "getPendingNodeID"
	getPendingProvider = "getPendingValidator"
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
	WithEVM(evm *EVM) (ValidatorSnapshot, error)
}

type ValidatorSnapshot interface {
	SetValidator(provider common.Address, nodeID ids.ShortID) error

	UpdateValidators() error

	GetPendingNodeID(provider common.Address) (ids.ShortID, error)
	GetPendingProvider(nodeID ids.ShortID) (common.Address, error)

	GetActiveNodeID(provider common.Address) (ids.ShortID, error)
	GetActiveProvider(nodeID ids.ShortID) (common.Address, error)
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
	case setValidator:
		cost = 1_000_000
	case updateValidators:
		cost = 4_000_000
	default:
		cost = 200_000
	}
	if cost > gas {
		return nil, 0, ErrOutOfGas
	}
	gas = gas - cost

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, fmt.Errorf("could not get validator's method: %w", err)
	}

	switch method.Name {

	case setValidator:

		provider := caller.Address()

		nodeID := args[0].(ids.ShortID)

		err := snapshot.SetValidator(provider, nodeID)
		if err != nil {
			return nil, gas, fmt.Errorf("could not set provider node: %w", err)
		}

	case updateValidators:

		err = snapshot.UpdateValidators()
		if err != nil {
			return nil, gas, fmt.Errorf("could not update active validators: %w", err)
		}

	case getPendingNodeID:

		provider := args[0].(common.Address)

		nodeID, err := snapshot.GetPendingNodeID(provider)
		if err != nil {
			return nil, gas, fmt.Errorf("could not get pending node: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, nodeID[:])
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, gas, nil

	case getActiveNodeID:

		provider := args[0].(common.Address)

		nodeID, err := snapshot.GetActiveNodeID(provider)
		if err != nil {
			return nil, gas, fmt.Errorf("could not get active node: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, nodeID[:])
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, gas, nil

	case getPendingProvider:

		nodeID := args[0].(ids.ShortID)

		provider, err := snapshot.GetPendingProvider(nodeID)
		if err != nil {
			return nil, gas, fmt.Errorf("could not get pending validator: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, provider)
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, gas, nil

	case getActiveProvider:

		nodeID := args[0].(ids.ShortID)

		provider, err := snapshot.GetActiveProvider(nodeID)
		if err != nil {
			return nil, gas, fmt.Errorf("could not get active validator: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, provider)
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, gas, nil
	}

	return nil, gas, nil
}
