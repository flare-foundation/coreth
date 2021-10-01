package bitcoin

import (
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare"
	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin/api"
)

func init() {

	cli, err := api.NewClient(
	// api.WithHost(""),
	// api.WithUser(""),
	// api.WithPassword(""),
	// api.WithSecure(true),
	)
	if err != nil {
		panic(fmt.Sprintf("could not register litecoin state connector: %s", err))
	}

	connector := bitcoin.NewConnector(cli,
		bitcoin.WithCurrency(flare.CurrencyLitecoin),
	)

	flare.Register(flare.ChainLitecoin, connector)
}
