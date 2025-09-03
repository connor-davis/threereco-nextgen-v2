package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var OrganizationSchema = openapi3.NewSchema().
	WithProperties(properties.OrganizationProperties).
	WithProperty("address", AddressSchema.Value.WithNullable()).
	WithProperty("bankDetails", BankDetailSchema.Value.WithNullable()).
	WithRequired([]string{
		"id",
		"name",
		"description",
		"address",
		"bankDetails",
		"createdAt",
		"updatedAt",
	}).NewRef()

var OrganizationsSchema = openapi3.NewArraySchema().WithItems(OrganizationSchema.Value).NewRef()
