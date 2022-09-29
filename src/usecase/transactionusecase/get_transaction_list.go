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
	txs, err := use.transactionRepository.FindAll(in)
	if err != nil {
		return nil, err
	}

	outTxs := make([]*output.Transaction, len(txs))
	for i, tx := range txs {

		outTxs[i] = &output.Transaction{
			Id:               tx.Id().Value(),
			SenderAddress:    tx.SenderAddress(),
			RecipientAddress: tx.RecipientAddress(),
			Amount:           tx.Value(),
			CreatedAt:        tx.CreatedAt().Value(),
			UpdatedAt:        tx.UpdatedAt().Value(),
		}
	}

	return &output.Transactions{TransactionList: outTxs}, nil
}
