package middleware

import "github.com/gofiber/fiber/v2"

func Notfound(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "route not found",
	})
}
