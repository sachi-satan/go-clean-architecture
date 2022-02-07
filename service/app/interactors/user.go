package interactors

import (
	"backend/app/inputs"
	"backend/app/outputs"
	repositories "backend/app/repositories/interface"
	"backend/app/services"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

// List

type UserListInteractor struct {
	userRepo repositories.IUserRepository
}

func NewUserListInteractor(userRepo repositories.IUserRepository) *UserListInteractor {
	return &UserListInteractor{
		userRepo: userRepo,
	}
}

func (r *UserListInteractor) Handle(in *inputs.UserListInputData) (int, *outputs.UserListOutputData, error) {
	users, err := r.userRepo.Find(in.ID, in.Username)
	if err != nil {
		return http.StatusNotFound, nil, errors.New("user is not found")
	}

	out := outputs.NewUserListOutputData(users)
	return http.StatusOK, out, nil
}

// Get

type UserGetInteractor struct {
	userRepo repositories.IUserRepository
}

func NewUserGetInteractor(userRepo repositories.IUserRepository) *UserGetInteractor {
	return &UserGetInteractor{
		userRepo: userRepo,
	}
}

func (r *UserGetInteractor) Handle(in *inputs.UserGetInputData) (int, *outputs.UserGetOutputData, error) {
	user, err := r.userRepo.GetById(in.ID)
	if err != nil {
		return http.StatusNotFound, nil, errors.New("user is not found")
	}

	out := outputs.NewUserGetOutputData(user)
	return http.StatusOK, out, nil
}

// Update

type UserUpdateInteractor struct {
	userRepo   repositories.IUserRepository
	jwtService *services.Jwt
}

func NewUserUpdateInteractor(userRepo repositories.IUserRepository, jwtService *services.Jwt) *UserUpdateInteractor {
	return &UserUpdateInteractor{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (r *UserUpdateInteractor) Handle(in *inputs.UserUpdateInputData) (int, *outputs.UserUpdateOutputData, error) {
	user, err := r.userRepo.GetById(in.ID)
	if err != nil {
		return http.StatusNotFound, nil, errors.New("user is not found")
	}

	user.SetUser(in.Username, in.Email, in.DisplayName, in.Bio)

	err = r.userRepo.DB().(*gorm.DB).Transaction(func(tx *gorm.DB) error {
		err := r.userRepo.Save(tx, user)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	out := outputs.NewUserUpdateOutputData(user)

	return http.StatusOK, out, nil
}

// Delete

type UserDeleteInteractor struct {
	userRepo   repositories.IUserRepository
	jwtService *services.Jwt
}

func NewUserDeleteInteractor(userRepo repositories.IUserRepository, jwtService *services.Jwt) *UserDeleteInteractor {
	return &UserDeleteInteractor{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (r *UserDeleteInteractor) Handle(in *inputs.UserDeleteInputData) (int, error) {
	user, err := r.userRepo.GetById(in.ID)
	if err != nil {
		return http.StatusNotFound, errors.New("user is not found")
	}

	err = r.userRepo.DB().(*gorm.DB).Transaction(func(tx *gorm.DB) error {
		err := r.userRepo.Delete(tx, user)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
