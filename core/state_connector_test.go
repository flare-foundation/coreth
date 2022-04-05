package core

import (
	"errors"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/core/types"
	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
)

type mockMessage struct {
	from       common.Address
	to         common.Address
	gas        uint64
	gasPrice   *big.Int
	gasFeeCap  *big.Int
	gasTipCap  *big.Int
	value      *big.Int
	data       []byte
	accessList types.AccessList
}

func (m *mockMessage) From() common.Address         { return m.from }
func (m *mockMessage) To() *common.Address          { return &m.to }
func (m *mockMessage) GasPrice() *big.Int           { return m.gasPrice }
func (m *mockMessage) GasFeeCap() *big.Int          { return m.gasFeeCap }
func (m *mockMessage) GasTipCap() *big.Int          { return m.gasTipCap }
func (m *mockMessage) Value() *big.Int              { return m.value }
func (m *mockMessage) Gas() uint64                  { return m.gas }
func (m *mockMessage) Nonce() uint64                { return 0 }
func (m *mockMessage) Data() []byte                 { return m.data }
func (m *mockMessage) AccessList() types.AccessList { return m.accessList }
func (m *mockMessage) IsFake() bool                 { return true }

type mockStateConnectorCaller struct {
	context  vm.BlockContext
	msg      *mockMessage
	CallFunc func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
}

func (m *mockStateConnectorCaller) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64,
	value *big.Int) (ret []byte, leftOverGas uint64, err error) {

	return m.CallFunc(m.context, caller, addr, input, gas, value)
}

func (m *mockStateConnectorCaller) SetBlockContext(bc vm.BlockContext) { m.context = bc }

func (m *mockStateConnectorCaller) BlockContext() vm.BlockContext { return m.context }

func buildStateConnectorMock(from string, addr common.Address) (*mockMessage, *mockStateConnectorCaller) {
	msg := mockMessage{
		from: common.BytesToAddress([]byte(from)),
	}
	scc := mockStateConnectorCaller{
		msg: &msg,
		context: vm.BlockContext{
			Coinbase: addr,
		},
	}
	return &msg, &scc
}

