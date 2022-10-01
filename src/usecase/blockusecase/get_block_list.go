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
	bc, err := use.blockRepository.FindAll(in)
	if err != nil {
		return nil, err
	}

	outBc := make([]*output.Block, len(bc))
	for i, b := range bc {
		outBc[i] = &output.Block{
			ID:           b.Id(),
			Nonce:        b.Nonce(),
			Transactions: b.Transactions(),
			PreviousHash: b.PreviousHash(),
			Timestamp:    b.Timestamp(),
			Hash:         b.Hash(),
		}
	}

	return &output.BlockList{
		BlockChain: outBc,
	}, nil
}
