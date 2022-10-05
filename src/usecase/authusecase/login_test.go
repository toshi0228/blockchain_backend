package authusecase_test

import (
	"github.com/golang/mock/gomock"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/mock/mock_authusecase"
	"github.com/toshi0228/blockchain/src/test"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"testing"
)

var (
	signUpInput      = input.SignUpInput{Name: "test1", Password: "test1"}
	loginInput       = input.LoginInput{Name: "test1", Password: "test1"}
	wrongLoginInput1 = input.LoginInput{Name: "test11", Password: "test1"}
	wrongLoginInput2 = input.LoginInput{Name: "test1", Password: "test2"}
)

func Test_Login_UseCaseExec(t *testing.T) {

	//テスト用の便利関数
	test := test.NewTestLib(t)

	//DBの初期化
	err := test.DBInit()
	if err != nil {
		test.NG(err)
	}

	// ログインするために新規登録を行う
	authRepo := database.NewAuthRepositoryImpl()
	walletRepo := database.NewWalletRepositoryImpl()

	////case : ユーザーの新規登録を行う
	signUpusecase := authusecase.NewSignUp(authRepo, walletRepo)
	signUpOut, err := signUpusecase.Exec(&in)
	if err != nil {
		test.NG(err)
	}

	//　モックの呼び出しを管理するControllerを生成する
	ctrl := gomock.NewController(t)

	// 必要なrepository
	authRepoMock := mock_authusecase.NewMockIAuthRepository(ctrl)

	//ログインして返ってくる値
	user1, _ := entity.NewUserPager(
		signUpOut.User.ID,
		loginInput.Name,
		loginInput.Password,
		signUpOut.User.CreatedAt,
		signUpOut.User.UpdatedAt,
		1,
		signUpOut.User.ID,
		signUpOut.Wallet.Address,
		signUpOut.User.CreatedAt, signUpOut.User.UpdatedAt)

	// テスト中に呼ばれるべき関数と返り値を指定
	authRepoMock.EXPECT().Login(&loginInput).Return(user1, nil)

	//テストケース : ユーザーのログイン
	usecase := authusecase.NewLogin(authRepoMock)
	loginOut, err := usecase.Exec(&loginInput)

	t.Run("新規登録したユーザーでログインできるべき", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}

		if loginOut.User.ID != signUpOut.User.ID {
			t.Errorf("id is not match. result is %v, expected is %v", loginOut.User.ID, signUpOut.User.ID)
		}

		if loginOut.User.Name != signUpOut.User.Name {
			t.Errorf("id is not match. result is %v, expected is %v", loginOut.User.Name, signUpOut.User.Name)
		}

		if loginOut.Wallet.Address != signUpOut.Wallet.Address {
			t.Errorf("id is not match. result is %v, expected is %v", loginOut.Wallet.Address, signUpOut.Wallet.Address)
		}
	})

	t.Run("ログインの際に間違った名前の場合にエラーになるべき", func(t *testing.T) {
		usecase := authusecase.NewLogin(authRepo)
		_, err := usecase.Exec(&wrongLoginInput1)
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("ログインの際に間違ったパスワードの場合にエラーになるべき", func(t *testing.T) {
		usecase := authusecase.NewLogin(authRepo)
		_, err := usecase.Exec(&wrongLoginInput2)
		if err == nil {
			t.Error(err)
		}
	})
}
