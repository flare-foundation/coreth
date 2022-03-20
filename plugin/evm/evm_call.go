// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/internal/ethapi"
)

var errNoReturnData = errors.New("no return data")

type EVMBind struct {
	blockchain *core.BlockChain
}

type EVMSnapshot struct {
	blockchain *core.BlockChain
	hash       common.Hash
}

type EVMCall struct {
	blockchain *core.BlockChain
	hash       common.Hash
	contract   EVMContract
}

func BindEVM(blockchain *core.BlockChain) *EVMBind {
	return &EVMBind{blockchain: blockchain}
}

func (e *EVMBind) AtBlock(hash common.Hash) *EVMSnapshot {
	return &EVMSnapshot{blockchain: e.blockchain, hash: hash}
}

func (e *EVMSnapshot) OnContract(contract EVMContract) *EVMCall {
	return &EVMCall{blockchain: e.blockchain, hash: e.hash, contract: contract}
}

func (e *EVMCall) Execute(method string, params ...interface{}) *EVMReturn {

	header := e.blockchain.GetHeaderByHash(e.hash)
	if header == nil {
		return &EVMReturn{err: fmt.Errorf("unknown block (hash: %x)", e.hash)}
	}

	state, err := e.blockchain.StateAt(header.Root)
	if err != nil {
		return &EVMReturn{err: fmt.Errorf("could not get blockchain state (root: %x): %w", header.Root, err)}
	}

	data, err := e.contract.abi.Pack(method, params...)
	if err != nil {
		return &EVMReturn{err: fmt.Errorf("could not pack parameters: %w", err)}
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &e.contract.address, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return &EVMReturn{err: fmt.Errorf("could not convert arguments to message: %w", err)}
	}

	txContext := core.NewEVMTxContext(msg)
	blkContext := core.NewEVMBlockContext(header, e.blockchain, nil)
	chainConfig := e.blockchain.Config()
	evm := vm.NewEVM(blkContext, txContext, state, chainConfig, vm.Config{NoBaseFee: true})
	defer evm.Cancel()

	gp := new(core.GasPool).AddGas(math.MaxUint64)
	result, err := core.ApplyMessage(evm, msg, gp)
	if err != nil {
		return &EVMReturn{err: fmt.Errorf("could not apply message: %w", err)}
	}
	if result.Err != nil {
		return &EVMReturn{err: fmt.Errorf("could not execute transaction: %w", err)}
	}
	if len(result.ReturnData) == 0 {
		return &EVMReturn{err: errNoReturnData}
	}

	values, err := e.contract.abi.Unpack(method, result.ReturnData)
	if err != nil {
		return &EVMReturn{err: fmt.Errorf("could not unpack return data: %w", err)}
	}

	return &EVMReturn{values: values}

}
