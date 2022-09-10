package router

import (
	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusOK, in)
	})
}
