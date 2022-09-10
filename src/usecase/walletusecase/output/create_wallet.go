package output

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
)

type CreateWallet struct {
	id                vo.ID        `json:"id"`
	userID            vo.ID        `json:"userId"`
	blockchainAddress string       `json:"blockchainAddress"`
	createdAt         vo.CreatedAt `json:"createdAt"`
	updatedAt         vo.UpdatedAt `json:"updatedAt"`
}
