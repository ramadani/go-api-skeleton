package providers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/middleware"
	"github.com/ramadani/go-api-skeleton/routes"
)

// HTTP contains the library of framework.
type HTTP struct {
	fw  *echo.Echo
	cog *config.Config
	md  *middleware.Middleware
}

// Boot the http.
func (p *HTTP) Boot() {
	if p.cog.Config.GetBool("debug") {
		p.fw.Use(p.md.Logger())
	}
	p.fw.Use(p.md.Recover())

	p.fw.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome Great Developer")
	})

	routes.APIRoutes(p.fw)
}

// InitHTTP returns route.
func InitHTTP(fw *echo.Echo, cog *config.Config, md *middleware.Middleware) *HTTP {
	return &HTTP{fw, cog, md}
}
