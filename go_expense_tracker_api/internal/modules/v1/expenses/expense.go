package expense

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/middleware"
)

func New(router fiber.Router, dbs storage.DBs) {
	expenseController := Controller{
		ReaderDB: &dbs.ReaderDB,
		WriterDB: &dbs.WriterDB,
	}
	authMiddleware := middleware.Auth(&dbs.ReaderDB)
	router.Get("/", authMiddleware, expenseController.Get)
	router.Post("/", authMiddleware, expenseController.Create)
	router.Get("/:id", expenseController.GetById)
	router.Delete("/:id", expenseController.DeleteById)
	router.Put("/:id", expenseController.UpdateById)
}
