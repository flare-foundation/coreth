package ftso

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/internal/ethapi"
)

var (
	errNoReturnData = errors.New("no return data")
)

// contractCall is a convenience wrapper around the EVM to execute read-only smart
// contract calls from Go.
type contractCall struct {
	contract evmContract
	evm      *vm.EVM
}

// evmContract is a convenience wrapper grouping together a smart contract address
// and its corresponding compiled ABI.
type evmContract struct {
	address common.Address
	abi     abi.ABI
}

// newContractCall will create a new convenience wrapper to execute read-only smart
// contract calls against the given call on top of the provided EVM state.
func newContractCall(evm *vm.EVM, contract evmContract) *contractCall {

	c := &contractCall{
		contract: contract,
		evm:      evm,
	}

	return c
}

// execute executes the function with the given method name against the underlying
// smart contract encoding the provided parameters as arguments.
func (c *contractCall) execute(method string, params ...interface{}) *contractReturn {

	// We use the contract's compiled ABI to pack the arguments.
	data, err := c.contract.abi.Pack(method, params...)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not pack parameters: %w", err)}
	}

	// We then create the corresponding message/transaction that we can apply to the
	// EVM to execute the computation.
	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &c.contract.address, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not convert arguments to message: %w", err)}
	}

	// Next, we apply the message against the EVM with infinite gas, as this is a
	// read-only execution. If there was an error, or there is no return data, we
	// error, because a read-only execution without return data makes little sense.
	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(c.evm, msg, gp)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not apply message: %w", err)}
	}
	if result.Err != nil {
		return &contractReturn{err: fmt.Errorf("could not execute transaction: %w", result.Err)}
	}
	if len(result.ReturnData) == 0 {
		return &contractReturn{err: errNoReturnData}
	}

	// Finally, we unpack the return data into an interface slice that can later
	// be used to properly extract the types.
	values, err := c.contract.abi.Unpack(method, result.ReturnData)
	if err != nil {
		return &contractReturn{err: fmt.Errorf("could not unpack return data: %w", err)}
	}

	return &contractReturn{values: values}
}

// contractReturn is a convenience wrapper around the results of a smart contract
// execution that enables proper extraction of data types.
type contractReturn struct {
	values []interface{}
	err    error
}

// decode can be used to extract the proper data types from smart contract calls by
// copying them into the provided values.
func (c *contractReturn) decode(values ...interface{}) error {

	// We defer the evaluation of errors, because this is supposed to be a chainable
	// API. If an error occurred during execution of the smart contract, we return
	// it now.
	if c.err != nil {
		return c.err
	}

	// Otherwise, if more values were returned from the smart contract call than are
	// provided to the decode function, there is an error in the logic.
	if len(c.values) != len(values) {
		return fmt.Errorf("invalid number of decode values (have: %d, want: %d)", len(values), len(c.values))
	}

	// For each given value, we try to copy the corresponding return value into
	// it. `nil` values can be given for return values that should be ignored.
	for i, val := range values {

		if val == nil {
			continue
		}

		ret := c.values[i]

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
		if iv.Kind() == reflect.Array && rv.Len() != iv.Len() {
			return fmt.Errorf("invalid len for return array (have: %T, want: %T)", iv.Len(), rv.Len())
		}

		iv.Set(rv)
	}

	return nil
}
