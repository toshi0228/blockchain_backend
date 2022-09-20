package datamodel

import "time"

type User struct {
	Id        uint32    `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserPager struct {
	Id        uint32    `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	// relationされたwallet情報
	WalletId          uint32    `db:"wallet_id"`
	UserID            uint32    `db:"user_id"`
	BlockchainAddress string    `db:"blockchain_address"`
	WalletCreatedAt   time.Time `db:"wallet_created_at"`
	WalletUpdatedAt   time.Time `db:"wallet_updated_at"`
}
