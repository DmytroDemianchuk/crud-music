package main

import (
	"fmt"
	"github.com/dmytrodemianchuk/crud-music/internal/repository/psql"
	"github.com/dmytrodemianchuk/crud-music/internal/service"
	"log"
	"net/http"
	"os"

	grpc_client "github.com/dmytrodmeianchuk/audit-log"

	"github.com/sirupsen/logrus"

	"github.com/dmytrodemianchuk/crud-music/pkg/config"
	"github.com/dmytrodemianchuk/crud-music/pkg/database"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	//srv := gin.New()
	//
	//cfg, err := config.Parse()
	//if err != nil {
	//	logrus.Fatalf("error psring config: %s", err.Error())
	//}

	//init db
	db, err := database.CreateConn(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.SSLMode)
	if err != nil {
		logrus.Fatalf("failed to connection db: %s", err.Error())
	}

	//init deps
	hasher := hash.NewSHA1Hasher("salt")

	musicRepository := psql.NewMusic(db)
	musicsService := service.NewMusic(musicRepository)

	usersRepo := psql.NewUsers(db)
	tokensRepo := psql.Newtokens(db)

	auditClient, err := grpc_client.NewClient(9000)
	if err != nil {
		log.Fatal(err)
	}

	usersService := service.NewUsers(usersRepo, tokensRepo, auditClient, hasher, []byte("sample secret"))

	handler := rest.NewHandler(musicsService usersService)

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	logrus.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	//musicTransport := rest.NewMusic(musicService)
	//
	//srv.GET("/musics", musicTransport.List)
	//srv.GET("/music/:id", musicTransport.Get)
	//srv.POST("/music", musicTransport.Create)
	//srv.PUT("/music/:id", musicTransport.Update)
	//srv.DELETE("/music/:id", musicTransport.Delete)
	//
	//if err := srv.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
	//	logrus.Fatalf("error occured while running http server %s", err.Error())
	//}
}
