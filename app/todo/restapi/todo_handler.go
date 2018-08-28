package restapi

import (
	"net/http"

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
	return c.JSON(http.StatusOK, h.uc.Create())
}

func NewTodoHandler(uc todo.UseCase) *TodoHandler {
	return &TodoHandler{uc}
}
