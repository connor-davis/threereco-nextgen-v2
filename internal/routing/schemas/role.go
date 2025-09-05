package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var RoleSchema = openapi3.NewSchema().
	WithProperties(properties.RoleProperties).
	WithRequired([]string{
		"id",
		"name",
		"description",
		"permissions",
		"createdAt",
		"updatedAt",
	}).NewRef()

var RolesSchema = openapi3.NewArraySchema().WithItems(RoleSchema.Value).NewRef()

var CreateRoleSchema = openapi3.NewSchema().
	WithProperties(properties.CreateRoleProperties).
	WithRequired([]string{
		"name",
		"permissions",
	}).NewRef()

var UpdateRoleSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateRoleProperties).NewRef()
