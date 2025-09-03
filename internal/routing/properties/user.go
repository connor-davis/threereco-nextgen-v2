package properties

import "github.com/getkin/kin-openapi/openapi3"

var UserProperties = map[string]*openapi3.Schema{
	"id":                 openapi3.NewUUIDSchema(),
	"name":               openapi3.NewStringSchema(),
	"email":              openapi3.NewStringSchema(),
	"phone":              openapi3.NewStringSchema(),
	"mfaEnabled":         openapi3.NewBoolSchema(),
	"mfaVerified":        openapi3.NewBoolSchema(),
	"banned":             openapi3.NewBoolSchema(),
	"banReason":          openapi3.NewStringSchema().WithNullable(),
	"activeOrganization": openapi3.NewUUIDSchema(),
	"type":               openapi3.NewStringSchema().WithEnum("standard", "collector", "business", "system"),
	"createdAt":          openapi3.NewDateTimeSchema(),
	"updatedAt":          openapi3.NewDateTimeSchema(),
}
