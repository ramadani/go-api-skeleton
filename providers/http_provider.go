package providers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/middleware"
	"github.com/ramadani/go-api-skeleton/routes"
)

// HTTPProvider contains the library of framework.
type HTTPProvider struct {
	fw  *echo.Echo
	cog *config.Config
	md  *middleware.Middleware
}

// Boot the http.
func (p *HTTPProvider) Boot() {
	if p.cog.Config.GetBool("debug") {
		p.fw.Use(p.md.Logger())
	}
	p.fw.Use(p.md.Recover())

	p.fw.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome Great Developer")
	})

	routes.APIRoutes(p.fw, p.md)
}

// InitHTTP returns route.
func InitHTTP(fw *echo.Echo, cog *config.Config, md *middleware.Middleware) *HTTPProvider {
	return &HTTPProvider{fw, cog, md}
}
