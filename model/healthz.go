package model

import (
	"encoding/json"
	"log"
	"net/http"
)

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	Message string `json:"message"`
}

type HealthzHandler struct{}

func (h HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := HealthzResponse{
		Message: "OK",
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		log.Println("Failed to encode JSON:", err)
	}
}

