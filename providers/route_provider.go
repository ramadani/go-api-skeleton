package providers

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/routes"
)

// Route contains the library of framework.
type Route struct {
	fw *echo.Echo
}

// Boot the route.
func (r Route) Boot() {
	routes.APIRoutes(r.fw)
}

// InitRoute returns route.
func InitRoute(fw *echo.Echo) *Route {
	return &Route{fw}
}
