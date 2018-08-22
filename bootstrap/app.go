package bootstrap

import (
	"fmt"

	"github.com/ramadani/go-api-skeleton/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	e *echo.Echo
}

func (app App) InitMiddleware() {
	app.e.Use(middleware.Logger())
	app.e.Use(middleware.Recover())
}

func (app App) InitRoutes() {
	routes.NewApiRoutes(app.e)
}

func (app App) Run(port int) {
	app.e.Logger.Fatal(app.e.Start(fmt.Sprintf(":%d", port)))
}

func New(e *echo.Echo) *App {
	return &App{e}
}
