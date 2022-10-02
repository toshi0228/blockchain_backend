package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
)

type SignUpusecase struct {
	authRepository   IAuthRepository
	walletRepository IWalletRepository
}

func NewSignUp(authRepo IAuthRepository, walletRepo IWalletRepository) *SignUpusecase {
	return &SignUpusecase{
		authRepository:   authRepo,
		walletRepository: walletRepo,
	}
}

func (use *SignUpusecase) Exec(in *input.SignUpInput) (*output.SignUpUser, error) {

	user, err := use.authRepository.SignUp(in)
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
