package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var TransactionMaterialSchema = openapi3.NewSchema().
	WithProperties(properties.TransactionMaterialProperties).
	WithProperty("material", MaterialSchema.Value).
	WithRequired([]string{
		"id",
		"material",
		"weight",
		"createdAt",
		"updatedAt",
	}).NewRef()

var TransactionMaterialsArraySchema = openapi3.NewArraySchema().WithItems(TransactionMaterialSchema.Value).NewRef()

var TransactionSchema = openapi3.NewSchema().
	WithProperties(properties.TransactionProperties).
	WithProperty("collector", UserSchema.Value).
	WithProperty("organization", OrganizationSchema.Value).
	WithProperty("materials", TransactionMaterialsArraySchema.Value).
	WithRequired([]string{
		"id",
		"collector",
		"organization",
		"materials",
		"createdAt",
		"updatedAt",
	}).NewRef()

var TransactionsSchema = openapi3.NewArraySchema().WithItems(TransactionSchema.Value).NewRef()
