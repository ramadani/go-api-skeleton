package routes

import (
	todoRestAPI "github.com/ramadani/go-api-skeleton/app/todo/restapi"
)

// API routes
func (r *Route) API() {
	todoRestAPI.TodoRoutes(r.fw, r.md)
}
