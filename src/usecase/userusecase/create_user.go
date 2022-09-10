package userusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/output"
	"github.com/toshi0228/blockchain/src/usecase/walletusecase"
)

type CreateUserusecase struct {
	userRepository   IUserRepository
	walletRepository walletusecase.IWalletRepository
}

func NewCreateUser(userRepo IUserRepository, walletRepo walletusecase.IWalletRepository) *CreateUserusecase {
	return &CreateUserusecase{
		userRepository:   userRepo,
		walletRepository: walletRepo,
	}
}

func (use *CreateUserusecase) Exec(in *input.CreateUserInput) (*output.CreateUser, error) {

	u, err := use.userRepository.CreateUser(in)
	if err != nil {
		return nil, err
	}

	w, err := use.walletRepository.Create(u.Id().Value())
	if err != nil {
		return nil, err
	}

	return &output.CreateUser{
		Name:          u.Name(),
		WalletAddress: w.BlockchainAddress(),
	}, nil
}
