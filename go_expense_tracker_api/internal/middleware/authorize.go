package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
)

var ROLES map[string]string = map[string]string{
	"admin": "admin",
	"user":  "user",
}

func Authorize(roleName string) func(*fiber.Ctx) error {
	var isAdminRole bool = roleName == ROLES["admin"]
	return func(c *fiber.Ctx) error {

		account := c.Locals("account").(models.Accounts)

		if account.IsAdmin != isAdminRole {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		return c.Next()
	}
}
