package ripple

import (
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

	apis := os.Getenv(flare.EndpointsRipple)
	endpoints := strings.Split(apis, ",")
	if len(endpoints) == 0 {
		log.Fatal().Msg("no Ripple API endpoints configured")
	}

	matches := uint(len(endpoints)/2 + 1)

	connectors := make([]flare.Connector, 0, len(endpoints))
	for _, endpoint := range endpoints {

		rpcClient := jsonrpc.NewClient(endpoint)

		retryClient := rpc.NewRetryClient(rpcClient,
			rpc.WithMaxElapsed(30*time.Second),
		)

		apiClient := api.NewAPIClient(retryClient)

		connector := bitcoin.NewConnector(apiClient,
			bitcoin.WithCurrency(flare.CurrencyRipple),
		)

		connectors = append(connectors, connector)
	}

	multiConnector, err := multi.NewConnector(log, connectors,
		multi.WithMatchesRequired(matches),
	)
	if err != nil {
		panic(fmt.Sprintf("could not initialize Ripple connectors: %s", err))
	}

	flare.Register(flare.ChainRipple, multiConnector)
}
