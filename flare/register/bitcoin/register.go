package bitcoin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/flare-foundation/coreth/flare"
	api "github.com/flare-foundation/coreth/flare/client/bitcoin"
	"github.com/flare-foundation/coreth/flare/client/rpc"
	"github.com/flare-foundation/coreth/flare/connector/bitcoin"
	"github.com/flare-foundation/coreth/flare/connector/multi"
)

func init() {

	log := flare.Log()

	apis := os.Getenv(flare.EndpointsBitcoin)
	endpoints := strings.Split(apis, ",")
	if len(endpoints) == 0 {
		log.Fatal().Msg("no Bitcoin API endpoints configured")
	}

	matches := uint(len(endpoints)/2 + 1)

	connectors := make([]flare.Connector, 0, len(endpoints))
	for _, endpoint := range endpoints {

		hash := sha256.Sum256([]byte(endpoint))
		checksum := hex.EncodeToString(hash[0:4])
		username := os.Getenv(flare.PrefixUsernameBitcoin + checksum)
		password := os.Getenv(flare.PrefixPasswordBitcoin + checksum)

		opts := jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": fmt.Sprintf("Basic %s:%s", username, password),
			},
		}
		rpcClient := jsonrpc.NewClientWithOpts(endpoint, &opts)

		retryClient := rpc.NewRetryClient(rpcClient,
			rpc.WithMaxElapsed(30*time.Second),
		)

		apiClient := api.NewAPIClient(retryClient)

		connector := bitcoin.NewConnector(apiClient,
			bitcoin.WithCurrency(flare.CurrencyBitcoin),
		)

		connectors = append(connectors, connector)
	}

	multiConnector, err := multi.NewConnector(log, connectors,
		multi.WithMatchesRequired(matches),
	)
	if err != nil {
		panic(fmt.Sprintf("could not initialize Bitcoin connectors: %s", err))
	}

	flare.Register(flare.ChainBitcoin, multiConnector)
}
