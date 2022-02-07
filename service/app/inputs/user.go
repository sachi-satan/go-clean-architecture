package inputs

import (
	"github.com/go-playground/validator"
	"strings"
)

// List

type UserListInputData struct {
	ID       []string
	Username []string
}

func NewUserListInputData(id string, username string) *UserListInputData {
	return &UserListInputData{
		ID: strings.FieldsFunc(id, func(c rune) bool {
			return c == ','
		}),
		Username: strings.FieldsFunc(username, func(c rune) bool {
			return c == ','
		}),
	}
}

func (r *UserListInputData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

// Get

type UserGetInputData struct {
	ID string `validate:"required"`
}

func NewUserGetInputData(id string) *UserGetInputData {
	return &UserGetInputData{
		ID: id,
	}
}

func (r *UserGetInputData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

// Update

type UserUpdateInputData struct {
	ID          string
	Username    string
	Email       string `validate:"email"`
	DisplayName string
	Bio         string
}

func NewUserUpdateInputData(id string, username string, email string, displayName string, bio string) *UserUpdateInputData {
	return &UserUpdateInputData{
		ID:          id,
		Username:    username,
		Email:       email,
		DisplayName: displayName,
		Bio:         bio,
	}
}

func (r *UserUpdateInputData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

// Delete

type UserDeleteInputData struct {
	ID string `validate:"required"`
}

func NewUserDeleteInputData(id string) *UserDeleteInputData {
	return &UserDeleteInputData{
		ID: id,
	}
}

func (r *UserDeleteInputData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
