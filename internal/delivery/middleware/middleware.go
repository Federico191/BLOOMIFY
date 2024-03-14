package middleware

import (
	"github.com/gin-gonic/gin"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/jwt"
)

type MiddlewareItf interface {
	JwtAuthMiddleware(ctx *gin.Context)
}

type Middleware struct {
	jwt     jwt.JWTMakerItf
	useCase usecase.UseCase
}

func NewMiddleware(jwt jwt.JWTMakerItf, useCase usecase.UseCase) MiddlewareItf {
	return &Middleware{jwt: jwt, useCase: useCase}
}
