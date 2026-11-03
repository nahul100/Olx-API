package handlers

import (
	"fmt"
	"strings"
	"time"
)




type CreateListing struct {
	Title       string `json:"title"`
	Description string `json:"description"`	
	Price       string `json:"price"`
	City        string `json:"city"`
}

type CreateListingResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`	
	CreatedAt   time.Time `json:"created_at"`
}

type ValidationError struct {
	Field   string 
	Message string 
}
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)	
}

func (req CreateListing) Validate() error {
	if strings.TrimSpace(req.Title) == "" {
		return &ValidationError{
			Field:   "title",
			Message: "Title is required",
		}
	}
	return nil
}