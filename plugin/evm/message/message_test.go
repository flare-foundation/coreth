// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"testing"

<<<<<<< HEAD
	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/utils"
	"github.com/flare-foundation/flare/utils/units"
=======
	"github.com/flare-foundation/flare/utils"
	"github.com/flare-foundation/flare/utils/units"

	"github.com/stretchr/testify/assert"
>>>>>>> upstream-v0.8.5-rc.2
)

func TestAtomicTx(t *testing.T) {
	assert := assert.New(t)

	msg := []byte("blah")
	builtMsg := AtomicTx{
		Tx: msg,
	}
<<<<<<< HEAD
	builtMsgBytes, err := Build(&builtMsg)
	assert.NoError(err)
	assert.Equal(builtMsgBytes, builtMsg.Bytes())

	parsedMsgIntf, err := Parse(builtMsgBytes)
=======
	codec, err := BuildCodec()
	assert.NoError(err)
	builtMsgBytes, err := BuildMessage(codec, &builtMsg)
	assert.NoError(err)
	assert.Equal(builtMsgBytes, builtMsg.Bytes())

	parsedMsgIntf, err := ParseMessage(codec, builtMsgBytes)
>>>>>>> upstream-v0.8.5-rc.2
	assert.NoError(err)
	assert.Equal(builtMsgBytes, parsedMsgIntf.Bytes())

	parsedMsg, ok := parsedMsgIntf.(*AtomicTx)
	assert.True(ok)

	assert.Equal(msg, parsedMsg.Tx)
}

func TestEthTxs(t *testing.T) {
	assert := assert.New(t)

	msg := []byte("blah")
	builtMsg := EthTxs{
		Txs: msg,
	}
<<<<<<< HEAD
	builtMsgBytes, err := Build(&builtMsg)
	assert.NoError(err)
	assert.Equal(builtMsgBytes, builtMsg.Bytes())

	parsedMsgIntf, err := Parse(builtMsgBytes)
=======
	codec, err := BuildCodec()
	assert.NoError(err)
	builtMsgBytes, err := BuildMessage(codec, &builtMsg)
	assert.NoError(err)
	assert.Equal(builtMsgBytes, builtMsg.Bytes())

	parsedMsgIntf, err := ParseMessage(codec, builtMsgBytes)
>>>>>>> upstream-v0.8.5-rc.2
	assert.NoError(err)
	assert.Equal(builtMsgBytes, parsedMsgIntf.Bytes())

	parsedMsg, ok := parsedMsgIntf.(*EthTxs)
	assert.True(ok)

	assert.Equal(msg, parsedMsg.Txs)
}

func TestEthTxsTooLarge(t *testing.T) {
	assert := assert.New(t)

	builtMsg := EthTxs{
		Txs: utils.RandomBytes(1024 * units.KiB),
	}
<<<<<<< HEAD
	_, err := Build(&builtMsg)
=======
	codec, err := BuildCodec()
	assert.NoError(err)
	_, err = BuildMessage(codec, &builtMsg)
>>>>>>> upstream-v0.8.5-rc.2
	assert.Error(err)
}

func TestParseGibberish(t *testing.T) {
	assert := assert.New(t)

<<<<<<< HEAD
	randomBytes := utils.RandomBytes(256 * units.KiB)
	_, err := Parse(randomBytes)
=======
	codec, err := BuildCodec()
	assert.NoError(err)
	randomBytes := utils.RandomBytes(256 * units.KiB)
	_, err = ParseMessage(codec, randomBytes)
>>>>>>> upstream-v0.8.5-rc.2
	assert.Error(err)
}
