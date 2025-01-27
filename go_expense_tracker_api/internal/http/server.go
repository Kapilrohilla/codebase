package http

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kapilrohilla/go_expense_tracker_api/internal/middleware"
)

func New(myEnv map[string]string) (*fiber.App, func(*fiber.App, serverConfig), func(*fiber.App), serverConfig) {
	app := fiber.New()

	serverPort, err := strconv.ParseInt(myEnv["PORT"], 10, 16)
	if err != nil {
		log.Fatalf("Failed to parse PORT ENV variable %s", err.Error())
		panic(err)
	}

	var sCfg serverConfig = serverConfig{
		port: int(serverPort),
	}
	app.Use(middleware.Cors())
	app.Use(middleware.Logger())
	return app, startServer, stopServer, sCfg
}

type serverConfig struct {
	port int
}

func startServer(app *fiber.App, sCfg serverConfig) {
	log.Printf("Starting server at %d", sCfg.port)
	app.Listen(":" + strconv.Itoa(sCfg.port))
	log.Printf("Server started at %d", sCfg.port)
}

func stopServer(app *fiber.App) {
	err := app.Shutdown()
	if err != nil {
		log.Fatalf("Failed to stop server %s", err.Error())
	} else {
		log.Println("Server stopped successfully")
	}

}
