package route

import (
	"net/http"

	"github.com/ramadani/go-api-skeleton/helpers/handler"
)

// Handler contains deps
type Handler struct {
	handler.Handler
}

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Index of user handlers
func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := user{"Ramadani", "email.ramadani@gmail.com"}

		h.Respose.JSON(w, user, http.StatusOK)
	}
}

// NewHandler user
func NewHandler() *Handler {
	return &Handler{}
}
