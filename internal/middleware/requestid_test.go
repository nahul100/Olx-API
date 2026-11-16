package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDPropagatesToContextAndResponse(t *testing.T) {
	tests := []struct {
		name       string
		incomingID string
	}{
		{name: "generates an ID when absent"},
		{name: "preserves a supplied ID", incomingID: "client-request-123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var contextID string
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				contextID = RequestIDFromContext(r.Context())
			})
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.incomingID != "" {
				req.Header.Set(RequestIDHeader, tt.incomingID)
			}
			rec := httptest.NewRecorder()

			RequestID(next).ServeHTTP(rec, req)

			if contextID == "" {
				t.Fatal("request ID was not available from the request context")
			}
			if got := rec.Header().Get(RequestIDHeader); got != contextID {
				t.Fatalf("response request ID = %q, want %q", got, contextID)
			}
			if tt.incomingID != "" && contextID != tt.incomingID {
				t.Fatalf("context request ID = %q, want supplied ID %q", contextID, tt.incomingID)
			}
		})
	}
}
