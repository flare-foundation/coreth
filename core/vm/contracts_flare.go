package vm

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/core/state/validatordb"
)

// The lines below represents a Solidity contract interface that is representative
// of the exact behaviour of our precompiled smart contract. So rather than implementing
// the interface explicitly in Solidity, we wrote Go code that behaves in the exact
// same way, but has the ability to perform many more interactions at the storage
// level and do so in a more efficient way.
//
// pragma solidity >=0.7.6 <0.9;
//
// interface IValidatorRegistry {
//
//     struct Validator {
//         address[] providers;
//         bytes20 nodeId;
//         uint64 weight;
//     }
//
//     function setProviderNodeId(address _provider) external;
//     function getProviderNodeId(address _provider) external view returns (bytes20 _nodeId);
//
//     function updateActiveValidators() external;
//     function getActiveValidators() external view returns (Validator[] memory _validators);
// }

const (
	// jsonValidator represents the ABI of the above interface, which allows us to
	// use the Ethereum ABI library to interact with our precompiled contract, which
	// exhibits the same behaviour as a Solidity smart contract that implements the
	// above interface would.
	jsonValidation = `[{"inputs":[],"name":"getActiveValidators","outputs":[{"components":[{"internalType":"address[]","name":"providers","type":"address[]"},{"internalType":"bytes20","name":"nodeId","type":"bytes20"},{"internalType":"uint64","name":"weight","type":"uint64"}],"internalType":"structIValidatorRegistry.Validator[]","name":"_validators","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_provider","type":"address"}],"name":"getProviderNodeId","outputs":[{"internalType":"bytes20","name":"_nodeId","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_provider","type":"address"}],"name":"setProviderNodeId","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"updateActiveValidators","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
)

const (
	setMapping   = "setProviderNodeID"      // called by FTSO data providers to set their validator node ID
	getMapping   = "getProviderNodeID"      // exhibit of mapping between data providers and validators to Solidity
	updateActive = "updateActiveValidators" // called as part of a daemon call to update active validators
	getActive    = "getActiveValidators"    // exhibit total set of active validators with all data to Solidity
)

func init() {

	// We compile the validator registry ABI as part of an init call, as it is only
	// needed once, and should never fail because it is a static string.
	var err error
	validationABI, err = abi.JSON(strings.NewReader(jsonValidation))
	if err != nil {
		panic(fmt.Sprintf("invalid ABI JSON: %s", err))
	}
}

var (
	registry      *validatorRegistry // holds the validator registry precompiled contract
	validationABI abi.ABI            // holds the validator registry compiled contract ABI
)

// abiValidator is needed to cleanly encode validators to ABI when providing return
// data to contract calls.
type abiValidator struct {
	Providers []common.Address `abi:"providers"`
	NodeID    [20]byte         `abi:"nodeId"`
	Weight    uint64           `abi:"weight"`
}

// ValidatorManager wraps the behaviour of the validator manager into an interface
// in order to have proper isolation between packages. It can instantiate a snapshot
// of the validator state at a given EVM state.
type ValidatorManager interface {
	WithEVM(evm *EVM) (ValidatorSnapshot, error)
}

// ValidatorSnapshot represents the state of the validator database linked to a specific
// EVM state. It can be used to read active validator and pending mappings from the state,
// as well as the modification of the pending mappings and the active validators.
type ValidatorSnapshot interface {
	SetMapping(provider common.Address, nodeID ids.ShortID) error
	GetMapping(provider common.Address) (ids.ShortID, error)

	UpdateValidators() error
	GetValidators() ([]*validatordb.Validator, error)

	Close() error
}

// validatorRegistry is a precompiled contract that exhibits the validator database
// state to the Solidity layer. It comes with a logger to help with debugging and
// wraps around a validator manager which allows accessing the persisted validator
// state.
type validatorRegistry struct {
	log logging.Logger
	mgr ValidatorManager
}

// InjectDependencies is a work-around for not having a logger or the validator
// database available as part of the EVM state. A proper implementation would propagate
// these through all of the constructing functions and then read them from the EVM
// instantiation in the `Run` call, but that would require modifying a significant
// number on code paths and is thus an exercise for a later point. Injecting this
// once on start-up means that they will both remain available
func InjectDependencies(log logging.Logger, mgr ValidatorManager) {
	registry.log = log
	registry.mgr = mgr
}

// Run is called if a transaction is sent to the ValidatorRegistry precompiled contract
// at its configured address, and passed a reference to the EVM at the state the transaction
// is processed, a reference to the caller data, the called address (which should always
// be the configured one), the input data, the remaining gas and whether it is a read-only
// call. We then decode the input, run all the code we need, and encode the output data
// into the desired return format before returning it.
func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, gas uint64, read bool) ([]byte, uint64, error) {

	// Initialize a validator database snapshot as of the given EVM state. This will
	// read the validator state root hash from the EVM and use it to retrieve the
	// validator database snapshot at the correct state.
	snapshot, err := v.mgr.WithEVM(evm)
	if err != nil {
		return nil, 0, fmt.Errorf("could not initialize validator snapshot: %w", err)
	}

	// Retrieve the method from the contract ABI using the first 4 bytes of the
	// input data, as done by Solidity.
	method, err := validationABI.MethodById(input[:4])
	if err != nil {
		return nil, 0, fmt.Errorf("could not get validator's method: %w", err)
	}

	// Depending on the function that is called, we have different predetermined
	// gas costs.
	var cost uint64
	switch method.Name {

	// Setting the validator node ID should cost the provider a relevant amount,
	// in order to always be more expensive than other calls on the EVM and thus
	// not increase DoS surface in a meaningful way.
	case setMapping:
		cost = 1_000_000

	// Getting a mapping is a purely informational call, and will mostly just be
	// done as part of read-only calls. We give it a high cost here in order to
	// again avoid exposing DoS surface, as this cost is irrelevant for local reads.
	case getMapping:
		cost = 200_000

	// Updating the active validators is a no-op except for the first call in a
	// new reward epoch (= validator epoch). As this is called by the FTSO manager
	// as part of the epoch transition, it will _always_ be called first as part
	// of the daemon call, which has up to 10M gas available. Users will will never
	// want to call it due to the cost, because it has no effect, and thus again
	// doesn't exhibit DoS surface.
	case updateActive:
		cost = 4_000_000

	// Getting the active validators  is a purely information call as well, and will
	// again mostly be called as read-only, for local reads. If someone wants to read
	// this as part of smart contract logic, that is fine, as it's a single DB read
	// and decode, which is much more efficient than regular EVM state reads. However,
	// we still make it reasonably expensive to avoid abuse.
	case getActive:
		cost = 800_000

	// There are no other functions on the ABI, so if another function is called, the
	// call will certainly fail. However, we still want to make this no-op somewhat
	// expensive, as there are still some computations being made, and we don't want
	// to provide any DoS surface.
	default:
		cost = 100_000
	}

	// If the cost of the call is higher than the available gas, we use up all of
	// the gas and return an EVM out-of-gas error. Otherwise, we reduce the available
	// gas by the cost.
	if cost > gas {
		return nil, 0, ErrOutOfGas
	}
	gas = gas - cost

	// We use the ABI's method definition to unpack the input for the function.
	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, fmt.Errorf("could not unpack input: %w", err)
	}

	// Depending on the function name, where we use a list of constant names based
	// on the smart contract ABI, we will execute different code.
	switch method.Name {

	// When setting the mapping, we always use the caller's address and assume it
	// is only called by FTSO data providers who want to set their node ID. We then
	// set the mapping in the underlying persistent state and close the snapshot in
	// order to write the new validator state root hash to the EVM.
	case setMapping:

		provider := caller.Address()

		nodeID := args[0].(ids.ShortID)

		err := snapshot.SetMapping(provider, nodeID)
		if err != nil {
			return nil, gas, fmt.Errorf("could not set provider node: %w", err)
		}

		err = snapshot.Close()
		if err != nil {
			return nil, gas, fmt.Errorf("could not close validator snapshot: %w", err)
		}

		return nil, gas, nil

	// When getting the mapping, we read the data provider address from the first
	// argument, get the node ID from the underlying state and encode it using the
	// smart contract ABI before returning.
	case getMapping:

		provider := args[0].(common.Address)

		nodeID, err := snapshot.GetMapping(provider)
		if err != nil {
			return nil, gas, fmt.Errorf("could not get pending node: %w", err)
		}

		ret, err := method.Outputs.Pack(getMapping, [20]byte(nodeID))
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack output %s: %w", method.Name, err)
		}

		return ret, gas, nil

	// When updating the active validators, we simply call the corresponding function
	// on the validator state snapshot and then close it. Once again, closing the
	// snapshot will write its root hash to the EVM so that it stays fork-aware.
	case updateActive:

		err = snapshot.UpdateValidators()
		if err != nil {
			return nil, gas, fmt.Errorf("could not update active validators: %w", err)
		}

		err = snapshot.Close()
		if err != nil {
			return nil, gas, fmt.Errorf("could not close validator snapshot: %w", err)
		}

		return nil, gas, nil

	// When getting the active validators, we retrieve them from the underlying validator
	// state and then encode it appropriately before returning.
	// TODO: make sure we properly encode the validator structs for this return.
	case getActive:

		validators, err := snapshot.GetValidators()
		if err != nil {
			return nil, gas, fmt.Errorf("could net get active validators: %w", err)
		}

		outputs := make([]abiValidator, 0, len(validators))
		for _, validator := range validators {

			output := abiValidator{
				Providers: validator.Providers,
				NodeID:    [20]byte(validator.NodeID),
				Weight:    validator.Weight,
			}

			outputs = append(outputs, output)
		}

		ret, err := method.Outputs.Pack(getActive, outputs)
		if err != nil {
			return nil, gas, fmt.Errorf("could not pack output: %w", err)
		}

		return ret, gas, nil

	// Any other call is invalid and we fail with reverted execution error, so that
	// all previous potential changes are reversed as well.
	default:

		return nil, gas, ErrExecutionReverted
	}

}
