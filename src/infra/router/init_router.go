package router

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	NewUserRouter(e)
	NewWalletRouter(e)
}
