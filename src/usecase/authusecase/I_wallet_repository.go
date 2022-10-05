package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
)

//go:generate mockgen -source=./I_wallet_repository.go -destination ../../mock/mock_authusecase/wallet_repository_mock.go

type IWalletRepository interface {
	Create(userId uint32) (*entity.Wallet, error)
}
