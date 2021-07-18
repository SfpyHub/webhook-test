package definitions

type PaymentNotification struct {
	Token          string `json:"token"`
	NotificationID string `json:"notification_id"`
	RequestID      string `json:"request_id"`
	PaymentID      string `json:"payment_id"`
	ChainID        uint   `json:"chain_id"`
	State          string `json:"state"`
	IsETH          bool   `json:"is_eth"`
	TxnHash        string `json:"txn_hash"`
	From           string `json:"from"`
	Amount         string `json:"amount"`
	TokenAddress   string `json:"token_address"`
	Rate           string `json:"rate"`
}

type RefundNotification struct {
	Token          string `json:"token"`
	NotificationID string `json:"notification_id"`
	RequestID      string `json:"request_id"`
	PaymentID      string `json:"payment_id"`
	RefundID       string `json:"refund_id"`
	ChainID        uint   `json:"chain_id"`
	IsETH          bool   `json:"is_eth"`
	TxnHash        string `json:"txn_hash"`
	From           string `json:"from"`
	To             string `json:"to"`
	Amount         string `json:"amount"`
	TokenAddress   string `json:"token_address"`
}
