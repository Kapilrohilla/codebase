package cmd

import (
	"github.com/kapilrohilla/go_expense_tracker_api/internal/config"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/http"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/infra/storage"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/models"
)

func Server() {
	// load configs
	var myEnv map[string]string = config.Load()
	// load db
	var dbs storage.DBs = storage.New(myEnv)
	// migrate models
	models.MigrateModels(dbs)
	// start app
	app, startServer, _, cwg := http.New(myEnv)
	// initiate routes
	http.InitiateRoutes(dbs, app)

	// listen server
	startServer(app, cwg)
	// gracefully shutdown
}
