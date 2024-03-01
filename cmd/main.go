package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	handler2 "projectIntern/internal/delivery/handler"
	"projectIntern/internal/delivery/routes"
	"projectIntern/internal/repository"
	"projectIntern/internal/usecase"
	"projectIntern/internal/util/config"
	"projectIntern/internal/util/token"
)

func main() {
	env, err := config.NewEnv("../")
	valid := validator.New()
	if err != nil {
		log.Fatal("cannot")
	}
	db, err := config.DBInit(env)
	repo := repository.Init(db)
	jwt := token.NewJWT(env.SecretToken)
	uc := usecase.Init(repo, jwt)
	handler := handler2.Init(uc, valid)
	router := gin.Default()
	authRoute := routes.AuthRoute{Router: router, AuthHandler: handler.Auth}
	authRoute.Register()

	err = router.Run(":8000")
	if err != nil {
		log.Fatal("cannot run localhost")
	}
}
