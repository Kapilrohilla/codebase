package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kapilrohilla/codebase/config"
)

func main() {
	fmt.Println("Welcome to Student API")
	// load config
	cfg := config.MustLoad()
	//database setup

	// server setup

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server %s\n", err.Error())
		panic(err)
	} else {
		log.Println("Server is running on", cfg.HTTPServer.Addr)
	}

}
