package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/middleware"
)

func New(router fiber.Router, dbs storage.DBs) {
	accountController := Controller{
		ReaderDb: &dbs.ReaderDB,
		WriterDb: &dbs.WriterDB,
	}

	router.Get("/", accountController.Get)
	router.Post("/", accountController.Create)
	router.Get("/profile", middleware.Auth(&dbs.ReaderDB), accountController.GetProfile)
	router.Get("/:id", accountController.GetById)
	router.Delete("/:id", accountController.DeleteById)
	router.Put("/:id", accountController.UpdateById)
	router.Post("/admin", accountController.CreateAdmin)
}
