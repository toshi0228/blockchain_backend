package db

import (
	"fmt"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
)

var tableList = map[string]string{
	"users":        dbtable.TableNameUser,
	"profiles":     dbtable.TableNameProfile,
	"wallets":      dbtable.TableNameWallet,
	"transactions": dbtable.TableNameTransaction,
}

func Reset() error {

	// dropする際に外部制約のチェックの解除 します
	Conn().Exec("SET FOREIGN_KEY_CHECKS = 0")

	// テーブルの削除
	for _, v := range tableList {
		fmt.Println("テーブル削除:" + v)
		err := dropTable(v)

		if err != nil {
			fmt.Printf("テーブル削除: + %s 失敗 \n", v)
		}
	}

	// 外部キーチェックの設定をデフォルトに戻す
	Conn().Exec("SET FOREIGN_KEY_CHECKS = 1")

	return GetConnErr()
}

func dropTable(tableName string) error {

	// language=sql
	cmd := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, tableName)
	_, err := Conn().Exec(cmd)

	return err
}
