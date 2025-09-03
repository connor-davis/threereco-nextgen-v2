package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var MaterialSchema = openapi3.NewSchema().
	WithProperties(properties.MaterialProperties).
	WithRequired([]string{
		"id",
		"name",
		"gwCode",
		"carbonFactor",
		"value",
		"createdAt",
		"updatedAt",
	}).NewRef()

var MaterialsSchema = openapi3.NewArraySchema().WithItems(MaterialSchema.Value).NewRef()
