package router

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/controller"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/input"
	"net/http"
)

func NewTransactionPoolRouter(e *echo.Echo) {

	//===========================================================
	//　トランザクションプールのデータを取得する
	//===========================================================
	e.GET("/transactionPool", func(c echo.Context) error {

		in := &input.GetTransactionPoolInput{}
		err := c.Bind(in)
		if err != nil {
			return err
		}

		transactionPoolRepoImpl := database.NewTransactionPoolRepositoryImpl()
		err = controller.NewTransactionPoolController(presenter.NewTransactionPoolPresenter(c), transactionPoolRepoImpl).GetTransactionPool(in)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

}

