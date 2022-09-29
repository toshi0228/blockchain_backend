package output

import "time"

type Transactions struct {
	TransactionList []*Transaction `json:"transactions"`
}

type Transaction struct {
	Id               uint32    `json:"id"`
	SenderAddress    string    `json:"senderAddress"`
	RecipientAddress string    `json:"recipientAddress"`
	Amount           uint64    `json:"amount"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
