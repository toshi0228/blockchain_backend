package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
)

type SignUpusecase struct {
	AuthRepository   IAuthRepository
	walletRepository IWalletRepository
}

func NewSignUp(AuthRepo IAuthRepository, walletRepo IWalletRepository) *SignUpusecase {
	return &SignUpusecase{
		AuthRepository:   AuthRepo,
		walletRepository: walletRepo,
	}
}

func (use *SignUpusecase) Exec(in *input.SignUpInput) (*output.SignUpUser, error) {

	user, err := use.AuthRepository.SignUp(in)
	if err != nil {
		return nil, err
	}

	wallet, err := use.walletRepository.Create(user.Id().Value())
	if err != nil {
		return nil, err
	}

	key := vo.NewCryptKey()

	return &output.SignUpUser{

		User: &output.SignUpU{
			ID:        user.Id().Value(),
			Name:      user.Name(),
			CreatedAt: user.CreatedAt().Value(),
			UpdatedAt: user.UpdatedAt().Value(),
		},

		Wallet: &output.SignUpUserWallet{
			Address: wallet.BlockchainAddress(),
		},

		CryptKey: &output.SignUpUserCryptKey{
			PrivateKey: key.PrivateKeyValue(),
			PublicKey:  key.PublicKeyValue(),
		},
	}, nil
}
