package multi

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"gitlab.com/flarenetwork/coreth/flare"
)

// Multi is a multi-connector wrapping around a number of other connectors and
// supports verification on n-of-m connectors before returning success on any
// state connector call.
type Multi struct {
	log        zerolog.Logger
	connectors []flare.Connector
	cfg        Config
}

// NewMulti creates a new multi-connector wrapping the given connectors and
// configured with the given n-of-m options. It uses the injected logger to log
// errors that happen _after_ successfully completing a call with some calls on
// underlying connectors still pending.
func NewMulti(log zerolog.Logger, connectors []flare.Connector, options ...Option) (*Multi, error) {

	cfg := DefaultConfig
	for _, option := range options {
		option(&cfg)
	}

	if cfg.MatchesRequired > uint(len(connectors)) {
		return nil, fmt.Errorf("insufficient number of connectors (required_matches: %d, connectors: %d)", cfg.MatchesRequired, len(connectors))
	}

	m := Multi{
		log:        log.With().Str("component", "multi_connector").Logger(),
		connectors: connectors,
		cfg:        cfg,
	}

	return &m, nil
}

// ProveAvailability will execute availability proof calls on the underlying
// state connectors and return success if and only if the configured number of
// connectors return success.
func (m *Multi) ProveAvailability(ret []byte) (bool, error) {
	calls := make([]Call, 0, len(m.connectors))
	for _, connector := range m.connectors {
		calls = append(calls, connector.ProveAvailability)
	}
	return m.execute(ret, calls)
}

// ProvePayment will execute payment proof calls on the underlying state
// connectors and return success if and only if the configured number of
// connectors return success.
func (m *Multi) ProvePayment(ret []byte) (bool, error) {
	calls := make([]Call, 0, len(m.connectors))
	for _, connector := range m.connectors {
		calls = append(calls, connector.ProvePayment)
	}
	return m.execute(ret, calls)
}

// DisprovePayment will execute payment disproof calls on the underlying state
// connectors and return success if and only if the configured number of
// connectors return success.
func (m *Multi) DisprovePayment(ret []byte) (bool, error) {
	calls := make([]Call, 0, len(m.connectors))
	for _, connector := range m.connectors {
		calls = append(calls, connector.DisprovePayment)
	}
	return m.execute(ret, calls)
}

// execute will execute the given state connector calls with the given return
// data. It will wait for state connector call to finish until either the
// configured number of successful calls has been reached, or until all of the
// state connector calls have returned. If the call is unsuccessful, all
// encountered errors are returned. If the call is successful, encountered
// errors are logged so as to allow detection of faulty connectors while still
// allowing the multi-connector to return as soon as possible.
func (m *Multi) execute(ret []byte, calls []Call) (bool, error) {

	results := make(chan bool, len(calls))
	errors := make(chan error, len(calls))

	wg := &sync.WaitGroup{}
	for _, call := range calls {
		wg.Add(1)
		go func(call Call) {
			defer wg.Done()
			ok, err := call(ret)
			if err != nil {
				errors <- err
			}
			results <- ok
		}(call)
	}

	var messages []string
	go func() {
		wg.Wait()
		close(errors)
		for err := range errors {
			m.log.Warn().Err(err).Msg("state connector call failed")
			messages = append(messages, err.Error())
		}
		close(results)
	}()

	accepted := uint(0)
	rejected := uint(0)
	for ok := range results {
		if ok {
			accepted++
		} else {
			rejected++
		}
		if accepted >= m.cfg.MatchesRequired {
			return true, nil
		}
	}

	if len(messages) > 0 {
		err := fmt.Errorf(strings.Join(messages, ", "))
		return false, err
	}

	return false, nil
}
