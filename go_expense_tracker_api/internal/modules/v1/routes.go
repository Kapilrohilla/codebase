package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/modules/v1/account"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/modules/v1/auth"
	expense "github.com/kapilrohilla/go_expense_tracker_api/internal/modules/v1/expenses"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/modules/v1/split"
)

func New(v1 fiber.Router, dbs storage.DBs) {
	// account routes
	accountRouter := v1.Group("/account")
	account.New(accountRouter, dbs)

	// auth routes
	authRouter := v1.Group("/auth")
	auth.New(authRouter, dbs)

	// expense routes
	expenseRouter := v1.Group("/expense")
	expense.New(expenseRouter, dbs)

	// split routes
	splitRouter := v1.Group("/split")
	split.New(splitRouter, dbs)

	// response middleware
}
