package main

import (
	"backend/app/controllers"
	"backend/app/interactors"
	"backend/app/middlewares"
	"backend/app/policies"
	"backend/app/repositories"
	"backend/app/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		os.Exit(1)
	}

	mysqlService, err := services.NewMySqlService(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	if err != nil {
		os.Exit(1)
	}

	jwtService, err := services.NewJwtService(
		os.Getenv("PRIVATE_KEY_FILE_PATH"),
		os.Getenv("PUBLIC_KEY_FILE_PATH"),
	)
	if err != nil {
		os.Exit(1)
	}

	userRepository := repositories.NewUserRepository(mysqlService)

	authSignUpInteractor := interactors.NewAuthSignUpInteractor(userRepository, jwtService, mysqlService)
	authSignInInteractor := interactors.NewAuthSingInInteractor(userRepository, jwtService)
	authController := controllers.NewAuthController(authSignUpInteractor, authSignInInteractor)

	userPolicy := policies.NewUserPolicy(userRepository)
	userListInteractor := interactors.NewUserListInteractor(userRepository, mysqlService)
	userGetInteractor := interactors.NewUserGetInteractor(userRepository)
	userUpdateInteractor := interactors.NewUserUpdateInteractor(userRepository, jwtService, mysqlService)
	userDeleteInteractor := interactors.NewUserDeleteInteractor(userRepository, jwtService, mysqlService)
	userController := controllers.NewUserController(userPolicy, userListInteractor, userGetInteractor, userUpdateInteractor, userDeleteInteractor)

	authMiddleWare := middlewares.AuthMiddleWare(jwtService, userRepository)

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
