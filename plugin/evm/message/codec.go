<<<<<<< HEAD
// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
=======
// (c) 2019-2022, Ava Labs, Inc. All rights reserved.
>>>>>>> upstream-v0.8.5-rc.2
// See the file LICENSE for licensing terms.

package message

import (
<<<<<<< HEAD
	"github.com/flare-foundation/flare/codec"
	"github.com/flare-foundation/flare/codec/linearcodec"
	"github.com/flare-foundation/flare/codec/reflectcodec"
	"github.com/flare-foundation/flare/utils/units"
	"github.com/flare-foundation/flare/utils/wrappers"
)

const (
	codecVersion   uint16 = 0
	maxMessageSize        = 512 * units.KiB
	maxSliceLen           = maxMessageSize
)

// Codec does serialization and deserialization
var c codec.Manager

func init() {
	c = codec.NewManager(maxMessageSize)
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSliceLen)

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&AtomicTx{}),
		lc.RegisterType(&EthTxs{}),
		c.RegisterCodec(codecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
=======
	"github.com/flare-foundation/flare/codec"
	"github.com/flare-foundation/flare/codec/linearcodec"
	"github.com/flare-foundation/flare/utils/units"
	"github.com/flare-foundation/flare/utils/wrappers"
)

const Version = uint16(0)
const maxMessageSize = 1 * units.MiB

func BuildCodec() (codec.Manager, error) {
	codecManager := codec.NewManager(maxMessageSize)
	c := linearcodec.NewDefault()
	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&AtomicTx{}),
		c.RegisterType(&EthTxs{}),
	)
	errs.Add(codecManager.RegisterCodec(Version, c))
	return codecManager, errs.Err
>>>>>>> upstream-v0.8.5-rc.2
}
