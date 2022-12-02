package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/dmytrodemianchuk/crud-music/internal/repository"
	"github.com/dmytrodemianchuk/crud-music/internal/service"
	"github.com/dmytrodemianchuk/crud-music/internal/transport/rest"
	"github.com/dmytrodemianchuk/crud-music/pkg/config"
	"github.com/dmytrodemianchuk/crud-music/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.New()

	cfg, err := config.Parse()
	if err != nil {
		logrus.Fatalf("error psring config: %s", err.Error())
	}

	db, err := database.CreateConn(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.SSLMode)
	if err != nil {
		logrus.Fatalf("failed to connection db: %s", err.Error())
	}

	musicRepository := repository.NewMovie(db)
	musicService := service.NewMovie(musicRepository)
	movieTransport := rest.NewMovie(musicService)

	srv.GET("/movies", movieTransport.List)
	srv.GET("/movie/:id", movieTransport.Get)
	srv.POST("/movie", movieTransport.Create)
	srv.PUT("/movie/:id", movieTransport.Update)
	srv.DELETE("/movie/:id", movieTransport.Delete)

	if err := srv.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		logrus.Fatalf("error occured while running http server %s", err.Error())
	}
}
