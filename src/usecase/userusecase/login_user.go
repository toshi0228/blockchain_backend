package userusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/output"
)

type LoginUserusecase struct {
	UserRepository IUserRepository
}

func NewLoginUser(UserRepo IUserRepository) *LoginUserusecase {
	return &LoginUserusecase{
		UserRepository: UserRepo,
	}
}

func (use *LoginUserusecase) Exec(in *input.LoginUserInput) (*output.LoginUser, error) {
	u, err := use.UserRepository.Login(in)
	if err != nil {
		return nil, err
	}

	return &output.LoginUser{
		Name:   u.Name(),
		Status: "ok",
	}, nil
}