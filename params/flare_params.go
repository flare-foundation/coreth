// (c) 2019-2020, Ava Labs, Inc.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare/ids"
)

var costonParams = FlareParams{
	flareDaemon:       common.HexToAddress("0x1000000000000000000000000000000000000001"),
	stateConnector:    common.HexToAddress("0x1000000000000000000000000000000000000002"),
	priceSubmitter:    common.HexToAddress("0x1000000000000000000000000000000000000003"),
	validatorRegistry: common.HexToAddress("0x1000000000000000000000000000000000000004"),
	defaultValidators: []ids.ShortID{},
	rootDegree:        1,
	stepSize:          1,
}

var songbirdParams = FlareParams{
	flareDaemon:       common.HexToAddress("0x1000000000000000000000000000000000000001"),
	stateConnector:    common.HexToAddress("0x1000000000000000000000000000000000000002"),
	priceSubmitter:    common.HexToAddress("0x1000000000000000000000000000000000000003"),
	validatorRegistry: common.HexToAddress("0x1000000000000000000000000000000000000004"),
	defaultValidators: []ids.ShortID{},
	rootDegree:        4,
	stepSize:          4,
}

var flareParams = FlareParams{
	flareDaemon:       common.HexToAddress("0x1000000000000000000000000000000000000001"),
	stateConnector:    common.HexToAddress("0x1000000000000000000000000000000000000002"),
	priceSubmitter:    common.HexToAddress("0x1000000000000000000000000000000000000003"),
	validatorRegistry: common.HexToAddress("0x1000000000000000000000000000000000000004"),
	defaultValidators: []ids.ShortID{},
	rootDegree:        4,
	stepSize:          2,
}

type FlareParams struct {
	flareDaemon       common.Address
	stateConnector    common.Address
	priceSubmitter    common.Address
	validatorRegistry common.Address
	defaultValidators []ids.ShortID
	rootDegree        uint
	stepSize          uint
}

func LoadFlareParams(chainID *big.Int) (FlareParams, error) {
	switch {
	case chainID.Cmp(CostonChainID) == 0:
		return costonParams, nil
	case chainID.Cmp(SongbirdChainID) == 0:
		return songbirdParams, nil
	case chainID.Cmp(FlareChainID) == 0:
		return flareParams, nil
	default:
		// TODO: this should parse the environment variables and create a new
		// FTOS parmas
		return FlareParams{}, nil
	}
}

func (f FlareParams) FlareDaemon(timestamp uint64) common.Address {
	return f.flareDaemon
}

func (f FlareParams) StateConnector(timestamp uint64) common.Address {
	return f.stateConnector
}

func (f FlareParams) PriceSubmitter(timestamp uint64) common.Address {
	return f.priceSubmitter
}

func (f FlareParams) ValidatorRegistry(timestamp uint64) common.Address {
	return f.validatorRegistry
}

func (f FlareParams) DefaultValidators(timestamp uint64) []ids.ShortID {
	return f.defaultValidators
}

func (f FlareParams) RootDegree(timestamp uint64) uint {
	return f.rootDegree
}

func (f FlareParams) StepSize(timestamp uint64) uint {
	return f.stepSize
}
