package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/util/token"
)

func JwtAuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		userClaims, err := tokenMaker.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("userClaims", userClaims)
		c.Next()
		return
	}
}
