package output

import "github.com/toshi0228/blockchain/src/entity"

type CreateUser struct {
	Name   string         `json:"name"`
	Wallet *entity.Wallet `json:"wallet"`
}
