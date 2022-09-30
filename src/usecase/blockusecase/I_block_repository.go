package blockusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
)

type IBlockRepository interface {
    FindAll(in *input.GetBlockListInput) ([]*entity.Block, error)

	Create(in *input.CreateBlockInput) ([]*entity.Block, error)
}

