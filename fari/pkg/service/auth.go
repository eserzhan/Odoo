package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"time"

	"github.com/eserzhan/rest"
	"github.com/eserzhan/rest/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
)

const (
	salt = "3p4tm24t1fdvsdk,v.g,rlw,gs"
	signingKey = "qrkjk#4#@35FSFJlja#4353KSFjH"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user todo.User) (int, error){
	user.Password = hashPasswordWithSalt(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	user_id, err := a.repo.GetUser(username, hashPasswordWithSalt(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
	user_id,})

	tokenString, err := token.SignedString([]byte(signingKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (a *AuthService) ParseToken(accessToken string) (int, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil 
}

func hashPasswordWithSalt(password string) string {
    hash := sha1.New()
    hash.Write([]byte(password + salt))
    hashBytes := hash.Sum(nil)
    return hex.EncodeToString(hashBytes)
}