// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

package core

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/flare-foundation/coreth/core/vm"
)

// Define errors
type ErrInvalidKeeperData struct{}

func (e *ErrInvalidKeeperData) Error() string { return "invalid return data from keeper trigger" }

type ErrKeeperDataEmpty struct{}

func (e *ErrKeeperDataEmpty) Error() string { return "return data from keeper trigger empty" }

type ErrMaxMintExceeded struct {
	mintMax     *big.Int
	mintRequest *big.Int
}

func (e *ErrMaxMintExceeded) Error() string {
	return fmt.Sprintf("mint request of %s exceeded max of %s", e.mintRequest.Text(10), e.mintMax.Text(10))
}

type ErrMintNegative struct{}

func (e *ErrMintNegative) Error() string { return "mint request cannot be negative" }

// Define interface for dependencies
type EVMCaller interface {
	Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
	GetBlockNumber() *big.Int
	GetGasLimit() uint64
	AddBalance(addr common.Address, amount *big.Int)
}

// Define maximums that can change by block height
func GetKeeperGasMultiplier(blockNumber *big.Int) uint64 {
	switch {
	default:
		return 100
	}
}

func GetSystemTriggerContractAddr(blockNumber *big.Int) string {
	switch {
	default:
		return "0x1000000000000000000000000000000000000002"
	}
}

func GetSystemTriggerSelector(blockNumber *big.Int) []byte {
	switch {
	default:
		return []byte{0x7f, 0xec, 0x8d, 0x38}
	}
}

func GetPrioritisedFTSOContract(blockTime *big.Int) string { //PriceSubmitter
	switch {
	default:
		return "0x1000000000000000000000000000000000000003"
	}
}

func GetMaximumMintRequest(blockNumber *big.Int) *big.Int {
	switch {
	default:
		maxRequest, _ := new(big.Int).SetString("50000000000000000000000000", 10)
		return maxRequest
	}
}

func triggerKeeper(evm EVMCaller) (*big.Int, error) {
	bigZero := big.NewInt(0)
	// Get the contract to call
	systemTriggerContract := common.HexToAddress(GetSystemTriggerContractAddr(evm.GetBlockNumber()))
	// Call the method
	triggerRet, _, triggerErr := evm.Call(
		vm.AccountRef(systemTriggerContract),
		systemTriggerContract,
		GetSystemTriggerSelector(evm.GetBlockNumber()),
		GetKeeperGasMultiplier(evm.GetBlockNumber())*evm.GetGasLimit(),
		bigZero)
	// If no error and a value came back...
	if triggerErr == nil && triggerRet != nil {
		// Did we get one big int?
		if len(triggerRet) == 32 {
			// Convert to big int
			// Mint request cannot be less than 0 as SetBytes treats value as unsigned
			mintRequest := new(big.Int).SetBytes(triggerRet)
			// return the mint request
			return mintRequest, nil
		} else {
			// Returned length was not 32 bytes
			return bigZero, &ErrInvalidKeeperData{}
		}
	} else {
		if triggerErr != nil {
			return bigZero, triggerErr
		} else {
			return bigZero, &ErrKeeperDataEmpty{}
		}
	}
}

func mint(evm EVMCaller, mintRequest *big.Int) error {
	// If the mint request is greater than zero and less than max
	max := GetMaximumMintRequest(evm.GetBlockNumber())
	if mintRequest.Cmp(big.NewInt(0)) > 0 &&
		mintRequest.Cmp(max) <= 0 {
		// Mint the amount asked for on to the keeper contract
		evm.AddBalance(common.HexToAddress(GetSystemTriggerContractAddr(evm.GetBlockNumber())), mintRequest)
	} else if mintRequest.Cmp(max) > 0 {
		// Return error
		return &ErrMaxMintExceeded{
			mintRequest: mintRequest,
			mintMax:     max,
		}
	} else if mintRequest.Cmp(big.NewInt(0)) < 0 {
		// Cannot mint negatives
		return &ErrMintNegative{}
	}
	// No error
	return nil
}

func triggerKeeperAndMint(evm EVMCaller, log log.Logger) {
	// Call the keeper
	mintRequest, triggerErr := triggerKeeper(evm)
	// If no error...
	if triggerErr == nil {
		// time to mint
		if mintError := mint(evm, mintRequest); mintError != nil {
			log.Warn("Error minting inflation request", "error", mintError)
		}
	} else {
		log.Warn("Keeper trigger in error", "error", triggerErr)
	}
}
