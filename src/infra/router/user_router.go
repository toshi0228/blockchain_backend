package router

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/usecase/userusecase"
	"net/http"
)

func NewUserRouter(e *echo.Echo) {
	e.GET("/user", func(c echo.Context) error {
		//in := &input.CreateUserInput{}

		userRepoImpl := database.NewUserRepositoryImpl()
		userList, err := userusecase.NewGetUserList(userRepoImpl).Exec()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, userList)
	})
}
