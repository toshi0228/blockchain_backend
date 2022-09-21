package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/infra/router"
	"strings"
)

type Sample struct {
	test string
}

func main() {
	fmt.Printf("%s main関数 %s \n", strings.Repeat("=", 25), strings.Repeat("=", 25))

	//データベースの接続
	db.Conn()

	// 開発環境の際は変更するたびにDBを削除する
	//db.Reset()

	// テーブルの作成 (テーブルがなければ作成される)
	db.CreateTables()

	// インスタンスを作成
	e := echo.New()

	router.InitRouter(e)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":4000"))

}
