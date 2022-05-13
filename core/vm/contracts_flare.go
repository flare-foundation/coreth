package vm

import (
	"fmt"

	"github.com/aws/smithy-go/logging"
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/plugin/evm/validators"
)

var (
	registry *validatorRegistry
)

var (
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

type ValidatorStorage interface {
}

type validatorRegistry struct {
	log     logging.Logger
	storage ValidatorStorage
}

func InjectDependencies(log logging.Logger, storage ValidatorStorage) {
	registry.log = log
	registry.storage = storage
}

func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, gas uint64, read bool) ([]byte, uint64, error) {

	// TODO: define gas cost per function and check there is sufficient
	ftso, err := NewFTSO(evm)
	if err != nil {
		return nil, 0, fmt.Errorf("could not initialize FTSO system: %w", err)
	}

	manager := validators.NewManager(v.log, v.storage, ftso)

	var sig [4]byte
	copy(sig[:], input[:4])
	switch sig {

	case sigSetValidatorNodeID:

		provider := caller.Address()

		// TODO: unpack node ID
		var nodeID ids.ShortID

		err := manager.SetValidatorNodeID(provider, nodeID)
		if err != nil {
			return nil, 0, fmt.Errorf("could not set provider node: %w", err)
		}

	case sigUpdateActiveValidators:

		err = manager.UpdateActiveValidators()
		if err != nil {
			return nil, 0, fmt.Errorf("could not update active validators: %w", err)
		}

	case sigGetPendingNodeID:

		// TODO: unpack provider address
		var provider common.Address

		nodeID, err := manager.GetPendingNodeID(provider)
		if err != nil {
			return nil, 0, fmt.Errorf("could not get pending node: %w", err)
		}

		// TODO: pack node ID
		return nodeID[:], 0, nil

	case sigGetActiveNodeID:

		// TODO: unpack provider address
		var provider common.Address

		nodeID, err := manager.GetActiveNodeID(provider)
		if err != nil {
			return nil, 0, fmt.Errorf("could not get active node: %w", err)
		}

		// TODO: pack node ID
		return nodeID[:], 0, nil

	case sigGetPendingValidator:

		// TODO: unpack node ID
		var nodeID ids.ShortID

		provider, err := manager.GetPendingValidator(nodeID)
		if err != nil {
			return nil, 0, fmt.Errorf("could not get pending validator: %w", err)
		}

		// TODO: pack provider address
		return provider[:], 0, nil

	case sigGetActiveValidator:

		// TODO: unpack node ID
		var nodeID ids.ShortID

		provider, err := manager.GetActiveValidator(nodeID)
		if err != nil {
			return nil, 0, fmt.Errorf("could not get active validator: %w", err)
		}

		// TODO: pack provider address
		return provider[:], 0, nil

	default:

		return nil, 0, fmt.Errorf("invalid function signature for validator registry (sig: %x)", sig)

	}

	return nil, gas, nil
}
