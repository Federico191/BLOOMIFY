package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTMakerItf interface {
	CreateToken(username string) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}

type JWTMaker struct {
	secretKey string
}

func NewJWT(secretKey string) JWTMaker {
	return JWTMaker{secretKey: secretKey}
}

func (j JWTMaker) CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j JWTMaker) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
