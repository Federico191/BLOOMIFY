package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/pkg/response"
	"strings"
)

func (m Middleware) JwtAuthMiddleware(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "Authorization token is required", errors.New(""))
		ctx.JSON(http.StatusUnauthorized, gin.H{"customerrors": "Authorization token is required"})
		ctx.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwt.VerifyToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "Invalid token", err)
		ctx.Abort()
		return
	}

	user, err := m.useCase.User.GetById(userId)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to get user", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
