package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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
	r.Router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	routerGroup := r.Router.Group("/api/v1")

	routerGroup.GET("/health-check", healthCheck)

	routerGroup.POST("/register", r.Handler.User.Register)
	routerGroup.POST("/login", r.Handler.User.Login)
	routerGroup.GET("/verify-email/:code", r.Handler.User.VerifyEmail)
	routerGroup.GET("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.GetUser)

	profile := routerGroup.Group("/profile")
	profile.POST("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.UpdatePhoto)

	service := routerGroup.Group("/service")

	beautyClinic := service.Group("/beauty-clinic")
	beautyClinic.GET("/search", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.GET("/:id", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.GET("/review", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.POST("/review", r.Middleware.JwtAuthMiddleware, r.Handler.Review.Create)
	beautyClinic.GET("/clinic-detail/:id", r.Handler.Service.GetById)

	salon := service.Group("/salon")
	salon.GET("/search", r.Handler.Service.GetAllSalon)
	salon.GET("/salon-detail/:id", r.Handler.Service.GetById)

	spaMassage := service.Group("/spa-massage")
	spaMassage.GET("/search", r.Handler.Service.GetAllSpaMassage)
	spaMassage.GET("/spa-massage-detail/:id", r.Handler.Service.GetById)

	fitnessCenter := service.Group("/fitness-center")
	fitnessCenter.GET("/search", r.Handler.Service.GetAllFitnessCenter)
	fitnessCenter.GET("/fitness-center-detail/:id", r.Handler.Service.GetById)

}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
