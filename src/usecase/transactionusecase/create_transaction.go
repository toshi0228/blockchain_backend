package transactionusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/output"
)

type CreateTransactionusecase struct {
	transactionRepository ITransactionRepository
}

func NewCreateTransaction(transactionRepo ITransactionRepository) *CreateTransactionusecase {
	return &CreateTransactionusecase{
		transactionRepository: transactionRepo,
	}
}

func (use *CreateTransactionusecase) Exec(in *input.CreateTransactionInput) (*output.CreateTransaction, error) {
	_, err := use.transactionRepository.Create(in)
	if err != nil {
		return nil, err
	}

	return &output.CreateTransaction{}, nil
}

