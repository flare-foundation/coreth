// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
)

var (
	errFTSONotDeployed = errors.New("FTSO not deployed")
	errFTSONotActive   = errors.New("FTSO not active")
)
