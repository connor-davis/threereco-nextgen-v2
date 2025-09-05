package organizations

import (
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/middleware"
	"github.com/connor-davis/threereco-nextgen/internal/routing"
	"github.com/connor-davis/threereco-nextgen/internal/services"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type OrganizationsRouter struct {
	Storage    storage.Storage
	Sessions   session.Store
	Services   services.Services
	Middleware middleware.Middleware
}

func NewOrganizationsRouter(
	storage storage.Storage,
	sessions session.Store,
	services services.Services,
	middleware middleware.Middleware,
) OrganizationsRouter {
	return OrganizationsRouter{
		Storage:    storage,
		Sessions:   sessions,
		Services:   services,
		Middleware: middleware,
	}
}

func (r *OrganizationsRouter) InitializeRoutes() []routing.Route {
	listRoute := r.ListRoute()
	findRoute := r.FindRoute()
	createRoute := r.CreateRoute()
	updateRoute := r.UpdateRoute()
	deleteRoute := r.DeleteRoute()

	return []routing.Route{
		listRoute,
		findRoute,
		createRoute,
		updateRoute,
		deleteRoute,
	}
}
