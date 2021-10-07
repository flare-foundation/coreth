package flare

import (
	"os"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/rs/zerolog"

	bstore "gitlab.com/flarenetwork/coreth/flare/store/badger"
)

var system *System

// TODO: Make these constants configuration parameters.
const (
	logLevel = zerolog.DebugLevel
	dbDir    = "cache"
)

func init() {

	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(logLevel)

	opts := badger.DefaultOptions(dbDir)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal().Err(err).Msg("could not open database")
	}

	store := bstore.NewStore(db)
	system = NewSystem(log, store)
}
