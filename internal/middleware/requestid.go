package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)
type ctxKey int
     
const requestIDkey ctxKey = iota
const RequestIDHeader = "X-Request-ID"
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      id:=r.Header.Get(RequestIDHeader)
	  if id == "" {
        id = uuid.NewString()
	  }
	  w.Header().Set(RequestIDHeader, id)
	  ctx := context.WithValue(r.Context(), "requestIDkey", id)
	  //context.WithValue(r.Context(), "requestIDkey", id) another packages
	  next.ServeHTTP(w, r.WithContext(ctx))
})
}

func RequestIDFromContext(ctx context.Context) string{
	requestID := ctx.Value(requestIDkey).(string)
	return requestID
}