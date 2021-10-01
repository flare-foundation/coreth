package bitcoin

import (
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare"
	api "gitlab.com/flarenetwork/coreth/flare/api/bitcoin"
	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

func init() {

	cli, err := api.NewClient(
	// api.WithHost(""),
	// api.WithUser(""),
	// api.WithPassword(""),
	// api.WithSecure(true),
	)
	if err != nil {
		panic(fmt.Sprintf("could not register bitcoin state connector: %s", err))
	}

	connector := bitcoin.NewConnector(cli,
		bitcoin.WithCurrency(flare.CurrencyBitcoin),
	)

	flare.Register(flare.ChainBitcoin, connector)
}
