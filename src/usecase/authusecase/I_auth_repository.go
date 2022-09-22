package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
)

type IAuthRepository interface {
	SignUp(in *input.SignUpInput) (*entity.User, error)
	Login(in *input.LoginInput) (*entity.UserPager, error)
}
