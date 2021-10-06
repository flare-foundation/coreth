package bitcoin

import (
	"fmt"
	"math"
	"sync"

	"github.com/rs/zerolog"

	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

// MultiClient is a wrapper around multiple Bitcoin API clients, which can use
// the responses from all clients to improve security and responsiveness of
// executed API requests.
type MultiClient struct {
	log     zerolog.Logger
	clients []bitcoin.API
	cfg     MultiConfig
}

// NewMultiClient creates a new wrapper around the given Bitcoin API clients,
// and will use the given options to determine when a request should be
// considered successful.
func NewMultiClient(log zerolog.Logger, clients []bitcoin.API, options ...MultiOption) (*MultiClient, error) {

	cfg := DefaultMultiConfig
	for _, option := range options {
		option(&cfg)
	}

	if uint(len(clients)) > cfg.RequiredMatches {
		return nil, fmt.Errorf("insufficient number of clients for required matches (clients: %d, matches: %d)", len(clients), cfg.RequiredMatches)
	}

	m := MultiClient{
		log:     log.With().Str("component", "bitcoin_multi_client").Logger(),
		clients: clients,
		cfg:     cfg,
	}

	return &m, nil
}

// Block retrieves a block with the given hash from the underlying Bitcoin API
// clients. It will only return the block once the configured number of clients
// agree on its contents.
func (m *MultiClient) Block(hash [32]byte) (*bitcoin.Block, error) {

	// In the first step, we prepare a channel to handle request results
	// (blocks), a channel to handle request errors, and a channel to
	// signal when all requests have finished.
	results := make(chan *bitcoin.Block, len(m.clients))
	errors := make(chan error, len(m.clients))
	done := make(chan struct{})
	wg := &sync.WaitGroup{}

	// Once the channels are in place, we spin off one goroutine for each
	// underlying client. We use the wait group to keep track of when the
	// requests on all clients are done.
	for _, client := range m.clients {
		wg.Add(1)
		go func(client bitcoin.API) {
			defer wg.Done()
			block, err := client.Block(hash)
			if err != nil {
				errors <- err
				return
			}
			results <- block
		}(client)
	}

	// Here, we spin off a goroutine to handle clean up after all requests are
	// done. It waits on the wait group and then closes and drains all channels.
	// It is also responsible for logging all request errors, just in case the
	// node operator wants to review the nodes he relies on.
	go func() {
		wg.Wait()
		close(results)
		close(errors)
		for range results {
		}
		for err := range errors {
			m.log.Warn().Err(err).Msg("could not get block from client")
		}
		close(done)
	}()

	// The main logic of the multi-client happens in the following section. We
	// wait for results (blocks) to come in on the results channel. Each
	// time we see a result, we check whether it corresponds to the hash and
	// output index we specified and keep track of it in a list of results with
	// the same contents (height). Once we receive the  configured amount of
	// blocks with same contents, we consider these contents the correct ones
	// and proceed to the next sections.
	// NOTE: It's possible that we never receive enough results with the same
	// contents. Once all requests are done, the cleanup goroutine sends the
	// done signal and we can return the corresponding error, and we thus avoid
	// getting stuck in an infinite loop.
	var selected *bitcoin.Block
	candidates := make(map[uint64][]*bitcoin.Block)
RequestLoop:
	for {
		select {
		case <-done:
			return nil, fmt.Errorf("insufficient number of matching blocks")
		case block := <-results:
			if block.Hash != hash {
				m.log.Warn().
					Hex("want", block.Hash[:]).
					Hex("have", hash[:]).
					Msg("block hash does not match requested hash")
				continue RequestLoop
			}
			candidates[block.Height] = append(candidates[block.Height], block)
			if uint(len(candidates[block.Height])) >= m.cfg.RequiredMatches {
				selected = block
				break RequestLoop
			}
		}
	}

	// At this point, we want to at least log all the candidate blocks
	// that were received with the wrong height. This warning gives the node
	// operator the chance to remove malicious nodes from the configuration.
	for height, blocks := range candidates {
		if height == selected.Height {
			continue
		}
		m.log.Warn().
			Uint64("want", selected.Height).
			Uint64("have", height).
			Int("blocks", len(blocks)).
			Msg("received block(s) with mismatching height")
	}

	// Finally, we should now have the required number of blocks that are the
	// same in every respect, _except_ the number of confirmations. In order to
	// keep the system as secure as possible, we use the lowest number of
	// confirmations returned, as all nodes agree that the block has _at least_
	// that number of confirmations.
	lowest := uint64(math.MaxUint64)
	blocks := candidates[selected.Height]
	for _, block := range blocks {
		if block.Confirmations < lowest {
			lowest = block.Confirmations
		}
	}
	selected.Confirmations = lowest

	return selected, nil
}

// Transaction retrieves the transaction with the given hash and the given
// output index from the underlying Bitcoin API clients. It will only return the
// transaction once the configured number of clients agree on its contents.
func (m *MultiClient) Transaction(hash [32]byte, index uint8) (*bitcoin.Transaction, error) {

	// In the first step, we prepare a channel to handle request results
	// (transactions), a channel to handle request errors, and a channel to
	// signal when all requests have finished.
	results := make(chan *bitcoin.Transaction, len(m.clients))
	errors := make(chan error, len(m.clients))
	done := make(chan struct{})

	// Once the channels are in place, we spin off one goroutine for each
	// underlying client. We use the wait group to keep track of when the
	// requests on all clients are done.
	wg := &sync.WaitGroup{}
	for _, client := range m.clients {
		wg.Add(1)
		go func(client bitcoin.API) {
			transaction, err := client.Transaction(hash, index)
			if err != nil {
				errors <- err
				return
			}
			results <- transaction
		}(client)
	}

	// Here, we spin off a goroutine to handle clean up after all requests are
	// done. It waits on the wait group and then closes and drains all channels.
	// It is also responsible for logging all request errors, just in case the
	// node operator wants to review the nodes he relies on.
	go func() {
		wg.Wait()
		close(results)
		close(errors)
		close(done)
		for range results {
		}
		for err := range errors {
			m.log.Warn().Err(err).Msg("could not get block from client")
		}
	}()

	// The main logic of the multi-client happens in the following section. We
	// wait for results (transactions) to come in on the results channel. Each
	// time we see a result, we check whether it corresponds to the hash and
	// output index we specified and keep track of it in a list of results with
	// the same contents (block hash and recipient). Once we receive the
	// configured amount of transactions with same contents, we consider these
	// contents the correct ones and proceed to the next sections.
	// NOTE: It's possible that we never receive enough results with the same
	// contents. Once all requests are done, the cleanup goroutine sends the
	// done signal and we can return the corresponding error, and we thus avoid
	// getting stuck in an infinite loop.
	var selected *bitcoin.Transaction
	type Key struct {
		block     [32]byte
		recipient string
	}
	candidates := make(map[Key][]*bitcoin.Transaction)
RequestLoop:
	for {
		select {
		case <-done:
			return nil, fmt.Errorf("insufficient number of matching transactions")
		case transaction := <-results:
			if transaction.Hash != hash {
				m.log.Warn().
					Hex("want", transaction.Hash[:]).
					Hex("have", hash[:]).
					Msg("transaction hash does not match requested hash")
				continue RequestLoop
			}
			if transaction.Index != index {
				m.log.Warn().
					Uint8("want", index).
					Uint8("have", transaction.Index).
					Msg("transaction index does not match requested index")
				continue RequestLoop
			}
			key := Key{block: transaction.Block, recipient: transaction.Recipient}
			candidates[key] = append(candidates[key], transaction)
			if uint(len(candidates[key])) >= m.cfg.RequiredMatches {
				selected = transaction
				break RequestLoop
			}
		}
	}

	// At this point, we want to at least log all the candidate transactions
	// that were received with the wrong height. This warning gives the node
	// operator the chance to remove malicious nodes from the configuration.
	for key, transactions := range candidates {
		if key.block == selected.Block && key.recipient == selected.Recipient {
			continue
		}
		m.log.Warn().
			Hex("want_block", selected.Block[:]).
			Hex("have_block", key.block[:]).
			Str("want_recipient", selected.Recipient).
			Str("have_recipient", key.recipient).
			Int("transactions", len(transactions)).Msg("received transaction(s) with mismatching block and/or recipient")
	}

	// Finally, we should now have the required number of transactions that are
	// the same in every respect, _except_ the number of confirmations. In order
	// to keep the system as secure as possible, we use the lowest number of
	// confirmations returned, as all nodes agree that the block has _at least_
	// that number of confirmations.
	lowest := uint64(math.MaxUint64)
	key := Key{block: selected.Block, recipient: selected.Recipient}
	transactions := candidates[key]
	for _, transaction := range transactions {
		if transaction.Confirmations < lowest {
			lowest = transaction.Confirmations
		}
	}
	selected.Confirmations = lowest

	return selected, nil
}
