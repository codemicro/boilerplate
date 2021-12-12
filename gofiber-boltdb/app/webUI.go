package app

import "github.com/gofiber/fiber/v2"

func (e *endpoints) uiIndex(c *fiber.Ctx) error {
	return c.SendString("Mars image hosting! :D")
}
