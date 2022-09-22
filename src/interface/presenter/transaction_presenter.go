package presenter

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/output"
	"net/http"
)

type transactionPresent struct {
	ctx echo.Context
}

func NewTransactionPresenter(p echo.Context) ITransactionPresenter {
	return &transactionPresent{ctx: p}
}

type ITransactionPresenter interface {
	CreateTransaction(out *output.CreateTransaction) error
}

func (p *transactionPresent) CreateTransaction(out *output.CreateTransaction) error {
	return p.ctx.JSON(http.StatusOK, out)
}
