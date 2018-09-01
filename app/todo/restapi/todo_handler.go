package restapi

import (
	"net/http"
	"strconv"

	"github.com/ramadani/go-api-skeleton/app/commons/response"

	"github.com/ramadani/go-api-skeleton/app/todo/resource"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/todo"
)

// TodoHandler represents the handler for todo.
type TodoHandler struct {
	uc todo.UseCase
}

// Index to show todo list
func (h *TodoHandler) Index(c echo.Context) error {
	result := resource.Collection(h.uc.All())

	return c.JSON(http.StatusOK, result)
}

// Create a new todo
func (h *TodoHandler) Create(c echo.Context) error {
	title := c.FormValue("title")
	body := c.FormValue("body")
	todos := h.uc.Create(title, body)
	result := resource.Item(todos)

	return c.JSON(http.StatusOK, result)
}

// Find a todo from collection
func (h *TodoHandler) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.uc.Find(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.Message{
			Message: err.Error(),
		})
	}

	result := resource.Item(todo)

	return c.JSON(http.StatusOK, result)
}

// Update an existing todo
func (h *TodoHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	todo, err := h.uc.Update(title, body, uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.Message{
			Message: err.Error(),
		})
	}

	result := resource.Item(todo)

	return c.JSON(http.StatusOK, result)
}

// Delete a todo from collection
func (h *TodoHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.uc.Delete(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.Message{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Message{
		Message: "Todo has been deleted",
	})
}

// NewHandler returns todo handler
func NewHandler(uc todo.UseCase) *TodoHandler {
	return &TodoHandler{uc}
}
