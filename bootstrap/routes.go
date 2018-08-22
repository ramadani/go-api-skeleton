package bootstrap

import "github.com/ramadani/go-api-skeleton/routes"

type RouteBoot struct {
	app App
}

func (b RouteBoot) Boot() {
	routes.NewApiRoutes(b.app.e)
}

func NewRouteBoot(app App) *RouteBoot {
	return &RouteBoot{app}
}
