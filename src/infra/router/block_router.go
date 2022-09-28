package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/controller"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
	"net/http"
)

func NewBlockRouter(e *echo.Echo) {

	e.POST("/block", func(c echo.Context) error {

		in := &input.CreateBlockInput{}
		err := c.Bind(in)
		if err != nil {
			return fmt.Errorf("エラー")
		}

		blockRepoImpl := database.NewBlockRepositoryImpl()
		err = controller.NewBlockController(presenter.NewBlockPresenter(c), blockRepoImpl).CreateBlock(in)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

}

