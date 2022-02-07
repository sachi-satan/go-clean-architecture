package controllers

import (
	"backend/app/inputs"
	"backend/app/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	signUpUseCase usecases.IAuthSignUpUseCase
	signInUseCase usecases.IAuthSignInUseCase
}

func NewAuthController(
	signUpUseCase usecases.IAuthSignUpUseCase,
	signInUseCase usecases.IAuthSignInUseCase,
) *AuthController {
	return &AuthController{
		signUpUseCase: signUpUseCase,
		signInUseCase: signInUseCase,
	}
}

func (r *AuthController) SignUp(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	displayName := c.FormValue("displayName")
	bio := c.FormValue("bio")

	in := inputs.NewAuthSignUpInputData(username, email, password, displayName, bio)
	err := in.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	code, out, err := r.signUpUseCase.Handle(in)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, out)
}

func (r *AuthController) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	in := inputs.NewAuthSignInInputData(email, password)
	err := in.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	code, out, err := r.signInUseCase.Handle(in)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, out)
}
