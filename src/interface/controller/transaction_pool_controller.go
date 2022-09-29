package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/input"
)

type transactionPoolController struct {
    delivery presenter.ITransactionPoolPresenter
	transactionPoolRepo transactionpoolusecase.ITransactionPoolRepository
}

func NewTransactionPoolController(p presenter.ITransactionPoolPresenter, transactionPoolRepo transactionpoolusecase.ITransactionPoolRepository) *transactionPoolController {
	return &transactionPoolController{
	    delivery: p,
		transactionPoolRepo: transactionPoolRepo,
	}
}

//===========================================================
//　トランザクションプールのデータを取得する
//===========================================================

func (c *transactionPoolController) GetTransactionPool(in *input.GetTransactionPoolInput) error {
	usecase := transactionpoolusecase.NewGetTransactionPool(c.transactionPoolRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.TransactionPool(out)
}
