package restapi

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/todo/repository"
	"github.com/ramadani/go-api-skeleton/app/todo/usecase"
)

// TodoRest contains the dependencies of todo rest api
type TodoRest struct {
	e *echo.Echo
}

// Boot the todo rest api
func (tr *TodoRest) Boot() {
	repo := repository.NewDummyRepo()
	useCase := usecase.NewUseCase(repo)
	handler := NewHandler(useCase)

	tr.e.GET("/todo", handler.Index)
	tr.e.POST("/todo", handler.Create)
	tr.e.GET("/todo/:id", handler.Find)
	tr.e.PUT("/todo/:id", handler.Update)
	tr.e.DELETE("/todo/:id", handler.Delete)
}

// New returns todo rest
func New(e *echo.Echo) *TodoRest {
	return &TodoRest{e}
}
