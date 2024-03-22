package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/jwt"
)

type MiddlewareItf interface {
	JwtAuthMiddleware(ctx *gin.Context)
	Cors() gin.HandlerFunc
}

type Middleware struct {
	jwt     jwt.JWTMakerItf
	useCase usecase.UseCase
}

func NewMiddleware(jwt jwt.JWTMakerItf, useCase usecase.UseCase) MiddlewareItf {
	return &Middleware{jwt: jwt, useCase: useCase}
}

func (m Middleware) Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")
		ctx.Header("Access-Control-Max-Age", "43200")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
