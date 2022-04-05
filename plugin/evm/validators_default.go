package evm

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"

	"github.com/flare-foundation/coreth/params"
)

const (
	costonValidatorWeight   = 200_000
	songbirdValidatorWeight = 50_000
	flareValidatorWeight    = 50_000
	customValidatorWeight   = 200_000
)

var costonNodeIDs = []string{
	"NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc",
	"NodeID-AQghDJTU3zuQj73itPtfTZz6CxsTQVD3R",
	"NodeID-EkH8wyEshzEQBToAdR7Fexxcj9rrmEEHZ",
	"NodeID-FPAwqHjs8Mw8Cuki5bkm3vSVisZr8t2Lu",
	"NodeID-HaZ4HpanjndqSuN252chFsTysmdND5meA",
}

var songbirdNodeIDs = []string{
	"NodeID-3M9KVT6ixi4gVMisbm5TnPXYXgFN5LHuv",
	"NodeID-4tStYRTi3KDxFmv1YHTZAQxbzeyMA7z52",
	"NodeID-7meEpyjmGbL577th58dm4nvvtVZiJusFp",
	"NodeID-8XnMh17zo6pB8Pa2zptRBi9TbbMZgij2t",
	"NodeID-9bWz6J61B8WbQtzeSyA1jsXosyVbuUJd1",
	"NodeID-9SqDo3MxpvEDN4bE4rLTyM7HkkKAw4h96",
	"NodeID-AzdF8JNU468uwZYGquHt7bhDrsggZpK67",
	"NodeID-B9HuZ5hDkRodyRRsiMEHWgMmmMF7xSKbj",
	"NodeID-Cn9P5wgg7d9RNLqm4dFLCUV2diCxpkj7f",
	"NodeID-DLMnewsEwtSH8Qk7p9RGzUVyZAaZVMKsk",
	"NodeID-Fdwp9Wtjh5rxzuTCF9z4zrSM31y7ZzBQS",
	"NodeID-FnvWuwvJGezs4uaBLujkfeM8U3gmAUY3Z",
	"NodeID-FqeGcnLAXbDTthd382aP9uyu1i47paRRh",
	"NodeID-JdEBRLS98PansyFKQUzFKqk4xqrVZ41nC",
	"NodeID-JeYnnrUkuArAAe2Sjo47Z3X5yfeF7cw43",
	"NodeID-Jx3E1F7mfkseZmqnFgDUFV3eusMxVdT6Z",
	"NodeID-LhVs6hzHjBcEkzA1Eu8Qxb9nEQAk1Qbgf",
	"NodeID-NnX4fajAmyvpL9RLfheNdc47FKKDuQW8i",
	"NodeID-PEDdah7g7Efiii1xw8ex2dH58oMfByzjb",
	"NodeID-QCt9AxMPt5nn445CQGoA3yktqkChnKmPY",
}

var flareNodeIDs = []string{}

type ValidatorsDefault struct {
	validators map[ids.ShortID]uint64
	steps      []Step
}

type Step struct {
	Epoch  uint64
	Cutoff int
}

func NewValidatorsDefault(chainID *big.Int) (*ValidatorsDefault, error) {

	var weight uint64
	var nodeIDs []string
	var steps []Step
	switch {
	case chainID.Cmp(params.CostonChainID) == 0:
		nodeIDs = costonNodeIDs
		weight = costonValidatorWeight
		steps = []Step{
			{Epoch: 1604, Cutoff: 4}, // go down to 4 default validators one week after hard fork
			{Epoch: 1772, Cutoff: 3}, // go down to 3 default validators two weeks after hard fork
		}
	case chainID.Cmp(params.SongbirdChainID) == 0:
		nodeIDs = songbirdNodeIDs
		weight = songbirdValidatorWeight
		steps = []Step{
			{Epoch: 42, Cutoff: 15}, // go down to 15 default validators two weeks after main net launch
			{Epoch: 44, Cutoff: 10}, // go down to 10 default validators four weeks after main net launch
			{Epoch: 46, Cutoff: 5},  // go down to 5 default validators six weeks after main net launch
		}
	case chainID.Cmp(params.FlareChainID) == 0:
		nodeIDs = flareNodeIDs
		weight = flareValidatorWeight
	default:
		customValidators := os.Getenv("CUSTOM_VALIDATORS")
		if customValidators == "" {
			return nil, fmt.Errorf("custom validators not set for non-standard network (chain: %s)", chainID)
		}
		nodeIDs = strings.Split(customValidators, ",")
		weight = customValidatorWeight
	}

	if len(nodeIDs) == 0 {
		return nil, fmt.Errorf("no default validators set")
	}

	validators := make(map[ids.ShortID]uint64, len(nodeIDs))
	for _, nodeID := range nodeIDs {
		validator, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			return nil, fmt.Errorf("could not parse validator (nodeid: %s): %w", nodeID, err)
		}
		validators[validator] = weight
	}

	v := ValidatorsDefault{
		validators: validators,
		steps:      steps,
	}

	return &v, nil
}

func (v *ValidatorsDefault) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	validatorIDs := make([]ids.ShortID, 0, len(v.validators))
	for validatorID := range v.validators {
		validatorIDs = append(validatorIDs, validatorID)
	}
	sort.Slice(validatorIDs, func(i int, j int) bool {
		return bytes.Compare(validatorIDs[i][:], validatorIDs[j][:]) < 0
	})

	for i := len(v.steps) - 1; i >= 0; i-- {
		step := v.steps[i]
		if epoch >= step.Epoch {
			validatorIDs = validatorIDs[:step.Cutoff]
			break
		}
	}

	reduced := make(map[ids.ShortID]uint64, len(validatorIDs))
	for _, validatorID := range validatorIDs {
		reduced[validatorID] = v.validators[validatorID]
	}

	if len(reduced) == 0 {
		return nil, fmt.Errorf("no default validators available for epoch")
	}

	return reduced, nil
}
