package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          int        `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"-"`
	Password    string     `json:"-"`
	DisplayName string     `json:"displayName"`
	Bio         string     `json:"bio"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

func NewUser() *User {
	return &User{
		ID:          0,
		Username:    "",
		Email:       "",
		DisplayName: "",
		Bio:         "",
		CreatedAt:   time.Time{},
		UpdatedAt:   nil,
		DeletedAt:   nil,
	}
}

func (r *User) SetUserID(id int) {
	r.ID = id
}

func (r *User) SetUser(userName string, email string, displayName string, bio string) {
	r.Username = userName
	r.Email = email
	r.DisplayName = displayName
	r.Bio = bio
}

func (r *User) SetNewUser(userName string, email string, password string, displayName string, bio string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	r.Username = userName
	r.Email = email
	r.Password = string(hashed)
	r.DisplayName = displayName
	r.Bio = bio

	return nil
}

func (r *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(password))
}
