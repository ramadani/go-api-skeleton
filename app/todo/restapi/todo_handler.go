package restapi

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/todo"
)

// TodoHandler represents the handler for todo.
type TodoHandler struct {
	uc todo.UseCase
}

// Index to show todo list
func (h *TodoHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, h.uc.All())
}

// Create a new todo
func (h *TodoHandler) Create(c echo.Context) error {
	title := c.FormValue("title")
	body := c.FormValue("body")
	result := h.uc.Create(title, body)
	return c.JSON(http.StatusOK, result)
}

// Find a todo from collection
func (h *TodoHandler) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := h.uc.Find(uint(id))
	return c.JSON(http.StatusOK, result)
}

// Update an existing todo
func (h *TodoHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	result := h.uc.Update(title, body, uint(id))
	return c.JSON(http.StatusOK, result)
}

// Delete a todo from collection
func (h *TodoHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := h.uc.Delete(uint(id))
	return c.JSON(http.StatusOK, result)
}

// NewHandler returns todo handler
func NewHandler(uc todo.UseCase) *TodoHandler {
	return &TodoHandler{uc}
}
