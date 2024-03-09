package main

import (
	"github.com/gin-gonic/gin"
	email3 "github.com/jordan-wright/email"
	"log"
	"projectIntern/internal/delivery/handler/rest"
	"projectIntern/internal/delivery/routes"
	"projectIntern/internal/repository"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/config"
	"projectIntern/pkg/database/mysql"
	email2 "projectIntern/pkg/email"
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

	mysql.InitSeed(db)

	jwtAuth := jwt.NewJWT(env.SecretToken)

	email := email3.NewEmail()

	emailVerify := email2.NewEmail(email, env)

	uc := usecase.Init(repo, jwtAuth, emailVerify)

	handler := rest.Init(uc)

	router := gin.Default()

	route := routes.NewRoute(handler, router)

	route.MountEndPoint()

	err = router.Run(env.APort)
	if err != nil {
		log.Fatalf("cannot run localhost: %v", err)
	}
}
