package routes

import (
	"github.com/gin-gonic/gin"
	"projectIntern/internal/delivery/handler/rest"
)

type Route struct {
	Handler *rest.Handler
	Router  *gin.Engine
}

func NewRoute(handler *rest.Handler, router *gin.Engine) *Route {
	return &Route{Handler: handler, Router: router}
}

func (r *Route) MountEndPoint() {
	routerGroup := r.Router.Group("/api/v1")
	routerGroup.POST("/register", r.Handler.Auth.Register)
	routerGroup.POST("/login", r.Handler.Auth.Login)
	routerGroup.GET("/verify_email/:code", r.Handler.Auth.VerifyEmail)

	place := routerGroup.Group("/place")
	place.GET("/", r.Handler.Place.GetAll)
	//place.GET("/", r.Handler.Place.GetByCity)
	//place.GET("/", r.Handler.Place.GetByTreatment)
}