func TestStateTransition_FinalisePreviousRound(t *testing.T) {
	currentRoundNumber := []byte("222")
	testCoinbaseAddress := common.BigToAddress(big.NewInt(1000))
	blockTime := big.NewInt(time.Date(2022, time.February, 28, 17, 1, 1, 0, time.UTC).Unix())

	t.Run("coston chain", func(t *testing.T) {
		t.Parallel()

		t.Run("nominal case", func(t *testing.T) {
			t.Parallel()

			mockMsg, mockSCC := buildStateConnectorMock("coston chain", testCoinbaseAddress)

			mockSCC.CallFunc = func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
				// this is used to return response for attestationResult
				if caller.Address() == costonDefaultAttestors[0] {
					rootHash := []byte("some_hash")
					return rootHash, 0, nil
				}

				// the rest is used for mocking finalization call

				if caller.Address() != stateConnectorCoinbaseSignalAddr(params.CostonChainID, blockTime) {
					return nil, 0, errors.New("caller address should be state connector coinbase signal addr")
				}

				if context.Coinbase == testCoinbaseAddress {
					return nil, 0, errors.New("original coinbase address should have been changed to state connector coinbase signal addr")
				}
				return nil, 0, nil
			}

			c := newConnector(mockSCC, mockMsg)
			err := c.finalizePreviousRound(params.CostonChainID, blockTime, currentRoundNumber)
			require.NoError(t, err)
			assert.Equal(t, testCoinbaseAddress, mockSCC.context.Coinbase, "coinbase address should be changed to the original address")
		})
		t.Run("handles count attestations error without reaching finalization", func(t *testing.T) {
			t.Parallel()

			mockMsg, mockSCC := buildStateConnectorMock("coston chain", testCoinbaseAddress)

			mockSCC.CallFunc = func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
				if caller.Address() == costonDefaultAttestors[0] {
					return nil, 0, errors.New("attestation error")
				}

				return nil, 0, errors.New("finalization should not be reached")
			}

			c := newConnector(mockSCC, mockMsg)
			err := c.finalizePreviousRound(params.CostonChainID, blockTime, currentRoundNumber)
			require.NoError(t, err)
		})
		t.Run("handles finalization error", func(t *testing.T) {
			t.Parallel()

			mockMsg, mockSCC := buildStateConnectorMock("coston chain", testCoinbaseAddress)

			mockSCC.CallFunc = func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
				// this is used to return response for attestationResult
				if caller.Address() == costonDefaultAttestors[0] {
					rootHash := []byte("some_hash")
					return rootHash, 0, nil
				}

				// the rest is used for mocking finalization call

				if caller.Address() != stateConnectorCoinbaseSignalAddr(params.CostonChainID, blockTime) {
					return nil, 0, errors.New("caller address should be state connector coinbase signal addr")
				}

				if context.Coinbase == testCoinbaseAddress {
					return nil, 0, errors.New("original coinbase address should have been changed to state connector coinbase signal addr")
				}
				return nil, 0, errors.New("finalization error")
			}

			c := newConnector(mockSCC, mockMsg)
			err := c.finalizePreviousRound(params.CostonChainID, blockTime, currentRoundNumber)
			assert.EqualError(t, err, "finalization error")
		})
		t.Run("with local attestors with finalization", func(t *testing.T) {
			t.Parallel()

			attestor1 := common.BigToAddress(big.NewInt(100))
			attestor2 := common.BigToAddress(big.NewInt(200))
			attestor3 := common.BigToAddress(big.NewInt(300))

			localAttestors := strings.Join([]string{attestor1.Hex(), attestor2.Hex(), attestor3.Hex()}, ",")

			oldEnv := os.Getenv(localAttestorEnv)
			defer os.Setenv(localAttestorEnv, oldEnv)
			_ = os.Setenv(localAttestorEnv, localAttestors)

			mockMsg, mockSCC := buildStateConnectorMock("default", testCoinbaseAddress)

			mockSCC.CallFunc = func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
				// this is used to return response for attestationResult
				if caller.Address() == attestor1 || caller.Address() == attestor2 || caller.Address() == attestor3 || caller.Address() == costonDefaultAttestors[0] {
					rootHash := []byte("some_hash")
					return rootHash, 0, nil
				}

				// the rest is used for mocking finalization call

				if caller.Address() != stateConnectorCoinbaseSignalAddr(params.CostonChainID, blockTime) {
					return nil, 0, errors.New("caller address should be state connector coinbase signal addr")
				}

				if context.Coinbase == testCoinbaseAddress {
					return nil, 0, errors.New("original coinbase address should have been changed to state connector coinbase signal addr")
				}
				return nil, 0, nil
			}

			c := newConnector(mockSCC, mockMsg)
			err := c.finalizePreviousRound(params.CostonChainID, blockTime, currentRoundNumber)
			require.NoError(t, err)
			assert.Equal(t, testCoinbaseAddress, mockSCC.context.Coinbase, "coinbase address should be changed to the original address")
		})

	})

	t.Run("default attestors", func(t *testing.T) {
		t.Parallel()

		attestor1 := common.BigToAddress(big.NewInt(10))
		attestor2 := common.BigToAddress(big.NewInt(20))
		attestor3 := common.BigToAddress(big.NewInt(30))

		defaultAttestors := strings.Join([]string{attestor1.Hex(), attestor2.Hex(), attestor3.Hex()}, ",")

		t.Run("nominal case", func(t *testing.T) {
			t.Parallel()

			oldEnv := os.Getenv(defaultAttestorEnv)
			defer os.Setenv(defaultAttestorEnv, oldEnv)

			_ = os.Setenv(defaultAttestorEnv, defaultAttestors)

			mockMsg, mockSCC := buildStateConnectorMock("default", testCoinbaseAddress)

			mockSCC.CallFunc = func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
				// this is used to return response for attestationResult
				if caller.Address() == attestor1 || caller.Address() == attestor2 || caller.Address() == attestor3 {
					rootHash := []byte("some_hash")
					return rootHash, 0, nil
				}

				// the rest is used for mocking finalization call

				if caller.Address() != stateConnectorCoinbaseSignalAddr(params.CostonChainID, blockTime) {
					return nil, 0, errors.New("caller address should be state connector coinbase signal addr")
				}

				if context.Coinbase == testCoinbaseAddress {
					return nil, 0, errors.New("original coinbase address should have been changed to state connector coinbase signal addr")
				}
				return nil, 0, nil
			}

			c := newConnector(mockSCC, mockMsg)
			err := c.finalizePreviousRound(params.CostonChainID, blockTime, currentRoundNumber)
			require.NoError(t, err)
			assert.Equal(t, testCoinbaseAddress, mockSCC.context.Coinbase, "coinbase address should be changed to the original address")
		})
		t.Run("handles count attestation errors without finalizing", func(t *testing.T) {
			t.Parallel()

			oldEnv := os.Getenv(defaultAttestorEnv)
			defer os.Setenv(defaultAttestorEnv, oldEnv)

			_ = os.Setenv(defaultAttestorEnv, defaultAttestors)

			mockMsg, mockSCC := buildStateConnectorMock("default", testCoinbaseAddress)

			mockSCC.CallFunc = func(context vm.BlockContext, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
				// this is used only to return response for attestationResult

				if caller.Address() == attestor1 {
					rootHash := []byte("some_hash")
					return rootHash, 0, nil
				}

				if caller.Address() == attestor2 {
					return nil, 0, errors.New("attestation error")
				}

				if caller.Address() == attestor3 {
					rootHash := []byte("some_other_hash")
					return rootHash, 0, nil
				}

				return nil, 0, errors.New("finalization should not be reached")
			}

			c := newConnector(mockSCC, mockMsg)
			err := c.finalizePreviousRound(params.CostonChainID, blockTime, currentRoundNumber)
			require.NoError(t, err)
			assert.Equal(t, testCoinbaseAddress, mockSCC.context.Coinbase, "coinbase address should not be changed")
		})
	})
}
