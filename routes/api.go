package routes

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/http/handlers"
)

// APIRoutes is method to register the routes and thier handlers.
func APIRoutes(fw *echo.Echo) {
	welcome := handlers.NewWelcomeHandler()

	fw.GET("/", welcome.Index)
}
