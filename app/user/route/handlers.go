package route

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/usecase"
	"github.com/ramadani/go-api-skeleton/app/user/validators"
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
	var input data.UserInput
	res := res.NewResponse(w)
	err := r.ParseForm()
	if err != nil {
		res.Fail(err.Error(), http.StatusBadRequest)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&input, r.PostForm)
	if err != nil {
		res.Fail(err.Error(), http.StatusBadRequest)
		return
	}

	validator := validators.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 {
		res.ValidationError(errs, http.StatusBadRequest)
		return
	}

	user, err := h.ucase.Create(input.Name, input.Email, input.Password)
	if err != nil {
		res.Fail(err.Error(), http.StatusInternalServerError)
		return
	}

	res.JSON(user, http.StatusOK)
}

// Find user by given id
func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	res := res.NewResponse(w)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.ucase.FindByID(uint(id))
	if err != nil {
		res.Fail(err.Error(), http.StatusNotFound)
		return
	}

	res.JSON(user, http.StatusOK)
}

// Update an existing user
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var input data.UserUpdateInput
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	res := res.NewResponse(w)
	err := r.ParseForm()
	if err != nil {
		res.Fail(err.Error(), http.StatusBadRequest)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&input, r.PostForm)
	if err != nil {
		res.Fail(err.Error(), http.StatusBadRequest)
		return
	}

	validator := validators.NewValidator()
	if errs := validator.Validate(input); len(errs) > 0 {
		res.ValidationError(errs, http.StatusBadRequest)
		return
	}

	user, err := h.ucase.Update(input.Name, uint(id))
	if err != nil {
		res.Fail(err.Error(), http.StatusInternalServerError)
		return
	}

	res.JSON(user, http.StatusOK)
}

// NewHandler user
func NewHandler(ucase usecase.Usecase) *Handler {
	return &Handler{ucase: ucase}
}
