package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/constants"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
	"gorm.io/gorm"
)

func Auth(readerDB *gorm.DB) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {

		var headers map[string][]string = c.GetReqHeaders()

		token := headers["Authorization"]

		if len(token) < 1 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Authroization header required",
			})
		}

		splittedToken := strings.Split(token[0], " ")

		if len(splittedToken) != 2 {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		if splittedToken[0] != constants.TOKEN_TYPE {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "invalid token type",
			})
		}

		isValid, claims, err := verifyToken(splittedToken[1])
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "invalid token, token might expired",
			})
		}

		if !isValid {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		accId := claims["accId"]
		var accountInstance models.Accounts
		readerDB.First(&accountInstance, accId)

		if accountInstance.ID == 0 {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "invalid account",
			})
		}
		c.Locals("account", accountInstance)
		return c.Next()
	}
}
func verifyToken(tokenString string) (bool, map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})

	if err != nil {
		return false, nil, err
	}

	if !token.Valid {
		return false, nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil, fmt.Errorf("invalid claims format")
	}
	claimData := make(map[string]interface{})
	for key, val := range claims {
		if val == nil {
			continue
		}
		claimData[key] = val
	}
	return true, claimData, nil
}
