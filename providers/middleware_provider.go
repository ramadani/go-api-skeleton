package providers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ramadani/go-api-skeleton/config"
)

// Middleware contains the libraries.
type Middleware struct {
	fw  *echo.Echo
	cog *config.Config
}

// Boot the middlewares when starting the app.
func (md Middleware) Boot() {
	if md.cog.Config.GetBool("debug") {
		md.fw.Use(middleware.Logger())
	}

	md.fw.Use(middleware.Recover())
}

// InitMiddleware returns middleware.
func InitMiddleware(fw *echo.Echo, cog *config.Config) *Middleware {
	return &Middleware{fw, cog}
}
