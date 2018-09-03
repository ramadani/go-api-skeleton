package restapi

import (
	"net/http"

	"github.com/ramadani/go-api-skeleton/app/commons/response"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/auth"
)

type AuthHandler struct {
	uc auth.UseCase
}

func (h *AuthHandler) Attempt(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	result, err := h.uc.Attempt(email, password)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.Message{err.Error()})
	}

	return c.String(http.StatusOK, result)
}

func (h *AuthHandler) Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, err := h.uc.Register(name, email, password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Message{err.Error()})
	}

	return c.JSON(http.StatusOK, response.Message{"registered. account needs to activation!"})
}

func NewHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc}
}
