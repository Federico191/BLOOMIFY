package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
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

	user := routerGroup.Group("/user")

	user.POST("/register", r.Handler.User.Register)
	user.POST("/login", r.Handler.User.Login)
	user.GET("/verify-email/:code", r.Handler.User.VerifyEmail)
	user.GET("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.GetUser)

	profile := routerGroup.Group("/profile")
	profile.POST("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.UpdatePhoto)

	service := routerGroup.Group("/service")

	beautyClinic := service.Group("/beauty-clinic")
	beautyClinic.GET("/search", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.GET("/:id", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.GET("/clinic-detail/:id", r.Handler.Service.GetById)
	//beautyClinic.POST("/clinic-detail/booking", r.Middleware.JwtAuthMiddleware, r.Handler.Booking.Create)

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

func (r *Route) Serve() {
	port := os.Getenv("PORT")

	err := r.Router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}
