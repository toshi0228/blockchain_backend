package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/usecase/userusecase"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
	"net/http"
)

func NewUserRouter(e *echo.Echo) {
	e.GET("/user", func(c echo.Context) error {

		in := &input.CreateUserInput{}
		return c.JSON(http.StatusOK, in)
	})

	e.POST("/user", func(c echo.Context) error {

		in := &input.CreateUserInput{}
		err := c.Bind(in)
		if err != nil {
			return fmt.Errorf("エラー")
		}

		userRepoImpl := database.NewUserRepositoryImpl()
		createUser, err := userusecase.NewCreateUser(userRepoImpl).Exec(in)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, createUser)
	})
}
