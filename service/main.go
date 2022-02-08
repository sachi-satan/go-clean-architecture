package main

import (
	"backend/app/controllers"
	"backend/app/interactors"
	"backend/app/middlewares"
	"backend/app/policies"
	"backend/app/repositories"
	"backend/app/services"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	service, err := services.NewService()
	if err != nil {
		os.Exit(1)
	}

	userRepository := repositories.NewUserRepository(service.Mysql)

	authSignUpInteractor := interactors.NewAuthSignUpInteractor(userRepository, service)
	authSignInInteractor := interactors.NewAuthSingInInteractor(userRepository, service)
	authController := controllers.NewAuthController(authSignUpInteractor, authSignInInteractor)

	userPolicy := policies.NewUserPolicy(userRepository)
	userListInteractor := interactors.NewUserListInteractor(userRepository, service)
	userGetInteractor := interactors.NewUserGetInteractor(userRepository)
	userUpdateInteractor := interactors.NewUserUpdateInteractor(userRepository, service)
	userDeleteInteractor := interactors.NewUserDeleteInteractor(userRepository, service)
	userController := controllers.NewUserController(userPolicy, userListInteractor, userGetInteractor, userUpdateInteractor, userDeleteInteractor)

	authMiddleWare := middlewares.AuthMiddleWare(service.Jwt, userRepository)

	e := echo.New()
	g := e.Group("/v1")

	g.POST("/signUp", authController.SignUp)
	g.POST("/signIn", authController.SignIn)

	g.GET("/users", userController.List)
	g.GET("/users/:ID", userController.Get)
	g.PATCH("/users/:ID", userController.Update, authMiddleWare)
	g.DELETE("/users/:ID", userController.Delete, authMiddleWare)

	e.Logger.Fatal(e.Start(":1323"))
}
