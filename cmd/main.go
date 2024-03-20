package main

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	midtrans2 "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
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
	"projectIntern/pkg/midtrans"
	"projectIntern/pkg/supabase"
)

func main() {
	err := godotenv.Load("../.env")
	envi := os.Getenv("ENV")
	if err != nil && envi == "" {
		log.Fatalf("cannot load env:%v", err)
	}

	db, err := mysql.DBInit()
	if err != nil {
		log.Fatalf("cannot initialize DB: %v", err)
	}

	mysql.Migration(db)

	repo := repository.Init(db)

	client := supabasestorageuploader.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), os.Getenv("SUPABASE_BUCKET"))

	spbs := supabase.NewSupabaseStorage(client)

	jwtAuth := jwt.NewJWT()

	emailVerify := email.NewEmail()

	midtransCoreApi := coreapi.Client{ServerKey: os.Getenv("SERVER_KEY"), Env: midtrans2.Sandbox}

	mtrans := midtrans.NewMidtrans(midtransCoreApi)

	uc := usecase.Init(repo, jwtAuth, emailVerify, spbs, mtrans)

	handler := rest.Init(uc)

	router := gin.Default()

	mdw := middleware.NewMiddleware(jwtAuth, *uc)

	route := routes.NewRoute(handler, router, mdw)

	route.MountEndPoint()

	route.Serve()
}
