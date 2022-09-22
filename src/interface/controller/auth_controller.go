package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
)

type authController struct {
	delivery   presenter.IAuthPresenter
	authRepo   authusecase.IAuthRepository
	walletRepo authusecase.IWalletRepository
}

func NewAuthController(p presenter.IAuthPresenter, authRepo authusecase.IAuthRepository, walletRepo authusecase.IWalletRepository) *authController {
	return &authController{
		delivery:   p,
		authRepo:   authRepo,
		walletRepo: walletRepo,
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
func (c *authController) SignUp(in *input.SignUpInput) error {
	usecase := authusecase.NewSignUp(c.authRepo, c.walletRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.SignUpUser(out)
}
