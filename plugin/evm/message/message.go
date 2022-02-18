// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"

<<<<<<< HEAD
	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/units"
=======
	"github.com/flare-foundation/flare/codec"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/units"
>>>>>>> upstream-v0.8.5-rc.2
)

const (
	// EthMsgSoftCapSize is the ideal size of encoded transaction bytes we send in
	// any [EthTxs] or [AtomicTx] message. We do not limit inbound messages to
	// this size, however. Max inbound message size is enforced by the codec
	// (512KB).
	EthMsgSoftCapSize = common.StorageSize(64 * units.KiB)
<<<<<<< HEAD
=======
	atomicTxType      = "atomic-tx"
	ethTxsType        = "eth-txs"
>>>>>>> upstream-v0.8.5-rc.2
)

var (
	_ Message = &AtomicTx{}
	_ Message = &EthTxs{}

	errUnexpectedCodecVersion = errors.New("unexpected codec version")
)

type Message interface {
	// Handle this message with the correct message handler
<<<<<<< HEAD
	Handle(handler Handler, nodeID ids.ShortID, requestID uint32) error
=======
	Handle(handler GossipHandler, nodeID ids.ShortID) error
>>>>>>> upstream-v0.8.5-rc.2

	// initialize should be called whenever a message is built or parsed
	initialize([]byte)

	// Bytes returns the binary representation of this message
	//
	// Bytes should only be called after being initialized
	Bytes() []byte
<<<<<<< HEAD
=======

	// Type returns user-friendly name for this object that can be used for logging
	Type() string
>>>>>>> upstream-v0.8.5-rc.2
}

type message []byte

func (m *message) initialize(bytes []byte) { *m = bytes }
func (m *message) Bytes() []byte           { return *m }

type AtomicTx struct {
	message

	Tx []byte `serialize:"true"`
}

<<<<<<< HEAD
func (msg *AtomicTx) Handle(handler Handler, nodeID ids.ShortID, requestID uint32) error {
	return handler.HandleAtomicTx(nodeID, requestID, msg)
=======
func (msg *AtomicTx) Handle(handler GossipHandler, nodeID ids.ShortID) error {
	return handler.HandleAtomicTx(nodeID, msg)
}

func (msg *AtomicTx) Type() string {
	return atomicTxType
>>>>>>> upstream-v0.8.5-rc.2
}

type EthTxs struct {
	message

	Txs []byte `serialize:"true"`
}

<<<<<<< HEAD
func (msg *EthTxs) Handle(handler Handler, nodeID ids.ShortID, requestID uint32) error {
	return handler.HandleEthTxs(nodeID, requestID, msg)
}

func Parse(bytes []byte) (Message, error) {
	var msg Message
	version, err := c.Unmarshal(bytes, &msg)
	if err != nil {
		return nil, err
	}
	if version != codecVersion {
=======
func (msg *EthTxs) Handle(handler GossipHandler, nodeID ids.ShortID) error {
	return handler.HandleEthTxs(nodeID, msg)
}

func (msg *EthTxs) Type() string {
	return ethTxsType
}

func ParseMessage(codec codec.Manager, bytes []byte) (Message, error) {
	var msg Message
	version, err := codec.Unmarshal(bytes, &msg)
	if err != nil {
		return nil, err
	}
	if version != Version {
>>>>>>> upstream-v0.8.5-rc.2
		return nil, errUnexpectedCodecVersion
	}
	msg.initialize(bytes)
	return msg, nil
}

<<<<<<< HEAD
func Build(msg Message) ([]byte, error) {
	bytes, err := c.Marshal(codecVersion, &msg)
=======
func BuildMessage(codec codec.Manager, msg Message) ([]byte, error) {
	bytes, err := codec.Marshal(Version, &msg)
>>>>>>> upstream-v0.8.5-rc.2
	msg.initialize(bytes)
	return bytes, err
}
