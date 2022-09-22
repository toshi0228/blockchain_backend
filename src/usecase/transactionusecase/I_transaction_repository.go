package transactionusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
)

type ITransactionRepository interface {
	Create(in *input.CreateTransactionInput) ([]*entity.Transaction, error)
}

