package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var LoginPayloadSchema = openapi3.NewSchema().
	WithProperties(properties.LoginPayloadProperties).
	WithRequired([]string{
		"emailOrPhone",
		"password",
	}).NewRef()

var SignUpPayloadSchema = openapi3.NewSchema().
	WithProperties(properties.SignUpPayloadProperties).
	WithRequired([]string{
		"name",
		"password",
	}).NewRef()

var AvailablePermissionsSchema = openapi3.NewArraySchema().
	WithItems(openapi3.NewSchema().
		WithProperties(properties.AvailablePermissionsProperties).
		WithRequired([]string{
			"name",
			"permissions",
		})).NewRef()
