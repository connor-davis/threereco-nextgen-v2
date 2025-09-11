package http

import (
	"fmt"
	"regexp"

	"github.com/connor-davis/threereco-nextgen/cmd/api/http/addresses"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/authentication"
	bankDetails "github.com/connor-davis/threereco-nextgen/cmd/api/http/bank-details"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/collections"
	collectionMaterials "github.com/connor-davis/threereco-nextgen/cmd/api/http/collections/materials"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/materials"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/middleware"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/organizations"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/permissions"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/roles"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/transactions"
	transactionMaterials "github.com/connor-davis/threereco-nextgen/cmd/api/http/transactions/materials"
	"github.com/connor-davis/threereco-nextgen/cmd/api/http/users"
	"github.com/connor-davis/threereco-nextgen/env"
	"github.com/connor-davis/threereco-nextgen/internal/routing"
	"github.com/connor-davis/threereco-nextgen/internal/routing/schemas"
	"github.com/connor-davis/threereco-nextgen/internal/services"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// HttpRouter encapsulates the dependencies and configuration required to set up HTTP routing.
// It includes references to storage, session management, service layer, middleware, and route definitions.
type HttpRouter struct {
	Storage    storage.Storage
	Sessions   session.Store
	Services   services.Services
	Middleware middleware.Middleware
	Routes     []routing.Route
}

// NewHttpRouter creates and initializes a new HttpRouter instance.
// It sets up middleware and authentication routes using the provided storage,
// session store, and services. The returned HttpRouter contains all configured
// routes and dependencies required for handling HTTP requests.
//
// Parameters:
//   - storage: Pointer to the application's storage layer.
//   - sessions: Pointer to the session store for managing user sessions.
//   - services: Pointer to the application's service layer.
//
// Returns:
//   - HttpRouter: A pointer to the initialized HttpRouter.
func NewHttpRouter(storage storage.Storage, sessions session.Store, services services.Services, middleware middleware.Middleware) HttpRouter {
	authenticationRouter := authentication.NewAuthenticationRouter(storage, sessions, services, middleware)
	authenticationRoutes := authenticationRouter.InitializeRoutes()

	materialsRouter := materials.NewMaterialsRouter(storage, sessions, services, middleware)
	materialsRoutes := materialsRouter.InitializeRoutes()

	usersRouter := users.NewUsersRouter(storage, sessions, services, middleware)
	usersRoutes := usersRouter.InitializeRoutes()

	rolesRouter := roles.NewRolesRouter(storage, sessions, services, middleware)
	rolesRoutes := rolesRouter.InitializeRoutes()

	organizationsRouter := organizations.NewOrganizationsRouter(storage, sessions, services, middleware)
	organizationsRoutes := organizationsRouter.InitializeRoutes()

	collectionsRouter := collections.NewCollectionsRouter(storage, sessions, services, middleware)
	collectionsRoutes := collectionsRouter.InitializeRoutes()

	collectionMaterialsRouter := collectionMaterials.NewCollectionMaterialsRouter(storage, sessions, services, middleware)
	collectionMaterialsRoutes := collectionMaterialsRouter.InitializeRoutes()

	transactionsRouter := transactions.NewTransactionsRouter(storage, sessions, services, middleware)
	transactionsRoutes := transactionsRouter.InitializeRoutes()

	transactionMaterialsRouter := transactionMaterials.NewTransactionsRouter(storage, sessions, services, middleware)
	transactionMaterialsRoutes := transactionMaterialsRouter.InitializeRoutes()

	addressesRouter := addresses.NewAddressesRouter(storage, sessions, services, middleware)
	addressesRoutes := addressesRouter.InitializeRoutes()

	bankDetailsRouter := bankDetails.NewBankDetailsRouter(storage, sessions, services, middleware)
	bankDetailsRoutes := bankDetailsRouter.InitializeRoutes()

	permissionsRouter := permissions.NewPermissionsRouter(storage, sessions, services, middleware)
	permissionsRoutes := permissionsRouter.InitializeRoutes()

	routes := []routing.Route{}

	routes = append(routes, authenticationRoutes...)
	routes = append(routes, materialsRoutes...)
	routes = append(routes, usersRoutes...)
	routes = append(routes, rolesRoutes...)
	routes = append(routes, organizationsRoutes...)
	routes = append(routes, collectionsRoutes...)
	routes = append(routes, collectionMaterialsRoutes...)
	routes = append(routes, transactionsRoutes...)
	routes = append(routes, transactionMaterialsRoutes...)
	routes = append(routes, addressesRoutes...)
	routes = append(routes, bankDetailsRoutes...)
	routes = append(routes, permissionsRoutes...)

	return HttpRouter{
		Storage:    storage,
		Sessions:   sessions,
		Services:   services,
		Middleware: middleware,
		Routes:     routes,
	}
}

// InitializeRoutes registers all HTTP routes defined in the HttpRouter with the provided Fiber router.
// It converts route paths from the format "{param}" to Fiber's ":param" syntax using regular expressions.
// For each route, it attaches the specified middlewares and handler to the corresponding HTTP method.
// Supported methods include GET, POST, PUT, PATCH, OPTIONS, and DELETE.
func (r *HttpRouter) InitializeRoutes(router fiber.Router) {
	for _, route := range r.Routes {
		path := regexp.MustCompile(`\{([^}]+)\}`).ReplaceAllString(route.Path, ":$1")

		switch route.Method {
		case routing.GetMethod:
			router.Get(path, append(route.Middlewares, route.Handler)...)
		case routing.PostMethod:
			router.Post(path, append(route.Middlewares, route.Handler)...)
		case routing.PutMethod:
			router.Put(path, append(route.Middlewares, route.Handler)...)
		case routing.PatchMethod:
			router.Patch(path, append(route.Middlewares, route.Handler)...)
		case routing.OptionsMethod:
			router.Options(path, append(route.Middlewares, route.Handler)...)
		case routing.DeleteMethod:
			router.Delete(path, append(route.Middlewares, route.Handler)...)
		}
	}
}

// InitializeOpenAPI generates and returns an OpenAPI 3 specification for the HTTP router.
// It iterates through the defined routes, constructs OpenAPI PathItem objects for each HTTP method,
// and sets up the API paths, operations, and components (schemas, servers, etc.).
// The resulting OpenAPI specification includes metadata such as title, version, server URLs, and
// reusable schemas for success and error responses.
//
// Returns:
//
//	*openapi3.T: The constructed OpenAPI 3 specification for the API.
func (h *HttpRouter) InitializeOpenAPI() *openapi3.T {
	paths := openapi3.NewPaths()

	for _, route := range h.Routes {
		pathItem := &openapi3.PathItem{}

		switch route.Method {
		case routing.GetMethod:
			pathItem.Get = &openapi3.Operation{
				Summary:     route.Summary,
				Description: route.Description,
				Tags:        route.Tags,
				Parameters:  route.Parameters,
				Responses:   route.Responses,
			}
		case routing.PostMethod:
			pathItem.Post = &openapi3.Operation{
				Summary:     route.Summary,
				Description: route.Description,
				Tags:        route.Tags,
				Parameters:  route.Parameters,
				RequestBody: route.RequestBody,
				Responses:   route.Responses,
			}
		case routing.PutMethod:
			pathItem.Put = &openapi3.Operation{
				Summary:     route.Summary,
				Description: route.Description,
				Tags:        route.Tags,
				Parameters:  route.Parameters,
				RequestBody: route.RequestBody,
				Responses:   route.Responses,
			}
		case routing.PatchMethod:
			pathItem.Patch = &openapi3.Operation{
				Summary:     route.Summary,
				Description: route.Description,
				Tags:        route.Tags,
				Parameters:  route.Parameters,
				RequestBody: route.RequestBody,
				Responses:   route.Responses,
			}
		case routing.OptionsMethod:
			pathItem.Options = &openapi3.Operation{
				Summary:     route.Summary,
				Description: route.Description,
				Tags:        route.Tags,
				Parameters:  route.Parameters,
				RequestBody: route.RequestBody,
				Responses:   route.Responses,
			}
		case routing.DeleteMethod:
			pathItem.Delete = &openapi3.Operation{
				Summary:     route.Summary,
				Description: route.Description,
				Tags:        route.Tags,
				Parameters:  route.Parameters,
				Responses:   route.Responses,
			}
		}

		path := fmt.Sprintf("/api%s", route.Path)

		existingPathItem := paths.Find(path)

		if existingPathItem != nil {
			switch route.Method {
			case routing.GetMethod:
				existingPathItem.Get = pathItem.Get
			case routing.PostMethod:
				existingPathItem.Post = pathItem.Post
			case routing.PutMethod:
				existingPathItem.Put = pathItem.Put
			case routing.PatchMethod:
				existingPathItem.Patch = pathItem.Patch
			case routing.DeleteMethod:
				existingPathItem.Delete = pathItem.Delete
			}
		} else {
			paths.Set(path, pathItem)
		}
	}

	return &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   "3rEco API",
			Version: "1.0.0",
		},
		Servers: openapi3.Servers{
			{
				URL:         fmt.Sprintf("http://localhost:%s", string(env.PORT)),
				Description: "Development",
			},
			{
				URL:         "https://3reco.co.za",
				Description: "Production",
			},
		},
		Tags:  openapi3.Tags{},
		Paths: paths,
		Components: &openapi3.Components{
			Schemas: openapi3.Schemas{
				"SuccessResponse":           schemas.SuccessResponseSchema,
				"ErrorResponse":             schemas.ErrorResponseSchema,
				"AvailablePermissions":      schemas.AvailablePermissionsSchema,
				"MfaVerifyPayload":          schemas.MfaVerifyPayloadSchema,
				"LoginPayload":              schemas.LoginPayloadSchema,
				"SignUpPayload":             schemas.SignUpPayloadSchema,
				"User":                      schemas.UserSchema,
				"Users":                     schemas.UsersSchema,
				"Role":                      schemas.RoleSchema,
				"Roles":                     schemas.RolesSchema,
				"Organization":              schemas.OrganizationSchema,
				"Organizations":             schemas.OrganizationsSchema,
				"Material":                  schemas.MaterialSchema,
				"Materials":                 schemas.MaterialsSchema,
				"Address":                   schemas.AddressSchema,
				"Addresses":                 schemas.AddressesSchema,
				"BankDetail":                schemas.BankDetailSchema,
				"BankDetails":               schemas.BankDetailsSchema,
				"Collection":                schemas.CollectionSchema,
				"Collections":               schemas.CollectionsSchema,
				"Transaction":               schemas.TransactionSchema,
				"Transactions":              schemas.TransactionsSchema,
				"CreateUser":                schemas.CreateUserSchema,
				"UpdateUser":                schemas.UpdateUserSchema,
				"CreateRole":                schemas.CreateRoleSchema,
				"UpdateRole":                schemas.UpdateRoleSchema,
				"CreateOrganization":        schemas.CreateOrganizationSchema,
				"UpdateOrganization":        schemas.UpdateOrganizationSchema,
				"CreateMaterial":            schemas.CreateMaterialSchema,
				"UpdateMaterial":            schemas.UpdateMaterialSchema,
				"CreateCollection":          schemas.CreateCollectionSchema,
				"UpdateCollection":          schemas.UpdateCollectionSchema,
				"CreateCollectionMaterial":  schemas.CreateCollectionMaterialSchema,
				"UpdateCollectionMaterial":  schemas.UpdateCollectionMaterialSchema,
				"CreateTransaction":         schemas.CreateTransactionSchema,
				"UpdateTransaction":         schemas.UpdateTransactionSchema,
				"CreateTransactionMaterial": schemas.CreateTransactionMaterialSchema,
				"UpdateTransactionMaterial": schemas.UpdateTransactionMaterialSchema,
				"CreateAddress":             schemas.CreateAddressSchema,
				"UpdateAddress":             schemas.UpdateAddressSchema,
				"CreateBankDetail":          schemas.CreateBankDetailSchema,
				"UpdateBankDetail":          schemas.UpdateBankDetailSchema,
			},
		},
	}
}
