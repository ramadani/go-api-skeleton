package route

import (
	"net/http"

	"github.com/ramadani/go-api-skeleton/app/user/usecase"

	"github.com/ramadani/go-api-skeleton/helpers/handler"
)

// Handler contains deps
type Handler struct {
	handler.Handler
	ucase usecase.Usecase
}

// Index of user handlers
func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.ucase.Paginate(uint(1), uint(10))
		if err != nil {
			h.Response.Fail(w, err.Error(), http.StatusInternalServerError)
		}

		h.Response.JSON(w, res, http.StatusOK)
	}
}

// NewHandler user
func NewHandler(ucase usecase.Usecase) *Handler {
	return &Handler{ucase: ucase}
}
