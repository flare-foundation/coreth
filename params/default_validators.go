package params

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils/constants"
)

var (
	testingSet  validation.Set
	costonSet   validation.Set
	songbirdSet validation.Set
	customSet   validation.Set
)

var testingNodeIDs = []string{
	"NodeID-MEHBQFqQnSz7KzS8u4t8nWy7fSaqN2Pdp",
}

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

func init() {

	var err error

	testingSet, err = buildValidators(testingNodeIDs, 1_000_000)
	if err != nil {
		panic(fmt.Sprintf("invalid testing default validators: %s", err))
	}

	costonSet, err = buildValidators(costonNodeIDs, 200_000)
	if err != nil {
		panic(fmt.Sprintf("invalid coston default validators: %s", err))
	}

	songbirdSet, err = buildValidators(songbirdNodeIDs, 50_000)
	if err != nil {
		panic(fmt.Sprintf("invalid songbird default validators: %s", err))
	}

	customNodeIDList := os.Getenv("CUSTOM_DEFAULT_VALIDATORS")
	customNodeIDs := strings.Split(customNodeIDList, ",")
	customSet, err = buildValidators(customNodeIDs, 500_000)
	if err != nil {
		panic(fmt.Sprintf("invalid custom default validators: %s", err))
	}
}

func DefaultValidators(chainID *big.Int) validation.Set {
	switch chainID {
	case TestingChainID:
		return testingSet
	case CostonChainID:
		return costonSet
	case SongbirdChainID:
		return songbirdSet
	}
	if customSet.Len() == 0 {
		panic("missing custom default validators for non-standard network, please set CUSTOM_DEFAULT_VALIDATORS")
	}
	return customSet
}

func buildValidators(nodeIDs []string, weight uint64) (validation.Set, error) {

	set := validation.NewSet()

	for _, nodeID := range nodeIDs {

		if nodeID == "" {
			continue
		}

		parsedID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			return nil, fmt.Errorf("invalid custom default validator for non-standard network (%s)", nodeID)
		}

		err = set.AddWeight(parsedID, weight)
		if err != nil {
			return nil, fmt.Errorf("could not add weight for custom default validator: %w", err)
		}
	}

	return set, nil
}
