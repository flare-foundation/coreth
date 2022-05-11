// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flare-foundation/coreth/core/vm"

	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/internal/ethapi"
)

var errNoReturnData = errors.New("no return data")

type ContractCall struct {
	contract EVMContract
	evm      *vm.EVM
}

func NewContractCall(evm *vm.EVM, contract EVMContract) *ContractCall {
	ec := &ContractCall{
		contract: contract,
		evm:      evm,
	}
	return ec
}

func (e *ContractCall) Execute(method string, params ...interface{}) *ContractReturn {
	data, err := e.contract.abi.Pack(method, params...)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not pack parameters: %w", err)}
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &e.contract.address, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not convert arguments to message: %w", err)}
	}

	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(e.evm, msg, gp)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not apply message: %w", err)}
	}
	if result.Err != nil {
		return &ContractReturn{err: fmt.Errorf("could not execute transaction: %w", result.Err)}
	}
	if len(result.ReturnData) == 0 {
		return &ContractReturn{err: errNoReturnData}
	}

	values, err := e.contract.abi.Unpack(method, result.ReturnData)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not unpack return data: %w", err)}
	}

	return &ContractReturn{values: values}
}
