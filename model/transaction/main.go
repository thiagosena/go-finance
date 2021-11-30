package transaction

import "time"

// Transaction type to describe the transaction variables of the application
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions type describe a list of transaction
type Transactions []Transaction
