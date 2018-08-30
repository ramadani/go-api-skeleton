package routes

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// Route represent http route instance
type Route struct {
	e  *echo.Echo
	md *middleware.Middleware
}

// New returns Route
func New(e *echo.Echo, md *middleware.Middleware) *Route {
	return &Route{e, md}
}
