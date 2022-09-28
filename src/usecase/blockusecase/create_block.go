package blockusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/output"
)

type CreateBlockusecase struct {
	blockRepository IBlockRepository
}

func NewCreateBlock(blockRepo IBlockRepository) *CreateBlockusecase {
	return &CreateBlockusecase{
		blockRepository: blockRepo,
	}
}

func (use *CreateBlockusecase) Exec(in *input.CreateBlockInput) (*output.CreateBlock, error) {
	_, err := use.blockRepository.Create(in)
	if err != nil {
		return nil, err
	}

	return &output.CreateBlock{}, nil
}

