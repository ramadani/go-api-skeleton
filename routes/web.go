package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// Web routes
func (r *Route) Web() {
	r.fw.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome Great Developer")
	})
}
