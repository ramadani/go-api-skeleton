package routes

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/todo/repository"
	"github.com/ramadani/go-api-skeleton/app/todo/usecase"
	"github.com/ramadani/go-api-skeleton/http/handlers"
)

// APIRoutes is method to register the routes and thier handlers.
func APIRoutes(fw *echo.Echo) {
	welcomeHandler := handlers.NewWelcomeHandler()

	todoRepoDummy := repository.NewTodoRepositoryDummy()
	todoUseCase := usecase.NewTodoUseCase(todoRepoDummy)
	todoHandler := handlers.NewTodoHandler(todoUseCase)

	fw.GET("/", welcomeHandler.Index)
	fw.GET("/todo", todoHandler.Index)
}
