package presenter

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/output"
	"net/http"
)

type transactionPoolPresent struct {
	ctx echo.Context
}

func NewTransactionPoolPresenter(p echo.Context) ITransactionPoolPresenter {
	return &transactionPoolPresent{ctx: p}
}

type ITransactionPoolPresenter interface {
	TransactionPool(out *output.TransactionPool) error
}

//===========================================================
//　トランザクションプールのデータを取得する
//===========================================================

func (p *transactionPoolPresent) TransactionPool(out *output.TransactionPool) error {
	return p.ctx.JSON(http.StatusOK, out)
}
