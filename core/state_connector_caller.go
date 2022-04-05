package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
)

// stateConnectorCaller is a lightweight wrapper around EVM that allows to
// change the context and adds ability to mock the EVM
type stateConnectorCaller struct {
	evm *vm.EVM
}

func newStateConnectorCaller(evm *vm.EVM) *stateConnectorCaller {
	return &stateConnectorCaller{evm: evm}
}

func (c *stateConnectorCaller) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	return c.evm.Call(caller, addr, input, gas, value)
}

func (c *stateConnectorCaller) SetBlockContext(bc vm.BlockContext) {
	c.evm.Context = bc
}

func (c *stateConnectorCaller) BlockContext() vm.BlockContext {
	return c.evm.Context
}
