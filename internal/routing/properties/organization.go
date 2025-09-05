package properties

import "github.com/getkin/kin-openapi/openapi3"

var OrganizationProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"name":      openapi3.NewStringSchema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}

var CreateOrganizationProperties = map[string]*openapi3.Schema{
	"name": openapi3.NewStringSchema(),
}

var UpdateOrganizationProperties = map[string]*openapi3.Schema{
	"name": openapi3.NewStringSchema().WithNullable(),
}
