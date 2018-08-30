package providers

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/middleware"
	"github.com/ramadani/go-api-skeleton/routes"
)

// HTTPProvider contains the library of framework.
type HTTPProvider struct {
	e   *echo.Echo
	cog *config.Config
	md  *middleware.Middleware
}

// Boot the http.
func (p *HTTPProvider) Boot() {
	if p.cog.Config.GetBool("debug") {
		p.e.Use(p.md.Logger())
	}
	p.e.Use(p.md.Recover())

	routes := routes.New(p.e, p.md)
	routes.Web()
	routes.API()
}

// NewHTTP returns route.
func NewHTTP(e *echo.Echo, cog *config.Config, md *middleware.Middleware) *HTTPProvider {
	return &HTTPProvider{e, cog, md}
}
