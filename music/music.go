package music

import (
	"github.com/gofiber/fiber/v2"
)

func GetMusics(c *fiber.Ctx) {
	c.Send("All musics")
}

func GetMusic(c *fiber.Ctx) {
	c.Send("A single music")
}

func NewMusic(c *fiber.Ctx) {
	c.Send("Add music")
}

func DeleteMusic(c *fiber.Ctx) {
	c.Send("Delete music")
}

func GetMusics(c *fiber.Ctx) {
	c.Send("All musics")
}
