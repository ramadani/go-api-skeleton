package restapi

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/todo/repository"
	"github.com/ramadani/go-api-skeleton/app/todo/usecase"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// TodoRoutes contains todo routes
func TodoRoutes(e *echo.Echo, md *middleware.Middleware) {
	todoRepo := repository.NewDummyTodoRepo()
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	todoHandler := NewTodoHandler(todoUseCase)

	e.GET("/todo", todoHandler.Index)
	e.POST("/todo", todoHandler.Create)
	e.GET("/todo/:id", todoHandler.Find)
	e.PUT("/todo/:id", todoHandler.Update)
	e.DELETE("/todo/:id", todoHandler.Delete)
}
