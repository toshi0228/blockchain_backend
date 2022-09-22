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

	//e.POST("/user/login", func(c echo.Context) error {
	//
	//	in := &input.LoginUserInput{}
	//	err := c.Bind(in)
	//	if err != nil {
	//		return fmt.Errorf("エラー")
	//	}
	//
	//	userRepoImpl := database.NewAuthRepositoryImpl()
	//	err = controller.NewAuthController(presenter.NewUserPresenter(c), userRepoImpl).LoginUser(in)
	//	if err != nil {
	//		return c.JSON(http.StatusInternalServerError, err.Error())
	//	}
	//
	//	return nil
	//})

	e.POST("/auth/login", func(c echo.Context) error {

		in := &input.LoginInput{}
		err := c.Bind(in)
		if err != nil {
			return fmt.Errorf("エラー")
		}

		authRepoImpl := database.NewAuthRepositoryImpl()
		err = controller.NewAuthController(presenter.NewAuthPresenter(c), authRepoImpl).Login(in)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	})

}
