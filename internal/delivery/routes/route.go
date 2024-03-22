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
	"time"
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
		AllowOrigins:     []string{"https://intern-bloomify.vercel.app"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	routerGroup := r.Router.Group("/api/v1")

	routerGroup.GET("/health-check", healthCheck)

	user := routerGroup.Group("/user")

	user.POST("/register", r.Handler.User.Register)
	user.POST("/login", r.Handler.User.Login)
	user.GET("/verify-email/:code", r.Handler.User.VerifyEmail)
	user.GET("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.GetUser)
	user.POST("/survey", r.Middleware.JwtAuthMiddleware, r.Handler.Personalization.Analyze)
	user.GET("/survey/result", r.Middleware.JwtAuthMiddleware, r.Handler.Product.GetByProblem)

	profile := routerGroup.Group("/profile")
	profile.POST("/", r.Middleware.JwtAuthMiddleware, r.Handler.User.UpdatePhoto)

	service := routerGroup.Group("/service")
	service.POST("/treatment/booking/", r.Middleware.JwtAuthMiddleware, r.Handler.Booking.CreateBookingTreatment)
	service.POST("/doctor/booking/", r.Middleware.JwtAuthMiddleware, r.Handler.Booking.CreateBookingDoctor)
	service.GET("/booking/get-status/:id", r.Handler.Booking.GetById)
	service.POST("/payment/update", r.Handler.Booking.Update)

	beautyClinic := service.Group("/beauty-clinic")
	beautyClinic.GET("/search", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.GET("/:id", r.Handler.Service.GetAllBeautyClinic)
	beautyClinic.GET("/clinic-detail/:id", r.Handler.Service.GetById)

	doctor := service.Group("/doctor")
	doctor.GET("/search", r.Handler.Doctor.GetAll)
	doctor.GET("/doctor-detail/:id", r.Handler.Doctor.GetById)

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
