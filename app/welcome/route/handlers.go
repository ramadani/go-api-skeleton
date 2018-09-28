package route

import (
	"net/http"

	"github.com/ramadani/go-api-skeleton/commons/http/res"
)

// Handler of welcome routes
type Handler struct{}

type welcome struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

// Index handler
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	res := res.NewResponse(w)
	welcome := welcome{
		"Go API Skeleton",
		"Go (Golang) API Skeleton for your great API",
		"v0.1.0",
	}

	res.JSON(welcome, http.StatusOK)
}

// NewHandler welcome handler
func NewHandler() *Handler {
	return &Handler{}
}
