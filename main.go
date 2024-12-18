package main

import (
	"net/http"
)

var uuidToPoint = make(map[string]int64)

func main() {
	mux := http.NewServeMux()

	// Register the handlers
	mux.Handle("/receipts/process", http.HandlerFunc(CalculateReceiptPointsHandler))
	mux.Handle("/receipts/{id}/points", http.HandlerFunc(FetchReceiptPointsHandler))

	// Start the server
	http.ListenAndServe(":8080", mux)
}
