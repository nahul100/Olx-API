package main

import (
	"log"
	"net/http"
	"time"
	"os"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/nahul100/Olx-API/internal/config"
	"github.com/nahul100/Olx-API/internal/db"

	"github.com/nahul100/Olx-API/internal/handlers"

)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: no .env file found or could not be loaded: %v", err)
	}
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	handler := slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{
		AddSource: true,
        Level : slog.LevelInfo,
	})
	logger:=slog.New(handler)
	slog.SetDefault(logger)
    
	lh:= handlers.NewListingHandler(db, logger)
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)
    mux.HandleFunc("GET /listings", lh.List)//
    mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

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