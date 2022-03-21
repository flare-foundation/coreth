// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

var DefaultFTSOConfig = FTSOConfig{
	RootDegree:      4,
	RatioMultiplier: 100.0,
}

type FTSOConfig struct {
	RootDegree      uint
	RatioMultiplier float64
}

type FTSOOption func(*FTSOConfig)

func WithRootDegree(degree uint) FTSOOption {
	return func(cfg *FTSOConfig) {
		cfg.RootDegree = degree
	}
}

func WithRatioMultiplier(multiplier float64) FTSOOption {
	return func(cfg *FTSOConfig) {
		cfg.RatioMultiplier = multiplier
	}
}

type FTSO interface {
	Current(hash common.Hash) (uint64, error)
	Details(epoch uint64) (FTSOEpoch, error)
	Snapshot(epoch uint64) (Snapshot, error)
}

type Snapshot interface {
	Providers() ([]common.Address, error)
	Validator(provider common.Address) (ids.ShortID, error)
	Votepower(provider common.Address) (float64, error)
	Rewards(provider common.Address) (float64, error)
}

// ValidatorsFTSO is responsible for retrieving the set of validators for the FTSO
// data providers, in accordance with the defined formula and configured root degree.
type ValidatorsFTSO struct {
	log        logging.Logger
	blockchain *core.BlockChain
	ftso       FTSO
	cfg        FTSOConfig
}

func NewValidatorsFTSO(log logging.Logger, blockchain *core.BlockChain, ftso FTSO, opts ...FTSOOption) *ValidatorsFTSO {

	cfg := DefaultFTSOConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	v := ValidatorsFTSO{
		log:        log,
		blockchain: blockchain,
		ftso:       ftso,
		cfg:        cfg,
	}

	return &v
}

func (v *ValidatorsFTSO) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	validators := make(map[ids.ShortID]uint64)

	snap, err := v.ftso.Snapshot(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO snapshot: %w", err)
	}

	providers, err := snap.Providers()
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO providers: %w", err)
	}

	for _, provider := range providers {

		validator, err := snap.Validator(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get FTSO validator (provider: %s): %w", provider, err)
		}
		if validator == ids.ShortEmpty {
			v.log.Debug("skipping provider %s with unset validator", provider.Hex())
			continue
		}

		votepower, err := snap.Votepower(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get vote power (provider: %s): %w", provider, err)
		}
		if votepower == 0 {
			v.log.Debug("skipping provider %s with validator %s and no votepower", provider.Hex(), validator)
			continue
		}

		rewards, err := snap.Rewards(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get rewards (provider: %s): %w", provider, err)
		}
		if rewards == 0 {
			v.log.Debug("skipping provider %s with validator %s and no rewards", provider.Hex(), validator)
			continue
		}

		weight := uint64(math.Pow(votepower, 1.0/float64(v.cfg.RootDegree)) * (v.cfg.RatioMultiplier * rewards / votepower))

		v.log.Debug("pro:%s val:%s vp:%f rw:%f w:%d", provider.Hex(), validator, votepower, rewards, weight)

		validators[validator] = weight
	}

	return validators, nil
}
