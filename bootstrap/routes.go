package bootstrap

import (
	"github.com/ramadani/go-api-skeleton/routes"
)

type Route struct {
	app App
}

func (b Route) Boot() {
	routes.NewApiRoutes(b.app.e)
}

func InitRoute(app App) *Route {
	return &Route{app}
}
