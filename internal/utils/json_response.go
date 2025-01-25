package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONErrorResponse represents the error response structure.
// @Description Error response with a message
type JSONErrorResponse struct {
	Error string `json:"error"` // Error message
}

// EmptyResponse is a placeholder for empty JSON responses.
type EmptyResponse struct{}

// RespondWithError sends an error response with the provided status code and message.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", message)
	}
	RespondWithJSON(w, code, JSONErrorResponse{Error: message})
}

// RespondWithJSON sends a JSON response with the provided status code and payload.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
