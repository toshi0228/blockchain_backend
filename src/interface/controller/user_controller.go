package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/userusecase"
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
