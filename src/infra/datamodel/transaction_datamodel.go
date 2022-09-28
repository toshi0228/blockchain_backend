package datamodel

import (
	"time"
)

type Transaction struct {
	Id               uint32    `db:"id"`
	SenderAddress    string    `db:"sender_address"`
	RecipientAddress string    `db:"recipient_address"`
	Amount           uint64    `db:"amount"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
