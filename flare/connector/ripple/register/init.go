package bitcoin

import (
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare"
	"gitlab.com/flarenetwork/coreth/flare/connector/ripple"
	"gitlab.com/flarenetwork/coreth/flare/connector/ripple/api"
)

func init() {

	cli, err := api.NewClient(
	// api.WithHost(""),
	// api.WithUser(""),
	// api.WithPassword(""),
	// api.WithSecure(true),
	)
	if err != nil {
		panic(fmt.Sprintf("could not register ripple state connector: %s", err))
	}

	connector := ripple.NewConnector(cli,
		ripple.WithCurrency(flare.CurrencyRipple),
	)

	flare.Register(flare.ChainRipple, connector)
}
