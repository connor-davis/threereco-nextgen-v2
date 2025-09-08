package transactionMaterials

import (
	"github.com/connor-davis/threereco-nextgen/internal/constants"
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/routing"
	"github.com/connor-davis/threereco-nextgen/internal/routing/schemas"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateParams struct {
	TransactionId uuid.UUID `json:"transactionId"`
}

func (r *TransactionMaterialsRouter) CreateRoute() routing.Route {
	responses := openapi3.NewResponses()

	responses.Set("200", &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Successful transaction material creation.").
			WithJSONSchema(schemas.SuccessResponseSchema.Value).
			WithContent(openapi3.Content{
				"text/plain": openapi3.NewMediaType().
					WithSchema(openapi3.NewStringSchema()).
					WithExample("example", "OK"),
			}),
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

	parameters := []*openapi3.ParameterRef{
		{
			Value: openapi3.NewPathParameter("transactionId").
				WithRequired(true).
				WithSchema(openapi3.NewUUIDSchema()),
		},
	}

	body := &openapi3.RequestBodyRef{
		Value: openapi3.NewRequestBody().
			WithRequired(true).
			WithDescription("Payload to create a new transaction material.").
			WithContent(openapi3.Content{
				"application/json": openapi3.NewMediaType().
					WithSchema(schemas.CreateTransactionMaterialSchema.Value).
					WithExample("example", schemas.CreateTransactionMaterialSchema.Value),
			}),
	}

	return routing.Route{
		OpenAPIMetadata: routing.OpenAPIMetadata{
			Summary:     "Create Transaction Material",
			Description: "Create a new transaction material in the system.",
			Tags:        []string{"Transaction Materials"},
			Responses:   responses,
			Parameters:  parameters,
			RequestBody: body,
		},
		Method: routing.PostMethod,
		Path:   "/transactions/:transactionId/materials",
		Middlewares: []fiber.Handler{
			r.Middleware.Authenticated(),
			r.Middleware.Authorized([]string{"transactions.materials.create"}),
		},
		Handler: func(c *fiber.Ctx) error {
			var params CreateParams

			if err := c.ParamsParser(&params); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   constants.BadRequestError,
					"message": constants.BadRequestErrorDetails,
				})
			}

			var payload models.CreateTransactionMaterialPayload

			if err := c.BodyParser(&payload); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   constants.BadRequestError,
					"message": constants.BadRequestErrorDetails,
				})
			}

			if err := r.Services.Transactions().Materials().Create(payload); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   constants.InternalServerError,
					"message": constants.InternalServerErrorDetails,
				})
			}

			return c.SendStatus(fiber.StatusOK)
		},
	}
}
