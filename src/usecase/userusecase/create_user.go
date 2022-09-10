package userusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
)

type CreateUserusecase struct {
	userRepository IUserRepository
}

func NewCreateUser(userRepo IUserRepository) *CreateUserusecase {
	return &CreateUserusecase{
		userRepository: userRepo,
	}
}

func (use *CreateUserusecase) Exec(in *input.CreateUserInput) (*entity.User, error) {

	u, err := use.userRepository.CreateUser(in)
	if err != nil {
		return nil, err
	}

	return u, nil
}
