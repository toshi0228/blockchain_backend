package blockusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/output"
)

type GetBlockListusecase struct {
	blockRepository IBlockRepository
}

func NewGetBlockList(blockRepo IBlockRepository) *GetBlockListusecase {
	return &GetBlockListusecase{
		blockRepository: blockRepo,
	}
}

//===========================================================
//　ブロックリスト(ブロックチェーン)を取得する
//===========================================================

func (use *GetBlockListusecase) Exec(in *input.GetBlockListInput) (*output.BlockList, error) {
	_, err := use.blockRepository.FindAll(in)
	if err != nil {
		return nil, err
	}

	return &output.BlockList{}, nil
}

