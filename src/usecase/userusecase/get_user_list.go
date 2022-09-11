package userusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/userusecase/output"
)

type GetUserListusecase struct {
	UserRepository IUserRepository
}

func NewGetUserList(UserRepo IUserRepository) *GetUserListusecase {
	return &GetUserListusecase{
		UserRepository: UserRepo,
	}
}

func (use *GetUserListusecase) Exec() (*output.UserList, error) {
	_, err := use.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &output.UserList{}, nil
}
