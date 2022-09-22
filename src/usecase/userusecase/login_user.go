package userusecase

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
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
	user, err := use.UserRepository.Login(in)
	if err != nil {
		return nil, err
	}

	key := vo.NewCryptKey()

	return &output.LoginUser{
		User: &output.User{
			ID:        user.Id().Value(),
			Name:      user.Name(),
			CreatedAt: user.CreatedAt().Value(),
			UpdatedAt: user.UpdatedAt().Value(),
		},

		Wallet: &output.UserWallet{
			ID:                user.Wallet().Id().Value(),
			UserID:            user.Wallet().UserID().Value(),
			BlockchainAddress: user.Wallet().BlockchainAddress(),
			CreatedAt:         user.Wallet().CreatedAt().Value(),
			UpdatedAt:         user.UpdatedAt().Value(),
		},

		CryptKey: &output.CryptKey{
			PrivateKey: key.PrivateKeyValue(),
			PublicKey:  key.PublicKeyValue(),
		},
	}, nil
}
