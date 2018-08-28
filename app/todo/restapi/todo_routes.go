package restapi

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/todo/repository"
	"github.com/ramadani/go-api-skeleton/app/todo/usecase"
	"github.com/ramadani/go-api-skeleton/middleware"
)

// TodoRoutes contains todo routes
func TodoRoutes(fw *echo.Echo, md *middleware.Middleware) {
	todoRepo := repository.NewDummyTodoRepo()
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	todoHandler := NewTodoHandler(todoUseCase)

	fw.GET("/todo", todoHandler.Index)
	fw.POST("/todo", todoHandler.Create)
}
