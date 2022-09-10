package db

import (
	"fmt"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
)

var createTableCMD = map[string]string{
	"users":    dbtable.Users(),
	"profiles": dbtable.Profiles(),
}

// CreateTables テーブルの作成 (テーブルがなければ作成される)
func CreateTables() error {

	// テーブル作成する際に外部制約のチェックの解除
	Conn().Exec("SET FOREIGN_KEY_CHECKS = 0")

	// テーブルの作成
	for k, v := range createTableCMD {
		fmt.Println("テーブル作成:" + k)
		err := createTable(v)

		if err != nil {
			fmt.Printf("テーブル作成: + %s 失敗 \n", k)
		}
	}

	// 外部キーチェックの設定をデフォルトに戻す
	Conn().Exec("SET FOREIGN_KEY_CHECKS = 1")

	return GetConnErr()
}

func createTable(cmd string) error {

	_, err := Conn().Exec(cmd)

	return err
}
