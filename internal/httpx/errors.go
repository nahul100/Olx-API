package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string
const (
	CodeInvalidID Code = "Invalid ID"
	CodeInternalError Code = "internal error"
)

type errorEnvelop struct{
	Error errorPayload `json:"error"`
}

type errorPayload struct{
	Code   Code   `json:"code"`
	Message   string  `json:"message"`
}

func Error(w http.ResponseWriter, status int, message string, code string) {
      w.Header().Set("Content-Type","application/json")
	  w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvelop{Error : errorPayload{
	    Code:   Code(code),
		Message:   message,
	}})

}