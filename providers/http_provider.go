package providers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// HTTPProvider contains the library of framework.
type HTTPProvider struct {
	e  *echo.Echo
	md *middleware.Middleware
}

// Boot the http.
func (pd *HTTPProvider) Boot() {
	pd.e.Use(pd.md.Logger())
	pd.e.Use(pd.md.Recover())

	pd.e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"name":        "go-api-skeleton",
			"creator":     "Ramadani <email.ramadani@gmail.com>",
			"description": "Go (Golang) API Skeleton",
			"version":     "0.1.0-dev",
		})
	})
}

// NewHTTP returns route.
func NewHTTP(e *echo.Echo, md *middleware.Middleware) *HTTPProvider {
	return &HTTPProvider{e, md}
}
