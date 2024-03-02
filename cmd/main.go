package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"projectIntern/internal/delivery/handler/rest"
	"projectIntern/internal/delivery/routes"
	"projectIntern/internal/repository"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/config"
	"projectIntern/pkg/database/mysql"
	"projectIntern/pkg/jwt"
)

func main() {
	env, err := config.NewEnv("../")
	if err != nil {
		log.Fatalf("cannot load env: %v", err)
	}

	db, err := mysql.DBInit(env)
	if err != nil {
		log.Fatalf("cannot initialize DB: %v", err)
	}

	mysql.Migration(db)

	repo := repository.Init(db)

	jwtAuth := jwt.NewJWT(env.SecretToken)

	uc := usecase.Init(repo, jwtAuth)

	handler := rest.Init(uc)

	router := gin.Default()

	authRoute := routes.AuthRoute{Router: router, AuthHandler: handler.Auth}
	authRoute.Register()

	err = router.Run(env.APort)
	if err != nil {
		log.Fatalf("cannot run localhost: %v", err)
	}
}
