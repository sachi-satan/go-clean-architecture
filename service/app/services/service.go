package services

import (
	"github.com/joho/godotenv"
	"os"
)

type Service struct {
	Mysql *MySql
	Jwt   *Jwt
}

func NewService() (*Service, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	mysql, err := NewMySqlService(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	if err != nil {
		return nil, err
	}

	jwt, err := NewJwtService(
		os.Getenv("PRIVATE_KEY_FILE_PATH"),
		os.Getenv("PUBLIC_KEY_FILE_PATH"),
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		Mysql: mysql,
		Jwt:   jwt,
	}, nil
}
