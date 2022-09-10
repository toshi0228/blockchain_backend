package userusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
)

type IUserRepository interface {
	CreateUser(in *input.CreateUserInput) (*entity.User, error)
}
