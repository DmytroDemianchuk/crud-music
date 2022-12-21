package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dmytrodemianchuk/crud-music/internal/config"
	"github.com/dmytrodemianchuk/crud-music/internal/repository/psql"
	"github.com/dmytrodemianchuk/crud-music/internal/service"
	grpc_client "github.com/dmytrodemianchuk/crud-music/internal/transport/grpc"
	"github.com/dmytrodemianchuk/crud-music/internal/transport/rest"
	"github.com/dmytrodemianchuk/crud-music/pkg/database"
	"github.com/dmytrodemianchuk/crud-music/pkg/hash"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	hasher := hash.NewSHA1Hasher("salt")

	booksRepo := psql.NewMUsics(db)
	booksService := service.NewMusics(booksRepo)

	usersRepo := psql.NewUsers(db)
	tokensRepo := psql.NewTokens(db)

	auditClient, err := grpc_client.NewClient(9000)
	if err != nil {
		log.Fatal(err)
	}

	usersService := service.NewUsers(usersRepo, tokensRepo, auditClient, hasher, []byte("sample secret"))

	handler := rest.NewHandler(booksService, usersService)

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
