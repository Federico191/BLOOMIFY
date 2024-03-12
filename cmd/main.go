package main

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"log"
	"projectIntern/internal/delivery/handler/rest"
	"projectIntern/internal/delivery/middleware"
	"projectIntern/internal/delivery/routes"
	"projectIntern/internal/repository"
	"projectIntern/internal/usecase"
	"projectIntern/pkg"
	"projectIntern/pkg/config"
	"projectIntern/pkg/database/mysql"
	"projectIntern/pkg/email"
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

	client := supabasestorageuploader.New(env.SupabaseUrl, env.SupabaseKey, env.SupabaseBucket)

	supabase := pkg.NewSupabaseStorage(client)

	jwtAuth := jwt.NewJWT(env)

	emailVerify := email.NewEmail(env)

	uc := usecase.Init(repo, jwtAuth, emailVerify, supabase)

	handler := rest.Init(uc)

	router := gin.Default()

	mdw := middleware.NewMiddleware(jwtAuth, *uc)

	route := routes.NewRoute(handler, router, mdw)

	route.MountEndPoint()

	err = router.Run(env.APort)
	if err != nil {
		log.Fatalf("cannot run localhost: %v", err)
	}
}
