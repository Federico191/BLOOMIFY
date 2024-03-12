package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"projectIntern/pkg/config"
	"strconv"
	"time"
)

type JWTMakerItf interface {
	CreateToken(id uuid.UUID) (string, error)
	VerifyToken(tokenString string) (uuid.UUID, error)
}

type JWTMaker struct {
	env *config.Env
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func NewJWT(env *config.Env) JWTMaker {
	return JWTMaker{env: env}
}

func (j JWTMaker) CreateToken(id uuid.UUID) (string, error) {
	expired, err := strconv.Atoi(j.env.ExpiredToken)
	if err != nil {
		log.Fatalf("cannot converse expired")
	}

	claims := &Claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expired))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.env.SecretToken))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j JWTMaker) VerifyToken(tokenString string) (uuid.UUID, error) {
	var (
		claims Claims
		userId uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.env.SecretToken), nil
	})

	if err != nil {
		return userId, err
	}

	if !token.Valid {
		return userId, err
	}

	userId = claims.UserId

	return userId, nil
}
