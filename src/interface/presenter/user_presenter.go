package presenter

import (
	"github.com/labstack/echo/v4"
)

type userPresent struct {
	ctx echo.Context
}

func NewUserPresenter(p echo.Context) IUserPresenter {
	return &userPresent{ctx: p}
}

type IUserPresenter interface {
}
