package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/controller"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"net/http"
)

func NewAuthRouter(e *echo.Echo) {
	e.POST("/auth/signup", func(c echo.Context) error {

		in := &input.SignUpInput{}
		err := c.Bind(in)
		if err != nil {
			return fmt.Errorf("エラー")
		}

		authRepoImpl := database.NewAuthRepositoryImpl()
		walletRepoImpl := database.NewWalletRepositoryImpl()

		err = controller.NewAuthController(presenter.NewAuthPresenter(c), authRepoImpl, walletRepoImpl).SignUp(in)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

	e.POST("/auth/login", func(c echo.Context) error {

		in := &input.LoginInput{}
		err := c.Bind(in)
		if err != nil {
			return fmt.Errorf("エラー")
		}

		authRepoImpl := database.NewAuthRepositoryImpl()
		walletRepoImpl := database.NewWalletRepositoryImpl()
		err = controller.NewAuthController(presenter.NewAuthPresenter(c), authRepoImpl, walletRepoImpl).Login(in)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

}
