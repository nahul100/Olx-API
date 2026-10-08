package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/nahul100/Olx-API/internal"
	"github.com/nahul100/Olx-API/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: no .env file found or could not be loaded: %v", err)
	}

	cfg := config.MustLoad()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)

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