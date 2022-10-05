package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
)

//go:generate mockgen -source=./I_auth_repository.go -destination ../../mock/mock_authusecase/auth_repository_mock.go

type IAuthRepository interface {
	SignUp(in *input.SignUpInput) (*entity.User, error)
	Login(in *input.LoginInput) (*entity.UserPager, error)
}
