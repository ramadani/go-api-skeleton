package handler

import (
	"encoding/json"
	"net/http"
)

// Handler helpers
type Handler struct{}

// JSON response
func (h *Handler) JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}
