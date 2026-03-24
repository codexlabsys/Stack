package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type StatusResponse struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Mode        string    `json:"mode"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := StatusResponse{
		Name:        "Stack",
		Status:      "coming_soon",
		Mode:        "autonomous",
		Description: "A self-funded AI bot that trades and reinvests profits to grow its own money.",
		Version:     "0.1.0",
		UpdatedAt:   time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Stack server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
