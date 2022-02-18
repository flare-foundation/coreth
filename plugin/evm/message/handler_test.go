// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"testing"

<<<<<<< HEAD
	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
=======
	"github.com/flare-foundation/flare/ids"

	"github.com/stretchr/testify/assert"
>>>>>>> upstream-v0.8.5-rc.2
)

type CounterHandler struct {
	AtomicTx, EthTxs int
}

<<<<<<< HEAD
func (h *CounterHandler) HandleAtomicTx(ids.ShortID, uint32, *AtomicTx) error {
=======
func (h *CounterHandler) HandleAtomicTx(ids.ShortID, *AtomicTx) error {
>>>>>>> upstream-v0.8.5-rc.2
	h.AtomicTx++
	return nil
}

<<<<<<< HEAD
func (h *CounterHandler) HandleEthTxs(ids.ShortID, uint32, *EthTxs) error {
=======
func (h *CounterHandler) HandleEthTxs(ids.ShortID, *EthTxs) error {
>>>>>>> upstream-v0.8.5-rc.2
	h.EthTxs++
	return nil
}

func TestHandleAtomicTx(t *testing.T) {
	assert := assert.New(t)

	handler := CounterHandler{}
	msg := AtomicTx{}

<<<<<<< HEAD
	err := msg.Handle(&handler, ids.ShortEmpty, 0)
=======
	err := msg.Handle(&handler, ids.ShortEmpty)
>>>>>>> upstream-v0.8.5-rc.2
	assert.NoError(err)
	assert.Equal(1, handler.AtomicTx)
	assert.Zero(handler.EthTxs)
}

func TestHandleEthTxs(t *testing.T) {
	assert := assert.New(t)

	handler := CounterHandler{}
	msg := EthTxs{}

<<<<<<< HEAD
	err := msg.Handle(&handler, ids.ShortEmpty, 0)
=======
	err := msg.Handle(&handler, ids.ShortEmpty)
>>>>>>> upstream-v0.8.5-rc.2
	assert.NoError(err)
	assert.Zero(handler.AtomicTx)
	assert.Equal(1, handler.EthTxs)
}

func TestNoopHandler(t *testing.T) {
	assert := assert.New(t)

<<<<<<< HEAD
	handler := NoopHandler{}

	err := handler.HandleAtomicTx(ids.ShortEmpty, 0, nil)
	assert.NoError(err)

	err = handler.HandleEthTxs(ids.ShortEmpty, 0, nil)
=======
	handler := NoopMempoolGossipHandler{}

	err := handler.HandleAtomicTx(ids.ShortEmpty, nil)
	assert.NoError(err)

	err = handler.HandleEthTxs(ids.ShortEmpty, nil)
>>>>>>> upstream-v0.8.5-rc.2
	assert.NoError(err)
}
