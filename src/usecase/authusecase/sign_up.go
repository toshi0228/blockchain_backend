package authusecase

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
	"github.com/toshi0228/blockchain/src/usecase/walletusecase"
)

type SignUpusecase struct {
	AuthRepository   IAuthRepository
	walletRepository walletusecase.IWalletRepository
}

func NewSignUp(AuthRepo IAuthRepository, walletRepo walletusecase.IWalletRepository) *SignUpusecase {
	return &SignUpusecase{
		AuthRepository:   AuthRepo,
		walletRepository: walletRepo,
	}
}

func (use *SignUpusecase) Exec(in *input.SignUpInput) (*output.SignUpUser, error) {

	u, err := use.AuthRepository.SignUp(in)
	if err != nil {
		return nil, err
	}

	w, err := use.walletRepository.Create(u.Id().Value())
	if err != nil {
		return nil, err
	}

	k := vo.NewCryptKey()

	return &output.SignUpUser{

		User: &output.SignUpU{
			ID:        u.Id().Value(),
			Name:      u.Name(),
			CreatedAt: u.CreatedAt().Value(),
			UpdatedAt: u.UpdatedAt().Value(),
		},

		Wallet: &output.SignUpUserWallet{
			Address: w.BlockchainAddress(),
		},

		CryptKey: &output.SignUpUserCryptKey{
			PrivateKey: k.PrivateKeyValue(),
			PublicKey:  k.PublicKeyValue(),
		},
	}, nil
}
