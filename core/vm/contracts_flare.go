package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare/ids"
)

var (
	registry *validatorRegistry
)

type ValidatorStorage interface {
	WithEVM(evm *EVM) (ValidatorManager, error)
}

type ValidatorManager interface {
	SetValidatorNodeID(address common.Address, nodeID ids.ShortID) error
	UpdateActiveValidators() error
	GetActiveValidators() (map[ids.ShortID]uint64, error)
}

type validatorRegistry struct {
	storage ValidatorStorage
}

func InitializeValidatorStorage(storage ValidatorStorage) {
	registry.SetValidatorStorage(storage)
}

func (v *validatorRegistry) SetValidatorStorage(storage ValidatorStorage) {
	v.storage = storage
}

func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, gas uint64, read bool) ([]byte, uint64, error) {

	manager, err := v.storage.WithEVM(evm)
	if err != nil {
		return nil, 0, err
	}

	// TODO switch-case the function call based on first 4 bytes of the input

	_ = manager

	return nil, 0, nil

}
