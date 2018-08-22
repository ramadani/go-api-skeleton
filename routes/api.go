package routes

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/http/handlers"
)

func NewApiRoutes(e *echo.Echo) {
	welcome := handlers.NewWelcomeHandler()

	e.GET("/", welcome.Index)
}
