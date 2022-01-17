package flareleader

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/flare-foundation/coreth/chain"
	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/coreth/core/types"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm"
	"github.com/flare-foundation/flare/ids"
	"math/big"
)

type leader struct {
	evm.VM
}

func New(vm evm.VM) *leader {
	return &leader{vm}
}

func (l *leader) GetValidators(hash chain.Hash) (map[ids.ID]float64, error) { // validators and their weights
	// todo make the evm call here after getting it from vm
	//l.VM.GetEthChain().BlockChain().GetBlock(hash, 1).Header() //todo what does number mean here and why do we need it if we already give hash??
	msg := types.NewMessage(
		common.Address{},  // from
		&common.Address{}, // to
		0,                 // nonce,
		nil,               // amount
		0,                 // gaslimit
		big.NewInt(5),     // gasprice
		nil,               // gasfeecap
		nil,               // gastipcap
		nil,               // data
		nil,               // accesslist
		true,              // isfake
	)
	blockchain := l.VM.GetEthChain().BlockChain()
	state, err := blockchain.State()
	if err != nil {
		return nil, fmt.Errorf("could not get blockchain state: %w", err)
	}

	tx := core.NewEVMTxContext(msg)
	header := &types.Header{
		BaseFee: nil,
		Number:  big.NewInt(1), //todo currently block height is 1. todo: api call needs to give us the block number we care about and parent hash
		// block number and hash.
		ParentHash: hash,
		Difficulty: big.NewInt(1),
	}
	header = l.VM.GetEthChain().BlockChain().GetBlock(hash, 1).Header()
	//block := core.NewEVMBlockContext(block.Header(), f.blockchain, nil)
	block := core.NewEVMBlockContext(header, blockchain, nil)
	chainConfig := params.ChainConfig{
		ChainID:             big.NewInt(4294967295),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
	}
	evm := vm.NewEVM(block, tx, state, &chainConfig, vm.Config{})
	evmCallValue := big.NewInt(0)
	//evmCallValue = nil
	caller := vm.AccountRef(getCreatorsContractAddress())
	creatorsByte, _, err := evm.Call(caller, getCreatorsContractAddress(), getValidatorsContractFunction4Bytes(), 100000, evmCallValue)
	//caller ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int
	if err != nil {
		return nil, fmt.Errorf("could not get block creators from contract: %w", err)
	}
	log.Info("result: ", "result: ", fmt.Sprintf("%x", creatorsByte))
	creators := make(map[ids.ID]float64) // todo make(map[string]float64)
	err = json.Unmarshal(creatorsByte, &creators)
	log.Info("creatorsByte..: ", "len(creatorsByte)", creatorsByte)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling error while trying to get block creators from contract: %w", err)
	}

	return creators, nil
}

func getCreatorsContractAddress() common.Address {
	return common.HexToAddress("0x1000000000000000000000000000000000000004")
}

func getCreatorsContractFunction4Bytes() []byte {
	switch {
	default:
		return []byte{0xe6, 0xad, 0xc1, 0xee} //getCreators()
	}
}

func getValidatorsContractFunction4Bytes() []byte {
	switch {
	default:
		return []byte{0xb7, 0xab, 0x4d, 0xb5} //getValidators()
	}
}

//getValidators() b7 ab 4d b5
