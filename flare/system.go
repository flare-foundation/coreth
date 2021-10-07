package flare

import (
	"github.com/rs/zerolog"
)

type System struct {
	log        zerolog.Logger
	store      Store
	connectors map[uint32]Connector
}

func NewSystem(log zerolog.Logger, store Store) *System {

	s := System{
		log:        log,
		store:      store,
		connectors: make(map[uint32]Connector),
	}

	return &s
}
