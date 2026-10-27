package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/anoopraju31/olx-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println("SERVER is STARTING...")

	mux := http.NewServeMux()

	// Health route
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"message": "All Good!"}`))
	})

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
