package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")//order matters here, set the header before writing the response
        w.WriteHeader(http.StatusOK)//if this is not set, the default status code is 200 OK, but it's good practice to set it explicitly.also write does it's task
		w.Write([]byte(`{"status": "healthy"}`))
	}