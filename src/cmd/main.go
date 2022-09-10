package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/router"
	"strings"
)

type Sample struct {
	test string
}

func main() {
	fmt.Printf("%s main関数 %s \n", strings.Repeat("=", 25), strings.Repeat("=", 25))

	// インスタンスを作成
	e := echo.New()
	router.InitRouter(e)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":4000"))

	walletM := entity.NewWallet()
	walletA := entity.NewWallet()
	walletB := entity.NewWallet()

	//FE
	sentT := entity.NewSentTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockChainAddress, 1.0)

	//BC Node
	blockchain := entity.NewBlockChain(walletM.BlockChainAddress)
	isAdded := blockchain.AddTransaction(walletA.BlockChainAddress, walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), sentT.GenerateSignature())
	fmt.Println("Added?", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("C %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
