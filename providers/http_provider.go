package providers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// HTTPProvider contains the library of framework.
type HTTPProvider struct {
	e   *echo.Echo
	cog *config.Config
	md  *middleware.Middleware
}

// Boot the http.
func (pd *HTTPProvider) Boot() {
	if pd.cog.Config.GetBool("debug") {
		pd.e.Use(pd.md.Logger())
	}

	pd.e.Use(pd.md.Recover())

	pd.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome Great Developer")
	})
}

// NewHTTP returns route.
func NewHTTP(e *echo.Echo, cog *config.Config, md *middleware.Middleware) *HTTPProvider {
	return &HTTPProvider{e, cog, md}
}
