package output

import "time"

type SignUpUser struct {
	User     *SignUpU            `json:"user"`
	Wallet   *SignUpUserWallet   `json:"wallet"`
	CryptKey *SignUpUserCryptKey `json:"cryptKey"`
}

type SignUpU struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SignUpUserWallet struct {
	Address string `json:"address"`
}

type SignUpUserCryptKey struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}
