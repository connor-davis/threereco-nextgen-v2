package properties

import "github.com/getkin/kin-openapi/openapi3"

var LoginPayloadProperties = map[string]*openapi3.Schema{
	"emailOrPhone": openapi3.NewStringSchema().WithFormat("email"),
	"password":     openapi3.NewStringSchema().WithMinLength(6).WithMaxLength(100),
}

var SignUpPayloadProperties = map[string]*openapi3.Schema{
	"name":     openapi3.NewStringSchema().WithMinLength(2).WithMaxLength(100),
	"email":    openapi3.NewStringSchema().WithFormat("email"),
	"phone":    openapi3.NewStringSchema().WithMinLength(10).WithMaxLength(15),
	"password": openapi3.NewStringSchema().WithMinLength(6).WithMaxLength(100),
	"type":     openapi3.NewStringSchema().WithEnum("standard", "collector", "business", "system").WithDefault("standard"),
}

var AvailablePermissionsProperties = map[string]*openapi3.Schema{
	"name": openapi3.NewStringSchema(),
	"permissions": openapi3.NewArraySchema().
		WithItems(openapi3.NewObjectSchema().
			WithProperties(map[string]*openapi3.Schema{
				"value":       openapi3.NewStringSchema(),
				"description": openapi3.NewStringSchema(),
			}).
			WithRequired([]string{
				"value",
				"description",
			})),
}
