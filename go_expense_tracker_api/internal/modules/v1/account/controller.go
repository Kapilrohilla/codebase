package account

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Controller struct {
	ReaderDb *gorm.DB
	WriterDb *gorm.DB
}

func (con *Controller) Get(c *fiber.Ctx) error {
	// get page, limit from query
	var query GetAccountQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// execute query
	offset, limit := utils.HandlePagination(query.Page, query.Limit)
	// return data
	var accounts []models.Accounts
	var count int64
	con.ReaderDb.Offset(offset).Limit(limit).Find(&accounts)
	con.ReaderDb.Table("accounts").Count(&count)

	var accountResponses []GetAccounts
	for _, account := range accounts {
		accountResponse := GetAccounts{
			ID:    account.ID,
			Name:  account.Name,
			Email: account.Email,
			Phone: account.Phone,
		}

		accountResponses = append(accountResponses, accountResponse)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":  accountResponses,
		"count": count,
	})
}

func (con *Controller) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": "account id must be a number",
		})
	}

	var account models.Accounts
	con.ReaderDb.Find(&account, id)

	if account.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"errors": "account not found",
		})
	}

	var responseAccount GetAccounts = GetAccounts{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
		Phone: account.Phone,
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": responseAccount,
	})
}

func (con *Controller) UpdateById(c *fiber.Ctx) error {
	// return c.SendString("update by id")
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	var updateDto UpdateAccount
	if err := c.BodyParser(&updateDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var accDetails models.Accounts

	// check if account exists
	tx := con.ReaderDb.Table("accounts").Find(&accDetails)

	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": tx.Error.Error(),
		})
	}

	if accDetails.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "account not found",
		})
	}
	var updateBody models.Accounts = models.Accounts{
		Name:  updateDto.Name,
		Email: updateDto.Email,
		Phone: updateDto.Phone,
	}

	if updateDto.Email != "" {
		updateBody.IsEmailVerified = false
	}
	if updateDto.Phone != "" {
		updateBody.IsPhoneVerified = false
	}
	// update account
	tx = con.WriterDb.Model(&updateBody).Where("id = ?", id).Updates(updateBody)

	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": tx.Error.Error(),
		})
	}

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"status": "success",
	})
}

func (con *Controller) DeleteById(c *fiber.Ctx) error {
	// return c.SendString("delete by id")
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	var account models.Accounts

	tx := con.WriterDb.Delete(&account, id)

	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": tx.Error.Error(),
		})
	}
	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"status": "success",
	})
}

func (con *Controller) Create(c *fiber.Ctx) error {
	var createDto Create
	// fetch data from request
	if err := c.BodyParser(&createDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// validation
	err := validator.New().Struct(&createDto)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// hash password
	hash, err := hashPassword(createDto.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// create model
	var account models.Accounts = models.Accounts{
		Name:     createDto.Name,
		Email:    createDto.Email,
		Phone:    createDto.Phone,
		Password: hash,
		IsAdmin:  false,
	}

	// execute query
	tx := con.WriterDb.Create(&account)

	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": tx.Error.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": createDto})
}

func (con *Controller) CreateAdmin(c *fiber.Ctx) error {

	var createAdmin CreateAdmin

	if err := c.BodyParser(&createAdmin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// validation
	err := validator.New().Struct(&createAdmin)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// create password
	hash, err := hashPassword(createAdmin.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// create model
	var admin models.Accounts = models.Accounts{
		Name:     createAdmin.Name,
		Email:    createAdmin.Email,
		Phone:    createAdmin.Phone,
		Password: hash,
		IsAdmin:  true,
	}

	tx := con.WriterDb.Create(&admin)

	if tx.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": tx.Error.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": createAdmin})

}

func (conn *Controller) GetProfile(c *fiber.Ctx) error {
	userDetails := c.Locals("account").(models.Accounts)

	var returnableData map[string]interface{} = make(map[string]interface{})
	returnableData["id"] = userDetails.ID
	returnableData["name"] = userDetails.Name
	returnableData["email"] = userDetails.Email
	returnableData["phone"] = userDetails.Phone
	returnableData["isAdimin"] = userDetails.IsAdmin

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   returnableData,
	})
}
func hashPassword(password string) (string, error) {
	pswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pswd), nil
}
