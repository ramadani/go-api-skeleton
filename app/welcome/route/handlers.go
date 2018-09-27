package route

import (
	"net/http"

	"github.com/ramadani/go-api-skeleton/commons/handler"
)

// Handler of welcome routes
type Handler struct {
	handler.Handler
}

type welcome struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

// Index handler
func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		welcome := welcome{
			"Go API Skeleton",
			"Go (Golang) API Skeleton for your great API",
			"v0.1.0",
		}

		h.Response.JSON(w, welcome, http.StatusOK)
	}
}

// NewHandler welcome handler
func NewHandler() *Handler {
	return &Handler{}
}
