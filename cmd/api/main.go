package main

import (
	"Olx-API/internal/config"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
  cfg := config.MustLoad()
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")//order matters here, set the header before writing the response
        w.WriteHeader(http.StatusOK)//if this is not set, the default status code is 200 OK, but it's good practice to set it explicitly.also write does it's task
		w.Write([]byte(`{"status": "healthy"}`))
	})
	srv := &http.Server{
		Addr: ":" + cfg.Port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 40 * time.Second,
		IdleTimeout: 60 * time.Second,
	}
	log.Printf("Starting server on port %s", srv.Addr) ;
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}