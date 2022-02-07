package outputs

import (
	"backend/app/models"
)

// List

type UserListOutputData []*models.User

func NewUserListOutputData(users []*models.User) *UserListOutputData {
	return (*UserListOutputData)(&users)
}

// Get

type UserGetOutputData struct {
	*models.User
}

func NewUserGetOutputData(user *models.User) *UserGetOutputData {
	return &UserGetOutputData{
		User: user,
	}
}

// Update

type UserUpdateOutputData struct {
	*models.User
}

func NewUserUpdateOutputData(user *models.User) *UserUpdateOutputData {
	return &UserUpdateOutputData{
		User: user,
	}
}
