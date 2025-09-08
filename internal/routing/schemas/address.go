package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var AddressSchema = openapi3.NewSchema().
	WithProperties(properties.AddressProperties).
	WithRequired([]string{
		"id",
		"lineOne",
		"lineTwo",
		"city",
		"zipCode",
		"province",
		"country",
		"createdAt",
		"updatedAt",
	}).NewRef()

var AddressesSchema = openapi3.NewArraySchema().WithItems(AddressSchema.Value).NewRef()

var CreateAddressSchema = openapi3.NewSchema().
	WithProperties(properties.CreateAddressProperties).
	WithRequired([]string{
		"lineOne",
		"city",
		"zipCode",
		"province",
		"country",
	}).NewRef()

var UpdateAddressSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateAddressProperties).
	NewRef()
