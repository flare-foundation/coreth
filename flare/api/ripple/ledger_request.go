package ripple

type LedgerRequest struct {
	LedgerIndex  uint32 `json:"ledger_index"`
	Accounts     bool   `json:"accounts"`
	Full         bool   `json:"full"`
	Transactions bool   `json:"transactions"`
	Expand       bool   `json:"expand"`
	OwnerFunds   bool   `json:"owner_funds"`
}
