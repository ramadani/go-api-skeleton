package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// WelcomeHandler represents the handler for welcome.
type WelcomeHandler struct{}

// Index return the welcome index's response.
func (h *WelcomeHandler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome Great Developer")
}

// NewWelcomeHandler returns welcome handler.
func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}
