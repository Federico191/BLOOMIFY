package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
	"projectIntern/internal/entity"
	"strconv"
	"time"
)

type JWTMakerItf interface {
	CreateToken(id uuid.UUID) (string, error)
	VerifyToken(tokenString string) (uuid.UUID, error)
	GetLoginUser(ctx *gin.Context) (*entity.User, error)
}

type JWTMaker struct {
}

func (j JWTMaker) GetLoginUser(ctx *gin.Context) (*entity.User, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return &entity.User{}, errors.New("failed to get user")
	}

	return user.(*entity.User), nil
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func NewJWT() JWTMakerItf {
	return JWTMaker{}
}

func (j JWTMaker) CreateToken(id uuid.UUID) (string, error) {
	expired, err := strconv.Atoi(os.Getenv("EXPIRED_TOKEN"))
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
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
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
		return []byte(os.Getenv("SECRET_TOKEN")), nil
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
