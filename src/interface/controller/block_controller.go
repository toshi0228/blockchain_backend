package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
)

type blockController struct {
    delivery presenter.IBlockPresenter
	blockRepo blockusecase.IBlockRepository
}

func NewBlockController(p presenter.IBlockPresenter, blockRepo blockusecase.IBlockRepository) *blockController {
	return &blockController{
	    delivery: p,
		blockRepo: blockRepo,
	}
}

func (c *blockController) CreateBlock(in *input.CreateBlockInput) error {
	usecase := blockusecase.NewCreateBlock(c.blockRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.CreateBlock(out)
}
//===========================================================
//　ブロックリスト(ブロックチェーン)を取得する
//===========================================================

func (c *blockController) GetBlockList(in *input.GetBlockListInput) error {
	usecase := blockusecase.NewGetBlockList(c.blockRepo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.BlockList(out)
}

