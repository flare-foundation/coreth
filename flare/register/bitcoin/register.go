package bitcoin

import (
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare"
	api "gitlab.com/flarenetwork/coreth/flare/api/bitcoin"
	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

func init() {

	client, err := api.NewClient()
	if err != nil {
		panic(fmt.Sprintf("could not register bitcoin state connector: %s", err))
	}

	connector := bitcoin.NewConnector(client,
		bitcoin.WithCurrency(flare.CurrencyBitcoin),
	)

	flare.Register(flare.ChainBitcoin, connector)
}
