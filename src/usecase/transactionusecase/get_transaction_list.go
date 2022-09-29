package transactionusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/output"
)

type GetTransactionListusecase struct {
	transactionRepository ITransactionRepository
}

func NewGetTransactionList(transactionRepo ITransactionRepository) *GetTransactionListusecase {
	return &GetTransactionListusecase{
		transactionRepository: transactionRepo,
	}
}

// ===========================================================
//　トランザクションを取得する
//===========================================================

func (use *GetTransactionListusecase) Exec(in *input.GetTransactionListInput) (*output.Transactions, error) {
	_, err := use.transactionRepository.FindAll(in)
	if err != nil {
		return nil, err
	}

	return &output.Transactions{}, nil
}
