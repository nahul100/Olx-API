package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	//"os"
	"log/slog"
	
)

type listing struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	City        string `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListingHandler struct {
	db *sql.DB
    logger *slog.Logger
}
func NewListingHandler(db *sql.DB) *ListingHandler {
	return &ListingHandler{
		db: db,
		logger: slog.Default(),
	}
	}

func (lh *ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()//request scoped context, which can be used to cancel the query if the client disconnects or the request times out
	
		rows, err := lh.db.QueryContext(ctx,
			`SELECT id, title, description, price, city, created_at
		     FROM listings
		     ORDER BY created_at DESC
		     LIMIT 100`)
		if err != nil {
			lh.logger.Error("Error querying listings", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
       
		listings := []listing{}
		for rows.Next() {
           var l listing
		   if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
			   lh.logger.Error("Error scanning listing", "error", err)
			   http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			   return
		   }

		   lh.logger.Info("listings fetched", "total", len(listings))
		   listings = append(listings, l)
		}
		if err := rows.Err(); err != nil {
			log.Printf("Error iterating listings: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(listings)
	
}
func (lh *ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // request scoped context, which can be used to cancel the query if the client disconnects or the request times out
	id := r.PathValue("id")
	fmt.Println("id", id)

	lh.logger.Debug("Attempting to delete listing", "listing_id", id)
	lh.logger.Info("Delete request received", "listing_id", id)

	_, err := lh.db.ExecContext(ctx, "DELETE FROM listings WHERE id = $1", id)
	if err != nil {
		//log.Printf("Error deleting listing: %v", err)
		lh.logger.Error("delete failed", "listing_id", id, "error", err)
		//log.Error("delete failed", "listing_id", id, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

		w.WriteHeader(http.StatusNoContent)
		
		// Implement deletion logic here
	}

