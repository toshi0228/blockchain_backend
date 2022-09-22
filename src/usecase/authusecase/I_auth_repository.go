package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
)

type IAuthRepository interface {
	Login(in *input.LoginInput) (*entity.UserPager, error)
}
