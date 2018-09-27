package route

import (
	"net/http"
	"strconv"

	"github.com/ramadani/go-api-skeleton/app/user/usecase"

	"github.com/ramadani/go-api-skeleton/commons/handler"
)

// Handler contains deps
type Handler struct {
	handler.Handler
	ucase usecase.Usecase
}

// Index of user handlers
func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := 1
		limit := 10

		if qPage := r.URL.Query().Get("page"); qPage != "" {
			page, _ = strconv.Atoi(qPage)
		}

		res, err := h.ucase.Paginate(uint(page), uint(limit))
		if err != nil {
			h.Response.Fail(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.Response.JSON(w, res, http.StatusOK)
	}
}

// Store a new user
func (h *Handler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := h.ucase.Create(name, email, password)
		if err != nil {
			h.Response.Fail(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.Response.JSON(w, user, http.StatusOK)
	}
}

// NewHandler user
func NewHandler(ucase usecase.Usecase) *Handler {
	return &Handler{ucase: ucase}
}
