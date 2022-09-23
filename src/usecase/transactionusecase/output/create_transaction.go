package output

import "time"

type CreateTransaction struct {
	ID               uint32    `json:"id"`
	SenderAddress    string    `json:"senderAddress"`
	RecipientAddress string    `json:"recipientAddress"`
	Value            uint64    `json:"value"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
