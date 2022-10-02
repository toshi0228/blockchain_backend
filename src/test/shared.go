package test

import (
	"fmt"
	"github.com/toshi0228/blockchain/src/infra/db"
	"strings"
	"testing"
)

type testLiv struct {
	liv *testing.T
}

func NewTestLib(t *testing.T) *testLiv {
	return &testLiv{liv: t}
}

// DBInit テスト用データベースリセット
func (t *testLiv) DBInit() error {
	err := db.Reset()
	err = db.CreateTables()

	if err != nil {
		return err
	}

	return nil
}

// Title テストタイトル
func (t *testLiv) Title(title string) {
	fmt.Println("")
	fmt.Printf("%s \n", strings.Repeat("-", 48))
	fmt.Printf("%v \n", title)
	fmt.Printf("%s \n", strings.Repeat("-", 48))

}

func (t *testLiv) OK() {
	fmt.Println("✅ OK")
}

func (t *testLiv) NG(err error) {
	fmt.Println("🙅‍️ NG")
	t.liv.Errorf("エラー")
}
