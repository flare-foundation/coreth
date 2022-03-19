package evm

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/accounts/abi/bind"
	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
	"github.com/flare-foundation/coreth/core"
)

func TestEVMCall_Execute(t *testing.T) {

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)

	balance := new(big.Int)
	balance.SetString("100000000000000000000", 10)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{
		Balance: balance,
	}
	gasLimit := uint64(9999999)

	be := backends.NewSimulatedBackend(alloc, gasLimit)

	abiStr := `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
	abiBin := `0x608060405234801561001057600080fd5b5060405161044f38038061044f8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186600182028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019080838360005b838110156100ba57808201518184015260208101905061009f565b50505050905090810190601f1680156100e75780820380516001836020036101000a031916815260200191505b50604052505050806000908051906020019061010492919061010b565b50506101b6565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826101415760008555610188565b82601f1061015a57805160ff1916838001178555610188565b82800160010185558215610188579182015b8281111561018757825182559160200191906001019061016c565b5b5090506101959190610199565b5090565b5b808211156101b257600081600090555060010161019a565b5090565b61028a806101c56000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806348f343f31461004657806354fd4d5014610088578063f56256c71461010b575b600080fd5b6100726004803603602081101561005c57600080fd5b8101908080359060200190929190505050610143565b6040518082815260200191505060405180910390f35b61009061015b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100d05780820151818401526020810190506100b5565b50505050905090810190601f1680156100fd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101416004803603604081101561012157600080fd5b8101908080359060200190929190803590602001909291905050506101f9565b005b60016020528060005260406000206000915090505481565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101f15780601f106101c6576101008083540402835291602001916101f1565b820191906000526020600020905b8154815290600101906020018083116101d457829003601f168201915b505050505081565b8060016000848152602001908152602001600020819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d48282604051808381526020018281526020019250505060405180910390a1505056fea2646970667358221220a7cedd091a0b3c6a40ba59ec51143153fbaa5bf19f0d8d0fd927e7f6dc84974164736f6c63430007060033`

	parsed, err := abi.JSON(strings.NewReader(abiStr))
	require.NoError(t, err)

	contractAddr, _, _, err := bind.DeployContract(auth, parsed, common.FromHex(abiBin), be, "v0")
	require.NoError(t, err)

	be.Commit(true)

	contract := EVMContract{
		address: contractAddr,
		abi:     parsed,
	}

	blockChain := be.Blockchain()

	evmBind := BindEVM(blockChain)

	t.Run("success_call", func(t *testing.T) {
		header := blockChain.GetHeaderByNumber(1)

		ret := evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("setItem", [32]byte{1}, [32]byte{1})
		require.NoError(t, ret.err)

		be.Commit(true)

		header = blockChain.GetHeaderByNumber(2)

		ret = evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("items", [32]byte{1})
		require.NoError(t, ret.err)
		require.Len(t, ret.values, 1)
	})

	t.Run("invalid_method", func(t *testing.T) {
		header := blockChain.GetHeaderByNumber(2)
		ret := evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("setItem2", [32]byte{1}, [32]byte{1})
		assert.EqualError(t, ret.err, "could not pack parameters: method 'setItem2' not found")
	})

	t.Run("invalid_params", func(t *testing.T) {
		header := blockChain.GetHeaderByNumber(2)
		ret := evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("setItem", [32]byte{1}, [32]byte{1}, [32]byte{1})
		assert.EqualError(t, ret.err, "could not pack parameters: argument count mismatch: got 3 for 2")
	})

	t.Run("invalid_header_number", func(t *testing.T) {
		header := blockChain.GetHeaderByNumber(2000)
		hash := header.Hash()
		ret := evmBind.AtBlock(hash).OnContract(contract).Execute("setItem", [32]byte{1}, [32]byte{1}, [32]byte{1})
		assert.EqualError(t, ret.err, fmt.Sprintf("unknown block (hash: %x)", hash))
	})
}
