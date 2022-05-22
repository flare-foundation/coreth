package vm

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/internal/ethapi"
)

var errNoReturnData = errors.New("no return data")

type contractCall struct {
	contract evmContract
	evm      *EVM
}

type evmContract struct {
	address common.Address
	abi     abi.ABI
}

func newContractCall(evm *EVM, contract evmContract) *contractCall {
	ec := &contractCall{
		contract: contract,
		evm:      evm,
	}
	return ec
}

func (e *contractCall) execute(method string, params ...interface{}) *contractReturn {

	data, err := e.contract.abi.Pack(method, params...)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not pack parameters: %w", err)}
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &e.contract.address, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not convert arguments to message: %w", err)}
	}

	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(e.evm, msg, gp)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not apply message: %w", err)}
	}
	if result.Err != nil {
		return &contractReturn{err: fmt.Errorf("could not execute transaction: %w", result.Err)}
	}
	if len(result.ReturnData) == 0 {
		return &contractReturn{err: errNoReturnData}
	}

	values, err := e.contract.abi.Unpack(method, result.ReturnData)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not unpack return data: %w", err)}
	}

	return &contractReturn{values: values}
}

type contractReturn struct {
	values []interface{}
	err    error
}

func (e *contractReturn) decode(values ...interface{}) error {

	if e.err != nil {
		return e.err
	}

	if len(e.values) != len(values) {
		return fmt.Errorf("invalid number of decode values (have: %d, want: %d)", len(values), len(e.values))
	}

	for i, val := range values {

		if val == nil {
			continue
		}

		ret := e.values[i]

		vv := reflect.ValueOf(val)
		if vv.IsNil() {
			continue
		}
		if vv.Kind() != reflect.Ptr {
			return fmt.Errorf("invalid non-pointer (index: %d, type: %T)", i, val)
		}

		iv := reflect.Indirect(vv)
		rv := reflect.ValueOf(ret)
		if iv.Kind() != rv.Kind() {
			return fmt.Errorf("invalid type for return value (have: %T, want: %T)", val, ret)
		}

		iv.Set(rv)
	}

	return nil
}
