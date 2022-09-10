package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
	"log"
)

type WalletRepositoryImpl struct{}

//go:embed wallet_repository_create.sql
var createWalletSql string

func NewWalletRepositoryImpl() *WalletRepositoryImpl {
	return &WalletRepositoryImpl{}
}

func (repo *WalletRepositoryImpl) Create(userID uint32) (*entity.Wallet, error) {

	w, err := entity.NewWallet(userID)
	if err != nil {
		return nil, err
	}

	cmd := fmt.Sprintf(createWalletSql, dbtable.TableNameWallet)

	_, err = db.Conn().Exec(
		cmd,
		w.Id(),
		w.UserID(),
		w.BlockchainAddress(),
		w.CreatedAt().Value(),
		w.UpdatedAt().Value(),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return w, nil
}
