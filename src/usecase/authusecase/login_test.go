package authusecase_test

import (
	"github.com/golang/mock/gomock"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/mock/mock_authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"testing"
	"time"
)

var (
	loginInput = input.LoginInput{Name: "test1", Password: "test1"}
	user1, _   = entity.NewUserPager(424262867, loginInput.Name, loginInput.Password, time.Now(), time.Now(),
		1, 1, "123", time.Now(), time.Now())
)

func TestLoginusecase_Exec(t *testing.T) {

	//テスト用の便利関数
	//ttt := test.NewTestLib(t)

	//DBの初期化
	//err := ttt.DBInit()
	//if err != nil {
	//	ttt.NG(err)
	//}

	//　モックの呼び出しを管理するControllerを生成する
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 必要なrepository
	authRepo := mock_authusecase.NewMockIAuthRepository(ctrl)
	// テスト中に呼ばれるべき関数と返り値を指定
	authRepo.EXPECT().Login(&loginInput).Return(user1, nil)

	//case : ユーザーの新規登録を行う
	usecase := authusecase.NewLogin(authRepo)
	out, err := usecase.Exec(&loginInput)

	t.Run("ウォレットが作成されているべき", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}

		if out.User.ID != user1.Id().Value() {
			t.Errorf("id is not match. result is , expected is ")
		}
	})

}
