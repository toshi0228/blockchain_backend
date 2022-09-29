package transactionpoolusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/output"
)

type GetTransactionPoolusecase struct {
	transactionPoolRepository ITransactionPoolRepository
}

func NewGetTransactionPool(transactionPoolRepo ITransactionPoolRepository) *GetTransactionPoolusecase {
	return &GetTransactionPoolusecase{
		transactionPoolRepository: transactionPoolRepo,
	}
}

//===========================================================
//　トランザクションプールのデータを取得する
//===========================================================

func (use *GetTransactionPoolusecase) Exec(in *input.GetTransactionPoolInput) (*output.TransactionPool, error) {
	txPool, err := use.transactionPoolRepository.FindAll(in)
	if err != nil {
		return nil, err
	}

	outTxPool := make([]*output.Transaction, len(txPool))
	for i, tx := range txPool {

		outTxPool[i] = &output.Transaction{
			Id:               tx.Id().Value(),
			SenderAddress:    tx.SenderAddress(),
			RecipientAddress: tx.RecipientAddress(),
			Amount:           tx.Amount(),
			CreatedAt:        tx.CreatedAt().Value(),
			UpdatedAt:        tx.UpdatedAt().Value(),
		}
	}

	return &output.TransactionPool{TxPool: outTxPool}, nil
}
