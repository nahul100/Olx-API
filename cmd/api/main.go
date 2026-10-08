package main

import (
	"log"
	"net/http"
	"time"

	"Olx-API/internal"
	"Olx-API/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: no .env file found or could not be loaded: %v", err)
	}

	cfg := config.MustLoad() // load configuration from environment variables and .env file when present
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health) // register the health handler to the /health endpoint

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 40 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Printf("Starting server on port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}