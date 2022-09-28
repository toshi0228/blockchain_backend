package blockusecase

import (
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
)

type IBlockRepository interface {
	Create(in *input.CreateBlockInput) ([]*entity.Block, error)
}

