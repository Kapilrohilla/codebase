package split

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/middleware"
)

func New(router fiber.Router, dbs storage.DBs) {

	controller := Controller{
		ReaderDB: &dbs.ReaderDB,
		WriterDB: &dbs.WriterDB,
	}

	auth := middleware.Auth(&dbs.ReaderDB)

	router.Post("/", auth, controller.Create)
	router.Get("/", auth, controller.Get)
	router.Delete("/:id", auth, controller.DeleteById)
	// implementation pending
	router.Post("/delete/bulk", auth, controller.BulkDelete)

}
