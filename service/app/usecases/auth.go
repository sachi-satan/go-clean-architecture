package usecases

import (
	"backend/app/inputs"
	"backend/app/outputs"
)

type IAuthSignUpUseCase interface {
	Handle(in *inputs.AuthSignUpInputData) (int, *outputs.AuthSignUpOutputData, error)
}

type IAuthSignInUseCase interface {
	Handle(in *inputs.AuthSignInInputData) (int, *outputs.AuthSignInOutputData, error)
}
