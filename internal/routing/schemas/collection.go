package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var CollectionMaterialSchema = openapi3.NewSchema().
	WithProperties(properties.CollectionMaterialProperties).
	WithProperty("material", MaterialSchema.Value).
	WithRequired([]string{
		"id",
		"material",
		"weight",
		"createdAt",
		"updatedAt",
	}).NewRef()

var CollectionMaterialsArraySchema = openapi3.NewArraySchema().WithItems(CollectionMaterialSchema.Value).NewRef()

var CreateCollectionMaterialSchema = openapi3.NewSchema().
	WithProperties(properties.CreateCollectionMaterialProperties).
	WithRequired([]string{
		"materialId",
		"weight",
		"value",
	}).NewRef()

var UpdateCollectionMaterialSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateCollectionMaterialProperties).
	NewRef()

var CollectionSchema = openapi3.NewSchema().
	WithProperties(properties.CollectionProperties).
	WithProperty("collector", UserSchema.Value).
	WithProperty("organization", OrganizationSchema.Value).
	WithProperty("materials", CollectionMaterialsArraySchema.Value).
	WithRequired([]string{
		"id",
		"collector",
		"organization",
		"materials",
		"createdAt",
		"updatedAt",
	}).NewRef()

var CollectionsSchema = openapi3.NewArraySchema().WithItems(CollectionSchema.Value).NewRef()

var CreateCollectionSchema = openapi3.NewSchema().
	WithProperties(properties.CreateCollectionProperties).
	WithRequired([]string{
		"collectorId",
		"organizationId",
	}).NewRef()

var UpdateCollectionSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateCollectionProperties).
	NewRef()
