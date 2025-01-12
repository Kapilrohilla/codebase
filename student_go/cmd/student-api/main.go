package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kapilrohilla/codebase/internal/config"
	"github.com/kapilrohilla/codebase/internal/http/handlers/student"
	"github.com/kapilrohilla/codebase/internal/storage/sqlite"
)

func main() {
	fmt.Println("Welcome to Student API")
	// load config
	cfg := *config.MustLoad()
	//database setup
	db, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
		slog.Error("failed to start db", err.Error())
		panic("failed to initiate db connection.")
	}
	// server setup

	router := http.NewServeMux()

	router.HandleFunc("/api/v1/students", student.New(db))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server %s\n", err.Error())
			panic(err)
		} else {
			log.Println("Server is running on", cfg.HTTPServer.Addr)
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("shutting down the server")

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
}
