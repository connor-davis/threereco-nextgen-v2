package properties

import "github.com/getkin/kin-openapi/openapi3"

var RoleProperties = map[string]*openapi3.Schema{
	"id":          openapi3.NewUUIDSchema(),
	"name":        openapi3.NewStringSchema(),
	"description": openapi3.NewStringSchema().WithNullable(),
	"permissions": openapi3.NewArraySchema().WithItems(openapi3.NewStringSchema()),
	"createdAt":   openapi3.NewDateTimeSchema(),
	"updatedAt":   openapi3.NewDateTimeSchema(),
}
