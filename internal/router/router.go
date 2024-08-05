package router

import (
	"banana-account-book.com/internal/libs/health"
	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	health.Check(r)
}
