package providers

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/routes"
)

type Route struct {
	fw *echo.Echo
}

func (r Route) Boot() {
	routes.ApiRoutes(r.fw)
}

func InitRoute(fw *echo.Echo) *Route {
	return &Route{fw}
}
