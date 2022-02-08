package controllers

import (
	"backend/app/inputs"
	"backend/app/models"
	"backend/app/policies"
	"backend/app/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userPolicy    *policies.UserPolicy
	listUseCase   usecases.IUserListUseCase
	getUseCase    usecases.IUserGetUseCase
	updateUseCase usecases.IUserUpdateUseCase
	deleteUseCase usecases.IUserDeleteUseCase
}

func NewUserController(
	userPolicy *policies.UserPolicy,
	listUseCase usecases.IUserListUseCase,
	getUseCase usecases.IUserGetUseCase,
	updateUseCase usecases.IUserUpdateUseCase,
	deleteUseCase usecases.IUserDeleteUseCase,
) *UserController {
	return &UserController{
		userPolicy:    userPolicy,
		listUseCase:   listUseCase,
		getUseCase:    getUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (r *UserController) List(c echo.Context) error {
	id := c.QueryParam("id")
	username := c.QueryParam("username")

	in := inputs.NewUserListInputData(id, username)
	err := in.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	code, out, err := r.listUseCase.Handle(in)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, out)
}

func (r *UserController) Get(c echo.Context) error {
	ID := c.Param("ID")

	in := inputs.NewUserGetInputData(ID)
	err := in.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	code, out, err := r.getUseCase.Handle(in)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, out)
}

func (r *UserController) Update(c echo.Context) error {
	ID := c.Param("ID")
	userName := c.FormValue("username")
	email := c.FormValue("email")
	displayName := c.FormValue("displayName")
	bio := c.FormValue("bio")

	in := inputs.NewUserUpdateInputData(ID, userName, email, displayName, bio)
	err := in.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = r.userPolicy.Update(c.Get("auth").(*models.User), ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	code, out, err := r.updateUseCase.Handle(in)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, out)
}

func (r *UserController) Delete(c echo.Context) error {
	ID := c.Param("ID")

	in := inputs.NewUserDeleteInputData(ID)
	err := in.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = r.userPolicy.Update(c.Get("auth").(*models.User), ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	code, err := r.deleteUseCase.Handle(in)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, nil)
}
