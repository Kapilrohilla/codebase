package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/middleware"
	v1 "github.com/kapilrohilla/go_expense_tracker_api/internal/modules/v1"
)

func InitiateRoutes(dbs storage.DBs, app *fiber.App) {

	app.Get("/metrics", middleware.Monitor())
	routes := app.Group("/api")

	// v1
	v1Router := routes.Group("/v1")
	v1.New(v1Router, dbs)
	// v2

	// not found
	app.Use(middleware.Notfound)

}

// func loadModules() {

// }
