package route

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

type welcome struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	welcome := welcome{
		"Go API Skeleton",
		"Go (Golang) API Skeleton for your great API",
		"v0.1.0",
	}
	res, _ := json.Marshal(welcome)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// NewHandler welcome handler
func NewHandler() *Handler {
	return &Handler{}
}
