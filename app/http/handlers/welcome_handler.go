package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type WelcomeHandler struct{}

func (h WelcomeHandler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome Great Developer")
}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}
