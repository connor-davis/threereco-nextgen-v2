package properties

import "github.com/getkin/kin-openapi/openapi3"

var AddressProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"lineOne":   openapi3.NewStringSchema(),
	"lineTwo":   openapi3.NewStringSchema().WithNullable(),
	"city":      openapi3.NewStringSchema(),
	"zipCode":   openapi3.NewStringSchema(),
	"province":  openapi3.NewStringSchema(),
	"country":   openapi3.NewStringSchema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}
