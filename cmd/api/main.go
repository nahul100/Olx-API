package main 

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")//order matters here, set the header before writing the response
        w.WriteHeader(http.StatusOK)//if this is not set, the default status code is 200 OK, but it's good practice to set it explicitly.also write does it's task
		w.Write([]byte(`{"status": "healthy"}`))
	})
	srv:=http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 40 * time.Second,
		IdleTimeout: 60 * time.Second,
	}
	
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}