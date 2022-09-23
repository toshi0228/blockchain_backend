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
	tx, err := use.transactionRepository.Create(in)
	if err != nil {
		return nil, err
	}

	return &output.CreateTransaction{
		ID:               tx.Id().Value(),
		SenderAddress:    tx.SenderAddress(),
		RecipientAddress: tx.RecipientAddress(),
		Value:            tx.Value(),
		CreatedAt:        tx.CreatedAt().Value(),
		UpdatedAt:        tx.UpdatedAt().Value(),
	}, nil
}
