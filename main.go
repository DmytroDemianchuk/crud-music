package main

import (
	"github.com/dmytrodemianchuk/crud-music/music"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/music", music.GetMusics)
	app.Get("/api/v1/music:id", music.GetMusic)
	app.Post("/api/v1/music", music.NewMusic)
	app.Delete("/api/v1/music/:id", music.DeleteBook)
}
func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
