package main

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"projectIntern/internal/delivery/handler/rest"
	"projectIntern/internal/delivery/middleware"
	"projectIntern/internal/delivery/routes"
	"projectIntern/internal/repository"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/database/mysql"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
	"projectIntern/pkg/supabase"
)

func main() {
	err := godotenv.Load("../app.env")
	envi := os.Getenv("ENV")
	if err != nil && envi == "" {
		log.Fatalf("cannot load env:%v", err)
	}

	db, err := mysql.DBInit()
	if err != nil {
		log.Fatalf("cannot initialize DB: %v", err)
	}

	mysql.Migration(db)

	//mysql.InitSeed(db)

	repo := repository.Init(db)

	client := supabasestorageuploader.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), os.Getenv("SUPABASE_BUCKET"))

	spbs := supabase.NewSupabaseStorage(client)

	jwtAuth := jwt.NewJWT()

	emailVerify := email.NewEmail()

	uc := usecase.Init(repo, jwtAuth, emailVerify, spbs)

	handler := rest.Init(uc)

	router := gin.Default()

	mdw := middleware.NewMiddleware(jwtAuth, *uc)

	route := routes.NewRoute(handler, router, mdw)

	route.MountEndPoint()

	err = router.Run(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("cannot run localhost: %v", err)
	}
}
