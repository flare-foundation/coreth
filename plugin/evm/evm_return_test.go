// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
)

func TestEVMReturn_Decode(t *testing.T) {

	inInt := big.NewInt(1)
	inInts := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	inBytes := []byte{1, 2, 3}
	inBytes20 := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	inAddress := common.HexToAddress("0xE93012B9d91b3C8F7C81eb1A4B92AaE06698DDda")
	inAddresses := []common.Address{{1}, {2}, {3}}

	ret := &EVMReturn{values: []interface{}{
		inInt,
		inInts,
		inBytes,
		inBytes20,
		inAddress,
		inAddresses,
	}}

	outInt := big.NewInt(0)
	outInts := []*big.Int{}
	outBytes := []byte{}
	outBytes20 := [20]byte{}
	outAddress := common.Address{}
	outAddresses := []common.Address{}
	err := ret.Decode(&outInt, &outInts, &outBytes, &outBytes20, &outAddress, &outAddresses)
	require.NoError(t, err)

	assert.Equal(t, inInt, outInt)
	assert.Equal(t, inInts, outInts)
	assert.Equal(t, inBytes, outBytes)
	assert.Equal(t, inBytes20, outBytes20)
	assert.Equal(t, inAddress, outAddress)
	assert.Equal(t, inAddresses, outAddresses)
}
