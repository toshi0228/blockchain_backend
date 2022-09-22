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
	users, err := use.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	userItems := make([]*output.UserItem, len(users))
	for i, user := range users {

		userItems[i] = &output.UserItem{
			ID:        user.Id().Value(),
			Name:      user.Name(),
			CreatedAt: user.CreatedAt().Value(),
			UpdatedAt: user.UpdatedAt().Value(),

			Wallet: &output.WalletItem{
				ID:        user.Wallet().Id().Value(),
				UserID:    user.Wallet().UserID().Value(),
				Address:   user.Wallet().BlockchainAddress(),
				CreatedAt: user.Wallet().CreatedAt().Value(),
				UpdatedAt: user.UpdatedAt().Value(),
			},
		}
	}

	return &output.UserList{Users: userItems}, nil
}
