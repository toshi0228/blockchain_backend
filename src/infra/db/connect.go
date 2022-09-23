package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"sync"
)

var sqlDriver *sqlx.DB
var connErr error

var once sync.Once

// Conn DBに接続　※外部から呼ばれる関数
func Conn() *sqlx.DB {
	once.Do(connImpl)
	return sqlDriver
}

//実際にDBに接続　※実際には一度しか呼ばれない
func connImpl() {

	//環境ごとにDBの接続先が切り替る
	dbConf := toggleDBConf()
	sqlDriver, connErr = sqlx.Open("mysql", dbConf)

	if connErr != nil {
		fmt.Println(connErr)
		log.Fatalln("データベース接続失敗")
	}
}

// 環境の確認
func toggleDBConf() string {
	env := os.Getenv("APP_ENV")
	fmt.Printf("現在接続されている環境: %s \n", env)

	switch env {

	case "dev": //コンテナ => DBコンテナ
		return "docker:docker@tcp(db:3306)/bc_db?parseTime=true&loc=Asia%2FTokyo"

	case "prod": // コンテナ　=> RDB
		return "docker:docker@tcp(db:3306)/bc_db?parseTime=true&loc=Asia%2FTokyo"

	default: //ローカル　=> DBコンテナ (テスト)
		return "docker:docker@tcp(127.0.0.1:3310)/bc_db?parseTime=true&loc=Asia%2FTokyo"
	}
}

func GetConnErr() error {
	return connErr
}
