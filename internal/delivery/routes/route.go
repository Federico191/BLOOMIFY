package routes

import (
	"github.com/gin-gonic/gin"
	"projectIntern/internal/delivery/handler/rest"
	"projectIntern/internal/delivery/middleware"
)

type Route struct {
	Handler    *rest.Handler
	Router     *gin.Engine
	Middleware middleware.MiddlewareItf
}

func NewRoute(handler *rest.Handler, router *gin.Engine, Middleware middleware.MiddlewareItf) *Route {
	return &Route{Handler: handler, Router: router, Middleware: Middleware}
}

func (r *Route) MountEndPoint() {
	r.Middleware.CorsMiddleware()

	routerGroup := r.Router.Group("/api/v1")
	routerGroup.POST("/register", r.Handler.User.Register)
	routerGroup.POST("/login", r.Handler.User.Login)
	routerGroup.GET("/verify_email/:code", r.Handler.User.VerifyEmail)
	routerGroup.GET("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.GetUser)

	beautyClinic := routerGroup.Group("/beauty_clinic")
	beautyClinic.GET("/", r.Handler.Place.GetAllBeautyClinic)
	beautyClinic.GET("/:id", r.Handler.Place.GetAllBeautyClinic)
	beautyClinic.GET("/review", r.Handler.Place.GetAllBeautyClinic)
	beautyClinic.POST("/review", r.Middleware.JwtAuthMiddleware, r.Handler.Review.Create)

	salon := routerGroup.Group("/salon")
	salon.GET("/", r.Handler.Place.GetAllSalon)

	spaMassage := routerGroup.Group("/spa_massage")
	spaMassage.GET("/", r.Handler.Place.GetAllSpaMassage)

	fitnessCenter := routerGroup.Group("/fitness_center")
	fitnessCenter.GET("/", r.Handler.Place.GetAllFitnessCenter)
}
