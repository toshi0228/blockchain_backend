package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/datamodel"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
	"log"
)

type BlockRepositoryImpl struct{}

func NewBlockRepositoryImpl() *BlockRepositoryImpl {
	return &BlockRepositoryImpl{}
}

//===========================================================
// ブロックの作成
//===========================================================

//ブロック作成のSQL
//go:embed block_repository_create.sql
var createBlockSql string

//一個前のハッシュを取得するSQL
//go:embed block_repository_find_prev_block.sql
var findPrevBlockSql string

//トランザクションプールにある全件のトランザクションデータを取得する
//go:embed block_repository_find_all_transaction_pool.sql
var findAllTransactionPoolSql string

//トランザクションプールのデータを削除
//go:embed block_repository_delete_transaction_pool.sql
var deleteTransactionPoolSql string

func (repo *BlockRepositoryImpl) Create(in *input.CreateBlockInput) ([]*entity.Block, error) {

	//1つ前のブロックのハッシュを取得
	row := db.Conn().QueryRow(findPrevBlockSql)
	var prevHash string
	err := row.Scan(&prevHash)

	var transactionPool []string

	// トランザクションプールの値を取得
	txRows, err := db.Conn().Queryx(findAllTransactionPoolSql)
	if err != nil {
		return nil, err
	}

	//dbから取得したtxPoolデータを処理する
	for txRows.Next() {
		//dbのカラムの型を定義
		tx := &datamodel.Transaction{}
		err := txRows.StructScan(tx)

		if err != nil {
			fmt.Println(err)
		}

		txInPool, err := entity.NewTransactionPool(
			tx.Id,
			tx.SenderAddress,
			tx.RecipientAddress,
			tx.Amount,
			tx.CreatedAt,
			tx.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		//トランザクションプールに入っているtxデータをtransactionPoolに入れる　また入れる際はjsonにして渡す
		transactionPool = append(transactionPool, txInPool.ToJSON())
	}

	//またblockがない場合は最初のブロックを作成する
	if err != nil {
		in.Nonce = 1
		b, err := entity.GenWhenCreateBlock(prevHash, []string{""})
		_, err = db.Conn().Exec(
			createBlockSql,
			b.Nonce(),
			b.PreviousHash(),
			b.Transactions(),
			b.Timestamp(),
			b.CalcHash(),
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		return nil, nil
	}

	//通常のブロック作成処理
	b, err := entity.GenWhenCreateBlock(prevHash, transactionPool)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	_, err = db.Conn().Exec(
		createBlockSql,
		b.Nonce(),
		b.PreviousHash(),
		b.Transactions(),
		b.Timestamp(),
		b.CalcHash(),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	//トランザクションをblockに詰め込んだらpoolを削除する
	_, err = db.Conn().Exec(
		deleteTransactionPoolSql,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//===========================================================
//　ブロックリスト(ブロックチェーン)を取得する
//===========================================================

//go:embed block_repository_find_all.sql
var findAllBlockSql string

func (repo *BlockRepositoryImpl) FindAll(in *input.GetBlockListInput) ([]*entity.Block, error) {

	//DBからブロックのリストを取得する
	var blockChain []*datamodel.BlocK
	err := db.Conn().Select(&blockChain, findAllBlockSql)
	if err != nil {
		return nil, err
	}

	// DBから取得したentity
	var entityBC []*entity.Block
	for _, b := range blockChain {
		blockEntity, err := entity.NewBlock(b.Id, b.Nonce, b.Hash, b.PreviousHash, b.Transactions, b.Timestamp)
		if err != nil {
			return nil, err
		}
		entityBC = append(entityBC, blockEntity)
	}

	return entityBC, nil
}
