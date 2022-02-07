package repositories

import (
	"backend/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db interface{}
}

func NewUserRepository(db interface{}) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) DB() interface{} {
	return r.db
}

func (r *UserRepository) Find(id []string, username []string) ([]*models.User, error) {
	var users []*models.User

	d := r.db.(*gorm.DB)

	if len(id) > 0 {
		d = d.Where("id IN ?", id)
	}

	if len(username) > 0 {
		d = d.Where("username IN ?", username)
	}

	d = d.Find(&users)
	return users, d.Error
}

func (r *UserRepository) GetById(id string) (*models.User, error) {
	u := models.NewUser()
	d := r.db.(*gorm.DB).First(u, id)
	return u, d.Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	u := models.NewUser()
	d := r.db.(*gorm.DB).Where("email = ?", email).First(u)
	return u, d.Error
}

func (r *UserRepository) Save(tx interface{}, user *models.User) error {
	t := tx.(*gorm.DB).Save(user)
	return t.Error
}

func (r *UserRepository) Delete(tx interface{}, user *models.User) error {
	t := tx.(*gorm.DB).Delete(user)
	return t.Error
}
