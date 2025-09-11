package permissions

import (
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/middleware"
	"github.com/connor-davis/threereco-nextgen/internal/routing"
	"github.com/connor-davis/threereco-nextgen/internal/services"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type PermissionsRouter struct {
	Storage    storage.Storage
	Sessions   session.Store
	Services   services.Services
	Middleware middleware.Middleware
}

func NewPermissionsRouter(
	storage storage.Storage,
	sessions session.Store,
	services services.Services,
	middleware middleware.Middleware,
) PermissionsRouter {
	return PermissionsRouter{
		Storage:    storage,
		Sessions:   sessions,
		Services:   services,
		Middleware: middleware,
	}
}

func (r *PermissionsRouter) InitializeRoutes() []routing.Route {
	listRoute := r.ListRoute()

	return []routing.Route{
		listRoute,
	}
}
