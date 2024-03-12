package middleware

import (
	"github.com/gin-gonic/gin"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/jwt"
)

type MiddlewareItf interface {
	JwtAuthMiddleware(ctx *gin.Context)
	CorsMiddleware() gin.HandlerFunc
}

type Middleware struct {
	jwt     jwt.JWTMakerItf
	useCase usecase.UseCase
}

func (m Middleware) CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "True")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if ctx.Request.Method == "OPTIONS" {
			ctx.Abort()
		}

		ctx.Next()
	}
}

func NewMiddleware(jwt jwt.JWTMakerItf, useCase usecase.UseCase) MiddlewareItf {
	return &Middleware{jwt: jwt, useCase: useCase}
}
