package route

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/http/handlers"
)

func ApiRoutes(fw *echo.Echo) {
	welcome := handlers.NewWelcomeHandler()

	fw.GET("/", welcome.Index)
}
