package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthResponse represents the response body
// returned by the health endpoint.
type healthResponse struct {
	Status string `json:"status"`
}

// NewRouter initializes the application's HTTP routes
// and returns an http.Handler.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Go 1.22+ routing syntax
	mux.HandleFunc("GET /health", handleHealth)

	return mux
}

// handleHealth returns the service health status.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := healthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode health response: %v", err)
	}
}
