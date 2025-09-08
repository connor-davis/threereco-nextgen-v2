package bankDetails

import (
	"github.com/connor-davis/threereco-nextgen/internal/constants"
	"github.com/connor-davis/threereco-nextgen/internal/routing"
	"github.com/connor-davis/threereco-nextgen/internal/routing/schemas"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FindParams struct {
	Id uuid.UUID `json:"id"`
}

func (r *BankDetailsRouter) FindRoute() routing.Route {
	responses := openapi3.NewResponses()

	responses.Set("200", &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Successful bank details retrieval.").
			WithJSONSchema(schemas.SuccessResponseSchema.Value).
			WithContent(openapi3.Content{
				"application/json": openapi3.NewMediaType().
					WithSchema(schemas.BankDetailSchema.Value).
					WithExample("example", map[string]any{
						"id":            uuid.New(),
						"accountHolder": "John Doe",
						"accountNumber": "12345678",
						"bankName":      "Bank of Examples",
						"branchCode":    "823494",
						"createdAt":     "2023-10-01T12:00:00Z",
						"updatedAt":     "2023-10-01T12:00:00Z",
					}),
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
			Value: openapi3.NewPathParameter("id").
				WithRequired(true).
				WithSchema(openapi3.NewUUIDSchema()),
		},
	}

	return routing.Route{
		OpenAPIMetadata: routing.OpenAPIMetadata{
			Summary:     "Find Bank Detail",
			Description: "Find existing bank details in the system.",
			Tags:        []string{"Bank Details"},
			Responses:   responses,
			Parameters:  parameters,
			RequestBody: nil,
		},
		Method: routing.GetMethod,
		Path:   "/bank-details/{id}",
		Middlewares: []fiber.Handler{
			r.Middleware.Authenticated(),
			r.Middleware.Authorized([]string{"bank_details.view"}),
		},
		Handler: func(c *fiber.Ctx) error {
			var params FindParams

			if err := c.ParamsParser(&params); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   constants.BadRequestError,
					"message": constants.BadRequestErrorDetails,
				})
			}

			bankDetails, err := r.Services.BankDetails().Find(params.Id)

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
				"item": bankDetails,
			})
		},
	}
}
