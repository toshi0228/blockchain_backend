package presenter

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/output"
	"net/http"
)

type blockPresent struct {
	ctx echo.Context
}

func NewBlockPresenter(p echo.Context) IBlockPresenter {
	return &blockPresent{ctx: p}
}

type IBlockPresenter interface {
	CreateBlock(out *output.CreateBlock) error
}

func (p *blockPresent) CreateBlock(out *output.CreateBlock) error {
	return p.ctx.JSON(http.StatusOK, out)
}
