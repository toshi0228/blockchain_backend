package presenter

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/output"
	"net/http"
)

type userPresent struct {
	ctx echo.Context
}

func NewUserPresenter(p echo.Context) IUserPresenter {
	return &userPresent{ctx: p}
}

type IUserPresenter interface {
	LoginUser(out *output.LoginUser) error
}

func (p *userPresent) LoginUser(out *output.LoginUser) error {
	return p.ctx.JSON(http.StatusOK, out)
}
