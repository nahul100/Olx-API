package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string
const (
	CodeInvalidID Code = "Invalid ID"
	CodeInternalError Code = "internal error"
	CodeNotFound Code = "not found"
	CodeMalformedJSON Code = "malformed json"
	CodeValidationError Code = "validation error"
	CodeUnauthenticated Code = "unauthenticated"
	Codeforbidden Code = "forbidden"
	CodeConflict Code = "conflict"
	CodeRateLimited Code = "rate limited"
)

type errorEnvelop struct{
	Error errorPayload `json:"error"`
}

type errorPayload struct{
	Code   Code   `json:"code"`
	Message   string  `json:"message"`
	Field   string  `json:"field,,omitempty"`
}

func Error(w http.ResponseWriter, status int, message string, code string) {
      w.Header().Set("Content-Type","application/json")
	  w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvelop{Error : errorPayload{
	    Code:   Code(code),
		Message:   message,
	}})

}
func ValidationError(w http.ResponseWriter, status int, message string, code string, field string) {
      w.Header().Set("Content-Type","application/json")
	  w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvelop{Error : errorPayload{
	    Code:   Code(code),
		Message:   message,
		Field:   field,
	}})

}