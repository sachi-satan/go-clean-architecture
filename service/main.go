package main

import (
	"backend/app/controllers"
	"backend/app/interactors"
	"backend/app/repositories"
	"backend/app/services"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
)

func main() {
	dsn := "docker:docker@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	jwtService, err := services.NewJwtService("app.rsa", "app.rsa.pub")
	if err != nil {
		os.Exit(1)
	}

	userRepository := repositories.NewUserRepository(db)

	authSignUpInteractor := interactors.NewAuthSignUpInteractor(userRepository, jwtService)
	authSignInInteractor := interactors.NewAuthSingInInteractor(userRepository, jwtService)
	authController := controllers.NewAuthController(authSignUpInteractor, authSignInInteractor)

	userListInteractor := interactors.NewUserListInteractor(userRepository)
	userGetInteractor := interactors.NewUserGetInteractor(userRepository)
	userUpdateInteractor := interactors.NewUserUpdateInteractor(userRepository, jwtService)
	userDeleteInteractor := interactors.NewUserDeleteInteractor(userRepository, jwtService)
	userController := controllers.NewUserController(userListInteractor, userGetInteractor, userUpdateInteractor, userDeleteInteractor)

	jwtAuth := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userIdParam := c.Param("ID")
			auth := c.Request().Header.Get("Authorization")

			subs := strings.Split(auth, " ")
			if len(subs) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "your client does not have permission to the requested")
			}

			userId, err := jwtService.ParseToken(subs[1])

			if err != nil || userId != userIdParam {
				return echo.NewHTTPError(http.StatusUnauthorized, "your client does not have permission to the requested")
			}

			return next(c)
		}
	}

	e := echo.New()
	g := e.Group("/v1")

	g.POST("/signUp", authController.SignUp)
	g.POST("/signIn", authController.SignIn)

	g.GET("/users", userController.List)
	g.GET("/users/:ID", userController.Get)
	g.PATCH("/users/:ID", userController.Update, jwtAuth)
	g.DELETE("/users/:ID", userController.Delete, jwtAuth)

	e.Logger.Fatal(e.Start(":1323"))
}
