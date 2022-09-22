package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
)

type Loginusecase struct {
	AuthRepository IAuthRepository
}

func NewLogin(AuthRepo IAuthRepository) *Loginusecase {
	return &Loginusecase{
		AuthRepository: AuthRepo,
	}
}

func (use *Loginusecase) Exec(in *input.LoginInput) (*output.LoginUser, error) {
	user, err := use.AuthRepository.Login(in)
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

		Wallet: &output.Wallet{
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
