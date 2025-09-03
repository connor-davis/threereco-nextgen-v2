package properties

import "github.com/getkin/kin-openapi/openapi3"

var MaterialProperties = map[string]*openapi3.Schema{
	"id":           openapi3.NewUUIDSchema(),
	"name":         openapi3.NewStringSchema(),
	"gwCode":       openapi3.NewStringSchema(),
	"carbonFactor": openapi3.NewStringSchema(),
	"value":        openapi3.NewFloat64Schema(),
	"createdAt":    openapi3.NewDateTimeSchema(),
	"updatedAt":    openapi3.NewDateTimeSchema(),
}
