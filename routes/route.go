package routes

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// Route represent http route instance
type Route struct {
	fw *echo.Echo
	md *middleware.Middleware
}

// New returns Route
func New(fw *echo.Echo, md *middleware.Middleware) *Route {
	return &Route{fw, md}
}
