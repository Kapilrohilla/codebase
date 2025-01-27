package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Controller struct {
	ReaderDb *gorm.DB
	WriterDb *gorm.DB
}

func (con *Controller) Login(c *fiber.Ctx) error {
	var body LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check if user exists
	var accModel models.Accounts
	con.ReaderDb.Where("email = ?", body.Email).First(&accModel)

	// extract password
	var hash = accModel.Password

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(body.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	// generate token
	var expiry_time int64 = time.Now().Add(time.Hour * 6).Unix()
	tokenString, err := GenerateToken(accModel.ID, "tracker", expiry_time)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
			"stack": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"type":       "Bearer",
			"token":      tokenString,
			"expired_at": expiry_time,
		},
	})
}

func (con *Controller) SentLoginLink(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{})
}

func GenerateToken(accountId uint, issuer string, expiry int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accId": accountId,
		"iss":   issuer,
		"exp":   expiry,
		"iat":   time.Now().Unix(),
	})
	return jwtToken.SignedString([]byte("SECRET"))
	// return jwtToken
}
