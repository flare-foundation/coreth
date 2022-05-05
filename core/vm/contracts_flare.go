package vm

import (
	"github.com/ethereum/go-ethereum/common"
)

type validatorRegistry struct {
}

func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, gas uint64, read bool) ([]byte, uint64, error) {

	// TODO
	// 1) Initialize FTSO system around the EVM.
	// 2) Retrieve current reward epoch from EVM.
	// 3) Retrieve last reward epoch from FTSO system.
	// 4) Check against last validator list in precompiled contract state (just some DB wrapped around the EVM DB).
	// 5) If the reward epoch increased, retrieve all provider details and compute FTSO provider weights.
	// 6) If new weights calculated, insert into the DB (by epoch) and update latest epoch.
	// 7) If we stored a new latest epoch, set precompiled contract address' byte code to hash of validator state.

}
