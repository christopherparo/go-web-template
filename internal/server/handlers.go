package server

import (
	"encoding/json"
	"net/http"
)

// healthResponse is the JSON payload returned by the health endpoint.
type healthResponse struct {
	Status string `json:"status"`
}

// handleHealth returns the service health status.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}
