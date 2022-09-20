package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/userusecase"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
)

type userController struct {
	delivery presenter.IUserPresenter
	userRepo userusecase.IUserRepository
}

func NewUserController(p presenter.IUserPresenter, userRepo userusecase.IUserRepository) *userController {
	return &userController{
		delivery: p,
		userRepo: userRepo,
	}
}

func (c *userController) LoginUser(in *input.LoginUserInput) error {
	usecase := userusecase.NewLoginUser(c.userRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.LoginUser(out)
}
