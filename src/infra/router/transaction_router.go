package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/controller"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
	"net/http"
)

func NewTransactionRouter(e *echo.Echo) {

	e.POST("/transaction/CreateTransaction", func(c echo.Context) error {

		in := &input.CreateTransactionInput{}
		err := c.Bind(in)
		if err != nil {
			return fmt.Errorf("エラー")
		}

		transactionRepoImpl := database.NewTransactionRepositoryImpl()
		err = controller.NewTransactionController(presenter.NewTransactionPresenter(c), transactionRepoImpl).CreateTransaction(in)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

}

