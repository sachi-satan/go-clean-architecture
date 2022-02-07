package services

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"time"
)

type Jwt struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

type CustomClaims struct {
	UserID string `json:"uid"`
	jwt.StandardClaims
}

func NewJwtService(privKeyPath string, pubKeyPath string) (*Jwt, error) {
	signBytes, err := ioutil.ReadFile(privKeyPath)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return &Jwt{
		signKey:   signKey,
		verifyKey: verifyKey,
	}, nil
}

func (r *Jwt) GenToken(userId string) (string, error) {
	claims := &CustomClaims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(r.signKey)
	if err != nil {
		return "", err
	}

	return ss, err
}

func (r *Jwt) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return r.verifyKey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", errors.New("token is invalid format")
	}

	return claims.UserID, nil
}
