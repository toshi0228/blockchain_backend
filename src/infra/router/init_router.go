package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(e *echo.Echo) {

	// CORSの設定追加
	e.Use(middleware.CORS())

	NewUserRouter(e)
	NewWalletRouter(e)
	NewAuthRouter(e)
	NewTransactionRouter(e)
	NewBlockRouter(e)
}
