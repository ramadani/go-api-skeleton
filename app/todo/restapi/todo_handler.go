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

func (h *TodoHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, h.uc.All())
}

func (h *TodoHandler) Create(c echo.Context) error {
	title := c.FormValue("title")
	body := c.FormValue("body")
	result := h.uc.Create(title, body)
	return c.JSON(http.StatusOK, result)
}

func (h *TodoHandler) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := h.uc.Find(uint(id))
	return c.JSON(http.StatusOK, result)
}

func (h *TodoHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	body := c.FormValue("body")
	result := h.uc.Update(title, body, uint(id))
	return c.JSON(http.StatusOK, result)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := h.uc.Delete(uint(id))
	return c.JSON(http.StatusOK, result)
}

func NewTodoHandler(uc todo.UseCase) *TodoHandler {
	return &TodoHandler{uc}
}
