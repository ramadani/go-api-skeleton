package providers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// Middleware contains the libraries.
type Middleware struct {
	fw  *echo.Echo
	cog *viper.Viper
}

// Boot the middlewares when starting the app.
func (md Middleware) Boot() {
	if md.cog.GetBool("debug") {
		md.fw.Use(middleware.Logger())
	}

	md.fw.Use(middleware.Recover())
}

// InitMiddleware returns middleware.
func InitMiddleware(fw *echo.Echo, cog *viper.Viper) *Middleware {
	return &Middleware{fw, cog}
}
