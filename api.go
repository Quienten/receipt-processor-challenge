package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// receipts/process
// This route is used to calculate the points for a receipt and store for later retrieval.
func CalculateReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Try to decode the request body into a Receipt struct
	// If the request body is invalid, return a 400 Bad Request
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	// Generate a new id for the receipt and calculate the points
	receiptUUID := uuid.New()
	uuidToPoint[receiptUUID.String()] = CalculatePoints(receipt)             // Store in local memory, not persistent
	json.NewEncoder(w).Encode(map[string]string{"id": receiptUUID.String()}) // Return the id to the client
}

// receipts/{id}/points
// {id} the receipt id to fetch.
// This route is used to fetch the points for a receipt using the UUID generated in the previous step.
func FetchReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id") // Retrieve the id from the URL
	if id == "" {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	points, ok := uuidToPoint[id]
	if !ok { // If the receipt is not found, return a 404
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}
	// Return the points to the client
	json.NewEncoder(w).Encode(map[string]int64{"points": points})
}
