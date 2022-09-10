package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewWalletRouter(e *echo.Echo) {

	e.GET("/wallet", func(c echo.Context) error {
		//u := &User{}
		//u.Name = "a"
		//u.Email = "b"
		//
		//return c.JSON(http.StatusOK, u)
		return nil
	})

	e.POST("/wallet", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
}

//func NewUserRouter(e *echo.Echo) {
//	e.GET("/user", func(c echo.Context) error {
//		return c.String(http.StatusOK, "Hello, World!!!!")
//	})
//}

//e.POST("/user", func(c echo.Context) error {
//
//	in := new(input.CreateUserInput)
//	err := c.Bind(in)
//	if err != nil {
//		fmt.Printf("err %v", err.Error())
//		return c.String(http.StatusBadRequest, "リクエストエラー")
//	}
//
//	userRepoImpl := database.NewUserRepositoryImpl()
//	profileRepoImpl := database.NewProfileRepositoryImpl()
//	skillRepoImpl := database.NewSkillRepositoryImpl()
//	careerRepoImpl := database.NewCareerRepositoryImpl()
//
//	return controller.NewUserController(
//		presenter.NewUserPresenter(c),
//		presenter.NewErrorPresent(c),
//		userRepoImpl, profileRepoImpl, skillRepoImpl, careerRepoImpl).CreateUser(in)
//
//})
