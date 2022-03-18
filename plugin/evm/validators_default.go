package evm

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

var costonNodeIDs = []string{
	"NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc",
	"NodeID-EkH8wyEshzEQBToAdR7Fexxcj9rrmEEHZ",
	"NodeID-FPAwqHjs8Mw8Cuki5bkm3vSVisZr8t2Lu",
	"NodeID-AQghDJTU3zuQj73itPtfTZz6CxsTQVD3R",
	"NodeID-HaZ4HpanjndqSuN252chFsTysmdND5meA",
}

var songbirdNodeIDs = []string{
	"NodeID-3M9KVT6ixi4gVMisbm5TnPXYXgFN5LHuv",
	"NodeID-NnX4fajAmyvpL9RLfheNdc47FKKDuQW8i",
	"NodeID-AzdF8JNU468uwZYGquHt7bhDrsggZpK67",
	"NodeID-FqeGcnLAXbDTthd382aP9uyu1i47paRRh",
	"NodeID-B9HuZ5hDkRodyRRsiMEHWgMmmMF7xSKbj",
	"NodeID-Jx3E1F7mfkseZmqnFgDUFV3eusMxVdT6Z",
	"NodeID-FnvWuwvJGezs4uaBLujkfeM8U3gmAUY3Z",
	"NodeID-LhVs6hzHjBcEkzA1Eu8Qxb9nEQAk1Qbgf",
	"NodeID-9SqDo3MxpvEDN4bE4rLTyM7HkkKAw4h96",
	"NodeID-4tStYRTi3KDxFmv1YHTZAQxbzeyMA7z52",
	"NodeID-8XnMh17zo6pB8Pa2zptRBi9TbbMZgij2t",
	"NodeID-Cn9P5wgg7d9RNLqm4dFLCUV2diCxpkj7f",
	"NodeID-PEDdah7g7Efiii1xw8ex2dH58oMfByzjb",
	"NodeID-QCt9AxMPt5nn445CQGoA3yktqkChnKmPY",
	"NodeID-9bWz6J61B8WbQtzeSyA1jsXosyVbuUJd1",
	"NodeID-DLMnewsEwtSH8Qk7p9RGzUVyZAaZVMKsk",
	"NodeID-7meEpyjmGbL577th58dm4nvvtVZiJusFp",
	"NodeID-JeYnnrUkuArAAe2Sjo47Z3X5yfeF7cw43",
	"NodeID-Fdwp9Wtjh5rxzuTCF9z4zrSM31y7ZzBQS",
	"NodeID-JdEBRLS98PansyFKQUzFKqk4xqrVZ41nC",
}

var flareNodeIDs = []string{}

func getDefaultValidators(chainID *big.Int) ([]ids.ShortID, error) {

	var nodeIDs []string
	switch {
	case chainID.Cmp(params.CostonChainID) == 0:
		nodeIDs = costonNodeIDs
	case chainID.Cmp(params.SongbirdChainID) == 0:
		nodeIDs = songbirdNodeIDs
	case chainID.Cmp(params.FlareChainID) == 0:
		nodeIDs = flareNodeIDs
	default:
		nodeIDs = strings.Split(os.Getenv("DEFAULT_VALIDATORS"), ",")
	}

	if len(nodeIDs) == 0 {
		return nil, fmt.Errorf("no default validators set")
	}

	validators := make([]ids.ShortID, 0, len(nodeIDs))
	for _, nodeID := range nodeIDs {
		validator, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			return nil, fmt.Errorf("could not parse validator (nodeid: %s): %w", nodeID, err)
		}
		validators = append(validators, validator)
	}

	return validators, nil
}
