package repositories

import (
	"backend/app/models"
)

type IUserRepository interface {
	DB() interface{}
	Find(id []string, username []string) ([]*models.User, error)
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Save(tx interface{}, user *models.User) error
	Delete(tx interface{}, user *models.User) error
}
