package presenter

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
	"net/http"
)

type authPresent struct {
	ctx echo.Context
}

func NewAuthPresenter(p echo.Context) IAuthPresenter {
	return &authPresent{ctx: p}
}

type IAuthPresenter interface {
	LoginUser(out *output.LoginUser) error
	SignUpUser(out *output.SignUpUser) error
}

func (p *authPresent) LoginUser(out *output.LoginUser) error {
	return p.ctx.JSON(http.StatusOK, out)
}
func (p *authPresent) SignUpUser(out *output.SignUpUser) error {
	return p.ctx.JSON(http.StatusOK, out)
}
