package userusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
)

type IUserRepository interface {
	Login(in *input.LoginUserInput) (*entity.UserPager, error)
	FindAll() ([]*entity.UserPager, error)
	CreateUser(in *input.CreateUserInput) (*entity.User, error)
}
