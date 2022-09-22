package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
)

type transactionController struct {
    delivery presenter.ITransactionPresenter
	transactionRepo transactionusecase.ITransactionRepository
}

func NewTransactionController(p presenter.ITransactionPresenter, transactionRepo transactionusecase.ITransactionRepository) *transactionController {
	return &transactionController{
	    delivery: p,
		transactionRepo: transactionRepo,
	}
}

func (c *transactionController) CreateTransaction(in *input.CreateTransactionInput) error {
	usecase := transactionusecase.NewCreateTransaction(c.transactionRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.CreateTransaction(out)
}
