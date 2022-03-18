// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"
)

var DefaultFTSOConfig = FTSOConfig{
	RootDegree: 4,
}

type FTSOConfig struct {
	RootDegree uint
}

type FTSOOption func(*FTSOConfig)

func WithRootDegree(degree uint) FTSOOption {
	return func(cfg *FTSOConfig) {
		cfg.RootDegree = degree
	}
}

type FTSO interface {
	Details(epoch uint64) (EpochDetails, error)
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
	ftso FTSO
	cfg  FTSOConfig
}

func NewValidatorsFTSO(ftso FTSO, opts ...FTSOOption) *ValidatorsFTSO {

	cfg := DefaultFTSOConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	v := ValidatorsFTSO{
		ftso: ftso,
		cfg:  cfg,
	}

	return &v
}

func (v *ValidatorsFTSO) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	snap, err := v.ftso.Snapshot(epoch)
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO snapshot: %w", err)
	}

	providers, err := snap.Providers()
	if err != nil {
		return nil, fmt.Errorf("could not get FTSO providers: %w", err)
	}

	validators := make(map[ids.ShortID]uint64, len(providers))
	for _, provider := range providers {

		id, err := snap.Validator(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get validator  (provider: %x): %w", provider, err)
		}

		votepower, err := snap.Votepower(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get vote power (provider: %x): %w", provider, err)
		}

		rewards, err := snap.Rewards(provider)
		if err != nil {
			return nil, fmt.Errorf("could not get rewards (provider: %x): %w", provider, err)
		}

		weight := uint64(math.Pow(votepower, 1.0/float64(v.cfg.RootDegree)) * (rewards / votepower))

		validators[id] = weight
	}

	return validators, nil
}
