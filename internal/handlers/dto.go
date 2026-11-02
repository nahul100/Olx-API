package handlers

import "time"

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