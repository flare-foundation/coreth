package api

type TransactionResponse struct {
	Hash        string `json:"hash"`
	LedgerIndex uint32 `json:"ledger_index"`
	Validated   bool   `json:"validated"`
	// the following fields are from the payment transaction object
	TransactionType string `json:"TransactionType"`
	Destination     string `json:"Destination"`
	DestinationTag  uint32 `json:"DestinationTag"`
	Amount          uint64 `json:"Amount"`
}
