package routes

import (
	"github.com/gin-gonic/gin"
	"projectIntern/internal/delivery/handler/rest"
)

type AuthRoute struct {
	Router      *gin.Engine
	AuthHandler *rest.AuthHandler
}

func (ar AuthRoute) Register() {
	Auth := ar.Router.Group("auth")
	Auth.POST("/register", ar.AuthHandler.Register)
	Auth.POST("/login", ar.AuthHandler.Login)

}
