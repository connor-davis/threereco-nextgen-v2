package transactionMaterials

import (
	"github.com/connor-davis/threereco-nextgen/internal/constants"
	"github.com/connor-davis/threereco-nextgen/internal/routing"
	"github.com/connor-davis/threereco-nextgen/internal/routing/schemas"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListParams struct {
	TransactionId uuid.UUID `json:"transactionId"`
}

type ListQueryParams struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}

func (r *TransactionMaterialsRouter) ListRoute() routing.Route {
	responses := openapi3.NewResponses()

	responses.Set("200", &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Successful transaction materials retrieval.").
			WithJSONSchema(schemas.SuccessResponseSchema.Value),
	})

	responses.Set("400", &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithJSONSchema(schemas.ErrorResponseSchema.Value).
			WithDescription(string(constants.BadRequestError)).
			WithContent(openapi3.Content{
				"application/json": openapi3.NewMediaType().
					WithSchema(schemas.ErrorResponseSchema.Value).
					WithExample("example", map[string]any{
						"error":   string(constants.BadRequestError),
						"message": string(constants.BadRequestErrorDetails),
					}),
			}),
	})

	responses.Set("401", &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithJSONSchema(schemas.ErrorResponseSchema.Value).
			WithDescription(string(constants.UnauthorizedError)).
			WithContent(openapi3.Content{
				"application/json": openapi3.NewMediaType().
					WithSchema(schemas.ErrorResponseSchema.Value).
					WithExample("example", map[string]any{
						"error":   string(constants.UnauthorizedError),
						"message": string(constants.UnauthorizedErrorDetails),
					}),
			}),
	})

	responses.Set("500", &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithJSONSchema(schemas.ErrorResponseSchema.Value).
			WithDescription("Internal server error.").
			WithContent(openapi3.Content{
				"application/json": openapi3.NewMediaType().
					WithSchema(schemas.ErrorResponseSchema.Value).
					WithExample("example", map[string]any{
						"error":   string(constants.InternalServerError),
						"message": string(constants.InternalServerErrorDetails),
					}),
			}),
	})

	paramters := []*openapi3.ParameterRef{
		{
			Value: openapi3.NewPathParameter("transactionId").
				WithRequired(true).
				WithSchema(openapi3.NewUUIDSchema()),
		},
		{
			Value: openapi3.NewQueryParameter("page").
				WithRequired(true).
				WithSchema(openapi3.NewInt64Schema().
					WithDefault(1).WithMin(1)).
				WithDescription("Page number for pagination. Defaults to 1."),
		},
		{
			Value: openapi3.NewQueryParameter("limit").
				WithRequired(true).
				WithSchema(openapi3.NewInt64Schema().
					WithDefault(10).WithMin(10)).
				WithDescription("Number of items per page. Defaults to 10."),
		},
		{
			Value: openapi3.NewQueryParameter("search").
				WithRequired(true).
				WithSchema(openapi3.NewStringSchema()).
				WithDescription("Search term for filtering transactions."),
		},
	}

	return routing.Route{
		OpenAPIMetadata: routing.OpenAPIMetadata{
			Summary:     "List Transaction Materials",
			Description: "List all transaction materials in the system.",
			Tags:        []string{"Transaction Materials"},
			Responses:   responses,
			Parameters:  paramters,
			RequestBody: nil,
		},
		Method: routing.GetMethod,
		Path:   "/transactions/:transactionId/materials",
		Middlewares: []fiber.Handler{
			r.Middleware.Authenticated(),
			r.Middleware.Authorized([]string{"transactions.materials.view"}),
		},
		Handler: func(c *fiber.Ctx) error {
			var params ListParams

			if err := c.ParamsParser(&params); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   constants.BadRequestError,
					"message": constants.BadRequestErrorDetails,
				})
			}

			var query ListQueryParams

			if err := c.QueryParser(&query); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   constants.BadRequestError,
					"message": constants.BadRequestErrorDetails,
				})
			}

			totalTransactionMaterials, err := r.Services.Transactions().Materials().Count(
				clause.Eq{
					Column: clause.Column{
						Name: "transaction_id",
					},
					Value: params.TransactionId,
				},
			)

			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
						"error":   constants.NotFoundError,
						"message": constants.NotFoundErrorDetails,
					})
				}

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   constants.InternalServerError,
					"message": constants.InternalServerErrorDetails,
				})
			}

			paginationClauses := []clause.Expression{
				clause.Eq{
					Column: clause.Column{
						Name: "transaction_id",
					},
					Value: params.TransactionId,
				},
				clause.Limit{
					Limit:  &query.Limit,
					Offset: (query.Page - 1) * query.Limit,
				},
			}

			totalPages := (totalTransactionMaterials + int64(query.Limit) - 1) / int64(query.Limit)

			transactionMaterials, err := r.Services.Transactions().Materials().List(paginationClauses...)

			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
						"error":   constants.NotFoundError,
						"message": constants.NotFoundErrorDetails,
					})
				}

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   constants.InternalServerError,
					"message": constants.InternalServerErrorDetails,
				})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"items": transactionMaterials,
				"pageDetails": map[string]any{
					"count":        totalTransactionMaterials,
					"nextPage":     query.Page + 1,
					"previousPage": query.Page - 1,
					"currentPage":  query.Page,
					"pages":        totalPages,
				},
			})
		},
	}
}
