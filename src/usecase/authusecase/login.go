package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
)

type Loginusecase struct {
	authRepository IAuthRepository
}

func NewLogin(authRepo IAuthRepository) *Loginusecase {
	return &Loginusecase{
		authRepository: authRepo,
	}
}

func (use *Loginusecase) Exec(in *input.LoginInput) (*output.LoginUser, error) {
	user, err := use.authRepository.Login(in)
	if err != nil {
		return nil, err
	}

	key := vo.NewCryptKey()

	return &output.LoginUser{
		User: &output.LoginU{
			ID:        user.Id().Value(),
			Name:      user.Name(),
			CreatedAt: user.CreatedAt().Value(),
			UpdatedAt: user.UpdatedAt().Value(),
		},

		Wallet: &output.LoginUserWallet{
			ID:        user.Wallet().Id().Value(),
			UserID:    user.Wallet().UserID().Value(),
			Address:   user.Wallet().BlockchainAddress(),
			CreatedAt: user.Wallet().CreatedAt().Value(),
			UpdatedAt: user.UpdatedAt().Value(),
		},

		CryptKey: &output.LoginUserCryptKey{
			PrivateKey: key.PrivateKeyValue(),
			PublicKey:  key.PublicKeyValue(),
		},
	}, nil

}
