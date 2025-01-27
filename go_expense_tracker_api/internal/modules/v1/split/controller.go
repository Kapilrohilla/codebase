package split

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/utils"
	"gorm.io/gorm"
)

type Controller struct {
	ReaderDB *gorm.DB
	WriterDB *gorm.DB
}

func (conn *Controller) Get(c *fiber.Ctx) error {
	var query GetQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid query, failed to parse",
			"reason": err.Error(),
		})
	}
	offset, limit := utils.HandlePagination(int(query.Page), int(query.Limit))

	var splits models.Split
	if err := conn.ReaderDB.Offset(offset).Limit(limit).Find(&splits); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "failed to fetch data",
			"reason": err.Error,
		})
	}
	var count int64
	if err := conn.ReaderDB.Table("splits").Count(&count); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "failed to fetch data(count)",
			"reason": err.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"count":  count,
			"splits": splits,
		},
	})
}

func (conn *Controller) Create(c *fiber.Ctx) error {
	var body CreateSplit
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "payload is in invalid format",
		})
	}
	var accountDetail models.Accounts = c.Locals("account").(models.Accounts)
	fmt.Println(body)
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var expenses models.Expenses
	if err := conn.ReaderDB.Where("id = ? and account_id = ? ", body.ExpenseId, accountDetail.ID).First(&expenses, models.Expenses{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  fmt.Sprintf("failed to search expense detail with id #%d", body.ExpenseId),
			"reason": err.Error,
		})
	}

	if expenses.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "expense detail not found",
		})
	}

	if expenses.OwnerSplit < body.Amount {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "split amount must be less then expense amount",
		})
	}

	var split models.Split = models.Split{
		ExpenseId:   body.ExpenseId,
		SplitAmount: body.Amount,
		AccountId:   body.AccountId,
	}
	var splitDetail models.Split
	if err := conn.WriterDB.Limit(1).Find(&splitDetail, models.Split{
		ExpenseId: split.ExpenseId,
		AccountId: split.AccountId,
	}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "failed to find already exiting split detail",
			"reason": err.Error(),
		})
	}

	if splitDetail.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "expense already splitted with given account.",
		})
	}

	if err := conn.WriterDB.Create(&split).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to create splits",
		})
	}
	if err := conn.WriterDB.Table("expenses").Where("id = ?", body.ExpenseId).Update("owner_split", gorm.Expr("owner_split - ?", body.Amount)).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "failed to update owner_split",
			"reason": err.Error(),
		})

	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"split": split},
	})

}
func (conn *Controller) DeleteById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse split id",
		})
	}
	var exisingSplit models.Split
	if err := conn.ReaderDB.First(&exisingSplit, id).Error; err != nil {
		return c.Status(200).JSON(fiber.Map{
			"error": "split not found to delete",
		})
	}

	var splitDetail models.Split
	if err := conn.WriterDB.Table("splits").Delete(&splitDetail, &id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to delete split with id #%d", id),
		})
	}

	if err := conn.WriterDB.Table("expenses").Where("id = ?", exisingSplit.ExpenseId).Update("owner_split", gorm.Expr("owner_split + ?", exisingSplit.SplitAmount)).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "failed to update amount in expense detail",
			"reason": err.Error,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"message": "successfully deleted",
		},
	})
}

func (conn *Controller) BulkDelete(c *fiber.Ctx) error {

	return c.SendStatus(500)
}
