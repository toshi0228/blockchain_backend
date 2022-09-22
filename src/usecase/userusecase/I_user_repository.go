package userusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
)

type IUserRepository interface {
	FindAll() ([]*entity.UserPager, error)
}
