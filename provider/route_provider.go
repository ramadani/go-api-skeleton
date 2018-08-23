package provider

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/route"
)

type Route struct {
	fw *echo.Echo
}

func (r Route) Boot() {
	route.ApiRoutes(r.fw)
}

func InitRoute(fw *echo.Echo) *Route {
	return &Route{fw}
}
