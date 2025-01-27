package expense

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
	"gorm.io/gorm"
)

type Controller struct {
	ReaderDB *gorm.DB
	WriterDB *gorm.DB
}

// Create new Expense
func (con *Controller) Create(c *fiber.Ctx) error {
	// create expenses
	account, ok := c.Locals("account").(models.Accounts)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Account not found",
		})
	}
	// fmt.Println(account)
	var dto CreateExpense
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	expense := models.Expenses{
		Title:       dto.Title,
		Amount:      dto.Amount,
		Description: dto.Description,
		AccountId:   account.ID,
		OwnerSplit:  uint(dto.Amount),
	}

	con.WriterDB.Create(&expense)
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   expense,
	})
}

// GET
func (con *Controller) Get(c *fiber.Ctx) error {

	account, ok := c.Locals("account").(models.Accounts)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	var expenses []models.Expenses
	var count int64
	con.ReaderDB.Where("account_id = ?", account.ID).Preload("Splits").Order("created_at DESC").Find(&expenses)
	con.ReaderDB.Table("expenses").Where("account_id = ?", account.ID).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"count":   count,
			"expense": expenses,
		},
	})
}

// GET by Id
func (con *Controller) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id must be an integer in params",
		})
	}
	var expense models.Expenses
	con.ReaderDB.Find(&expense, id)

	if expense.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "expense not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   expense,
	})

}

func (con *Controller) DeleteById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id in params",
		})
	}
	var ex models.Expenses
	tx := con.WriterDB.Delete(ex, id)

	if tx.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to delete",
		})
	}

	if tx.RowsAffected != 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to delete, not found",
		})
	}

	fmt.Println(ex)
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": "successfully deleted",
		},
	})
}

func (con *Controller) UpdateById(c *fiber.Ctx) error {
	var dto UpdateExpense
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id in params",
		})

	}

	err = c.BodyParser(&dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse body.",
		})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var expense models.Expenses
	con.ReaderDB.Find(&expense, id)

	if expense.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "expense not found",
		})
	}
	tx := con.WriterDB.Table("expenses").Where("id = ?", id).Updates(dto)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to Update",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": "successfully updated.",
		},
	})
}
