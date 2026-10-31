package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/anoopraju31/olx-api/internal/config"
	"github.com/anoopraju31/olx-api/internal/db"
	"github.com/anoopraju31/olx-api/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	db, dbErr := db.Connect(cfg.DatabaseUrl)
	if dbErr != nil {
		log.Fatalf("main.db.connect: %v", dbErr)
	}

	fmt.Println("DATABASE CONNECTED!")
	fmt.Println("SERVER is STARTING...")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listings", handlers.List(db))

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("SERVER RUNNING on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("SERVER FAILED: %v", err.Error())
	}
}
