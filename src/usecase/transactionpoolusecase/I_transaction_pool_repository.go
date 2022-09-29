package transactionpoolusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/input"
)

type ITransactionPoolRepository interface {
	FindAll(in *input.GetTransactionPoolInput) ([]*entity.TransactionPool, error)
}

