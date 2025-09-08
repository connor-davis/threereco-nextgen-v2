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

var CreateAddressProperties = map[string]*openapi3.Schema{
	"lineOne":  openapi3.NewStringSchema(),
	"lineTwo":  openapi3.NewStringSchema().WithNullable(),
	"city":     openapi3.NewStringSchema(),
	"zipCode":  openapi3.NewStringSchema(),
	"province": openapi3.NewStringSchema(),
	"country":  openapi3.NewStringSchema(),
}

var UpdateAddressProperties = map[string]*openapi3.Schema{
	"lineOne":  openapi3.NewStringSchema().WithNullable(),
	"lineTwo":  openapi3.NewStringSchema().WithNullable(),
	"city":     openapi3.NewStringSchema().WithNullable(),
	"zipCode":  openapi3.NewStringSchema().WithNullable(),
	"province": openapi3.NewStringSchema().WithNullable(),
	"country":  openapi3.NewStringSchema().WithNullable(),
}
