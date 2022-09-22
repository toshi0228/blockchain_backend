package output

import "time"

type LoginUser struct {
	User     *LoginU            `json:"user"`
	Wallet   *LoginUserWallet   `json:"wallet"`
	CryptKey *LoginUserCryptKey `json:"cryptKey"`
}

type LoginU struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginUserWallet struct {
	ID        uint32    `json:"iD"`
	UserID    uint32    `json:"userID"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginUserCryptKey struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}
