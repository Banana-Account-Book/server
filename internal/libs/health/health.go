package health

import "github.com/gofiber/fiber/v2"

func Check(r *fiber.App) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("pong")
	})
}
