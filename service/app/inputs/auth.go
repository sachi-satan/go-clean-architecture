package inputs

import "github.com/go-playground/validator"

// SignUp

type AuthSignUpInputData struct {
	Username    string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,gte=8"`
	DisplayName string `validate:"required"`
	Bio         string `validate:"required"`
}

func NewAuthSignUpInputData(username string, email string, password string, displayName string, bio string) *AuthSignUpInputData {
	return &AuthSignUpInputData{
		Username:    username,
		Email:       email,
		Password:    password,
		DisplayName: displayName,
		Bio:         bio,
	}
}

func (r *AuthSignUpInputData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

// SignIn

type AuthSignInInputData struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=8"`
}

func NewAuthSignInInputData(email string, password string) *AuthSignInInputData {
	return &AuthSignInInputData{
		Email:    email,
		Password: password,
	}
}

func (r *AuthSignInInputData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
