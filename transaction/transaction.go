package transaction

// Transaction represents a transfer of coins.
type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`

	// Digital signature fields
	PublicKey string `json:"public_key,omitempty"`
	Signature string `json:"signature,omitempty"`
}