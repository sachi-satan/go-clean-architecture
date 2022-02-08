package interactors

import (
	"backend/app/inputs"
	"backend/app/models"
	"backend/app/outputs"
	repositories "backend/app/repositories/interface"
	"backend/app/services"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type AuthSignUpInteractor struct {
	userRepo repositories.IUserRepository
	service  *services.Service
}

func NewAuthSignUpInteractor(userRepo repositories.IUserRepository, service *services.Service) *AuthSignUpInteractor {
	return &AuthSignUpInteractor{
		userRepo: userRepo,
		service:  service,
	}
}

func (r *AuthSignUpInteractor) Handle(in *inputs.AuthSignUpInputData) (int, *outputs.AuthSignUpOutputData, error) {
	user := models.NewUser()
	err := user.SetNewUser(in.Username, in.Email, in.Password, in.DisplayName, in.Bio)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	err = r.service.Mysql.DB.Transaction(func(tx *gorm.DB) error {
		err := r.userRepo.Save(tx, user)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	accessToken, err := r.service.Jwt.GenToken(strconv.Itoa(user.ID))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	out := outputs.NewAuthSignUpOutputData(accessToken)

	return http.StatusOK, out, nil
}

type AuthSingInInteractor struct {
	userRepo repositories.IUserRepository
	service  *services.Service
}

func NewAuthSingInInteractor(userRepo repositories.IUserRepository, service *services.Service) *AuthSingInInteractor {
	return &AuthSingInInteractor{
		userRepo: userRepo,
		service:  service,
	}
}

func (r *AuthSingInInteractor) Handle(in *inputs.AuthSignInInputData) (int, *outputs.AuthSignInOutputData, error) {
	user, err := r.userRepo.GetByEmail(in.Email)
	if err != nil || user.ValidatePassword(in.Password) != nil {
		return http.StatusUnauthorized, nil, errors.New("the credentials are incorrect")
	}

	accessToken, err := r.service.Jwt.GenToken(strconv.Itoa(user.ID))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	out := outputs.NewAuthSignInOutputData(accessToken)

	return http.StatusOK, out, nil
}
