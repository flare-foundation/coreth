package bitcoin

import (
	"fmt"

	"gitlab.com/flarenetwork/coreth/flare"
	api "gitlab.com/flarenetwork/coreth/flare/api/ripple"
	"gitlab.com/flarenetwork/coreth/flare/connector/ripple"
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
