package ripple

type LedgerResponse struct {
	LedgerHash  string `json:"ledger_hash"`
	LedgerIndex uint32 `json:"ledger_index"`
	Validated   bool   `json:"validated"`
}
