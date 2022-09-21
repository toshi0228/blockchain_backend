package output

import (
	"time"
)

type LoginUser struct {
	ID        uint32           `json:"id"`
	Name      string           `json:"name"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Wallet    *LoginUserWallet `json:"wallet"`
	CryptKey  *CryptKey        `json:"cryptKey"`
}

type LoginUserWallet struct {
	ID                uint32    `json:"iD"`
	UserID            uint32    `json:"userID"`
	BlockchainAddress string    `json:"blockchainAddress"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type CryptKey struct {
	PrivateKey string
	PublicKey  string
}
