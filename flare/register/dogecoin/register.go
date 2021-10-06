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
		panic(fmt.Sprintf("could not register dogecoin state connector: %s", err))
	}

	connector := bitcoin.NewConnector(client,
		bitcoin.WithCurrency(flare.CurrencyDogecoin),
	)

	flare.Register(flare.ChainDogecoin, connector)
}
