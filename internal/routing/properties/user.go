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

var CreateUserProperties = map[string]*openapi3.Schema{
	"name":          openapi3.NewStringSchema(),
	"email":         openapi3.NewStringSchema(),
	"phone":         openapi3.NewStringSchema(),
	"password":      openapi3.NewStringSchema(),
	"roles":         openapi3.NewArraySchema().WithItems(openapi3.NewUUIDSchema()),
	"type":          openapi3.NewStringSchema().WithEnum("standard", "collector", "business", "system"),
	"addressId":     openapi3.NewUUIDSchema().WithNullable(),
	"bankDetailsId": openapi3.NewUUIDSchema().WithNullable(),
}

var UpdateUserProperties = map[string]*openapi3.Schema{
	"name":     openapi3.NewStringSchema().WithNullable(),
	"email":    openapi3.NewStringSchema().WithNullable(),
	"phone":    openapi3.NewStringSchema().WithNullable(),
	"password": openapi3.NewStringSchema().WithNullable(),
	"roles":    openapi3.NewArraySchema().WithItems(openapi3.NewUUIDSchema()).WithNullable(),
	"type":     openapi3.NewStringSchema().WithEnum("standard", "collector", "business", "system").WithNullable(),
	"banned":   openapi3.NewBoolSchema().WithNullable(),
	"banReason": openapi3.NewStringSchema().
		WithNullable(),
	"addressId":          openapi3.NewUUIDSchema().WithNullable(),
	"bankDetailsId":      openapi3.NewUUIDSchema().WithNullable(),
	"activeOrganization": openapi3.NewUUIDSchema().WithNullable(),
}
