package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
)

type authController struct {
    delivery presenter.IAuthPresenter
	authRepo authusecase.IAuthRepository
}

func NewAuthController(p presenter.IAuthPresenter, authRepo authusecase.IAuthRepository) *authController {
	return &authController{
	    delivery: p,
		authRepo: authRepo,
	}
}

func (c *authController) Login(in *input.LoginInput) error {
	usecase := authusecase.NewLogin(c.authRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.LoginUser(out)
}
