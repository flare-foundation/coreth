package vm

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/plugin/evm/validators"
)

const (
	// TODO add abi for packing/unpacking
	jsonValidator = ``
)

func init() {
	var err error
	validatorABI, err = abi.JSON(strings.NewReader(jsonValidator))
	if err != nil {
		panic(fmt.Sprintf("invalid ABI JSON: %s", err))
	}
}

var (
	registry *validatorRegistry

	validatorABI abi.ABI

	// TODO we might use actual method names (method.Name) provided in the ABI
	sigSetValidatorNodeID     = [4]byte{0x00, 0x00, 0x00, 0x00}
	sigUpdateActiveValidators = [4]byte{0x00, 0x00, 0x00, 0x00}
	sigGetPendingNodeID       = [4]byte{0x00, 0x00, 0x00, 0x00}
	sigGetActiveNodeID        = [4]byte{0x00, 0x00, 0x00, 0x00}
	sigGetActiveValidator     = [4]byte{0x00, 0x00, 0x00, 0x00}
	sigGetPendingValidator    = [4]byte{0x00, 0x00, 0x00, 0x00}
)

type ValidatorManager interface {
	SetValidatorNodeID(provider common.Address, nodeID ids.ShortID) error
	UpdateActiveValidators() error

	GetActiveNodeID(provider common.Address) (ids.ShortID, error)
	GetPendingNodeID(provider common.Address) (ids.ShortID, error)
	GetActiveValidator(nodeID ids.ShortID) (common.Address, error)
	GetPendingValidator(nodeID ids.ShortID) (common.Address, error)
}

type validatorRegistry struct {
	log     logging.Logger
	storage validators.ValidatorRepository
}

func InjectDependencies(log logging.Logger, storage validators.ValidatorRepository) {
	registry.log = log
	registry.storage = storage
}

func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, suppliedGas uint64, read bool) ([]byte, uint64, error) {
	var sig [4]byte
	copy(sig[:], input[:4])

	gasCost := v.requiredGas(sig)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost

	method, err := validatorABI.MethodById(input[:4])
	if err != nil {
		return nil, 0, fmt.Errorf("could not get validator's method: %w", err)
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, fmt.Errorf("could not get validator's method: %w", err)
	}

	ftso, err := NewFTSO(evm)
	if err != nil {
		return nil, 0, fmt.Errorf("could not initialize FTSO system: %w", err)
	}

	manager := validators.NewManager(v.log, v.storage, ftso)

	switch sig {

	case sigSetValidatorNodeID:

		provider := caller.Address()

		nodeID := args[0].(ids.ShortID)

		err := manager.SetValidatorNodeID(provider, nodeID)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not set provider node: %w", err)
		}

	case sigUpdateActiveValidators:

		err = manager.UpdateActiveValidators()
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not update active validators: %w", err)
		}

	case sigGetPendingNodeID:

		provider := args[0].(common.Address)

		nodeID, err := manager.GetPendingNodeID(provider)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not get pending node: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, nodeID[:])
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, suppliedGas, nil

	case sigGetActiveNodeID:

		provider := args[0].(common.Address)

		nodeID, err := manager.GetActiveNodeID(provider)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not get active node: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, nodeID[:])
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, suppliedGas, nil

	case sigGetPendingValidator:

		nodeID := args[0].(ids.ShortID)

		provider, err := manager.GetPendingValidator(nodeID)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not get pending validator: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, provider)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, suppliedGas, nil

	case sigGetActiveValidator:

		nodeID := args[0].(ids.ShortID)

		provider, err := manager.GetActiveValidator(nodeID)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not get active validator: %w", err)
		}

		ret, err := validatorABI.Pack(method.Name, provider)
		if err != nil {
			return nil, suppliedGas, fmt.Errorf("could not pack return values for method %s: %w", method.Name, err)
		}

		return ret, suppliedGas, nil
	}

	return nil, suppliedGas, nil
}

// TODO: define gas cost per function and check there is sufficient
func (v *validatorRegistry) requiredGas(method [4]byte) uint64 {
	switch method {

	case sigSetValidatorNodeID:

	case sigUpdateActiveValidators:

	case sigGetPendingNodeID:

	case sigGetActiveNodeID:

	case sigGetPendingValidator:

	case sigGetActiveValidator:

	}

	return 0
}
