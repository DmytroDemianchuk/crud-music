package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dmytrodemianchuk/crud-music/internal/config"
	"github.com/dmytrodemianchuk/crud-music/internal/repository/psql"
	"github.com/dmytrodemianchuk/crud-music/internal/service"
	"github.com/dmytrodemianchuk/crud-music/internal/transport/rest"
	"github.com/dmytrodemianchuk/crud-music/pkg/database"

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
	musicsRepo := psql.NewMusics(db)
	musicsService := service.NewMusics(musicsRepo)
	handler := rest.NewHandler(musicsService)

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
