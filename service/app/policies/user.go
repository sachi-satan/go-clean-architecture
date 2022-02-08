package policies

import (
	"backend/app/models"
	"backend/app/repositories"
	"errors"
)

type UserPolicy struct {
	userRepository *repositories.UserRepository
}

func NewUserPolicy(userRepository *repositories.UserRepository) *UserPolicy {
	return &UserPolicy{
		userRepository: userRepository,
	}
}

func (r *UserPolicy) Update(auth *models.User, id string) error {
	user, err := r.userRepository.GetById(id)
	if err != nil {
		return errors.New("your client does not have permission to the requested")
	}

	if auth.ID != user.ID {
		return errors.New("your client does not have permission to the requested")
	}

	return nil
}

func (r *UserPolicy) Delete(auth *models.User, id string) error {
	user, err := r.userRepository.GetById(id)
	if err != nil {
		return errors.New("your client does not have permission to the requested")
	}

	if auth.ID != user.ID {
		return errors.New("your client does not have permission to the requested")
	}

	return nil
}
