package blockusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/output"
)

type VerifyBlockusecase struct {
	blockRepository IBlockRepository
}

func NewVerifyBlock(blockRepo IBlockRepository) *VerifyBlockusecase {
	return &VerifyBlockusecase{
		blockRepository: blockRepo,
	}
}

//===========================================================
//　ブロックが改竄されてないか検証
//===========================================================

func (use *VerifyBlockusecase) Exec(in *input.VerifyBlockInput) (*output.VerifyBlockList, error) {
	bc, err := use.blockRepository.VerifyBlockFindAll(in)
	if err != nil {
		return nil, err
	}

	outVerifyBc := make([]*output.VerifyBlock, len(bc))
	for i, b := range bc {
		outVerifyBc[i] = &output.VerifyBlock{
			ID:           b.Id(),
			Nonce:        b.Nonce(),
			Transactions: b.Transactions(),
			PreviousHash: b.PreviousHash(),
			Timestamp:    b.Timestamp(),
			Hash:         b.Hash(),
			IsCleanData:  b.IsCleanData(),
		}
	}

	return &output.VerifyBlockList{
		VerifyBlockChain: outVerifyBc,
	}, nil
}
