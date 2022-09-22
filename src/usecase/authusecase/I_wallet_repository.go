package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
)

type IWalletRepository interface {
	Create(userId uint32) (*entity.Wallet, error)
}
