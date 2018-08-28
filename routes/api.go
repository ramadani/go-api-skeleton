package routes

import (
	"github.com/labstack/echo"
	todo "github.com/ramadani/go-api-skeleton/app/todo/restapi"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// APIRoutes is method to register the routes and thier handlers.
func APIRoutes(fw *echo.Echo, md *middleware.Middleware) {
	todo.TodoRoutes(fw, md)
}
