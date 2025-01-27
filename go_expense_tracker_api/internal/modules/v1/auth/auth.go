package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
)

func New(authRouter fiber.Router, dbs storage.DBs) {

	controller := Controller{
		ReaderDb: &dbs.ReaderDB,
		WriterDb: &dbs.WriterDB,
	}

	authRouter.Post("/login", controller.Login)
}
