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
	WithProperty("seller", UserSchema.Value).
	WithProperty("buyer", OrganizationSchema.Value).
	WithProperty("materials", CollectionMaterialsArraySchema.Value).
	WithRequired([]string{
		"id",
		"seller",
		"buyer",
		"materials",
		"createdAt",
		"updatedAt",
	}).NewRef()

var CollectionsSchema = openapi3.NewArraySchema().WithItems(CollectionSchema.Value).NewRef()

var CreateCollectionSchema = openapi3.NewSchema().
	WithProperties(properties.CreateCollectionProperties).
	WithRequired([]string{
		"sellerId",
		"buyerId",
	}).NewRef()

var UpdateCollectionSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateCollectionProperties).
	NewRef()
