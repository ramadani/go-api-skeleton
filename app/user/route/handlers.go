package route

import (
	"net/http"
	"strconv"

	"github.com/ramadani/go-api-skeleton/app/user/usecase"
	"github.com/ramadani/go-api-skeleton/commons/http/res"
)

// Handler contains deps
type Handler struct {
	ucase usecase.Usecase
}

// Index of user handlers
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	res := res.NewResponse(w)
	page := 1
	limit := 10

	if qPage := r.URL.Query().Get("page"); qPage != "" {
		page, _ = strconv.Atoi(qPage)
	}

	data, err := h.ucase.Paginate(uint(page), uint(limit))
	if err != nil {
		res.Fail(err.Error(), http.StatusInternalServerError)
		return
	}

	res.JSON(data, http.StatusOK)
}

// Store a new user
func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	res := res.NewResponse(w)
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.ucase.Create(name, email, password)
	if err != nil {
		res.Fail(err.Error(), http.StatusInternalServerError)
		return
	}

	res.JSON(user, http.StatusOK)
}

// Find user by given id
func (h *Handler) Find(w http.ResponseWriter, r *http.Request, id uint) {
	res := res.NewResponse(w)
	user, err := h.ucase.FindByID(id)
	if err != nil {
		res.Fail(err.Error(), http.StatusNotFound)
		return
	}

	res.JSON(user, http.StatusOK)
}

// NewHandler user
func NewHandler(ucase usecase.Usecase) *Handler {
	return &Handler{ucase: ucase}
}
