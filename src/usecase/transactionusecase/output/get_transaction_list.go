package output

import "time"

type Transactions struct {
	TransactionList []*Transaction `json:"transaction"`
}

type Transaction struct {
	Id               uint32    `json:"id"`
	SenderAddress    string    `json:"sender_address"`
	RecipientAddress string    `json:"recipient_address"`
	Amount           uint64    `json:"amount"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
