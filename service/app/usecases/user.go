package usecases

import (
	"backend/app/inputs"
	"backend/app/outputs"
)

type IUserListUseCase interface {
	Handle(in *inputs.UserListInputData) (int, *outputs.UserListOutputData, error)
}

type IUserGetUseCase interface {
	Handle(in *inputs.UserGetInputData) (int, *outputs.UserGetOutputData, error)
}

type IUserUpdateUseCase interface {
	Handle(in *inputs.UserUpdateInputData) (int, *outputs.UserUpdateOutputData, error)
}

type IUserDeleteUseCase interface {
	Handle(in *inputs.UserDeleteInputData) (int, error)
}
