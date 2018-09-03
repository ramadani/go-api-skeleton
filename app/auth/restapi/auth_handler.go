package restapi

import (
	"net/http"

	"github.com/ramadani/go-api-skeleton/app/auth/dto"
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
	registrar := new(dto.RegistrarDto)

	if err := c.Bind(registrar); err != nil {
		return c.JSON(http.StatusBadRequest, response.Message{"Bad Request"})
	}

	_, err := h.uc.Register(registrar)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Message{err.Error()})
	}

	return c.JSON(http.StatusOK, response.Message{"registered. account needs to activation!"})
}

func NewHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{uc}
}
