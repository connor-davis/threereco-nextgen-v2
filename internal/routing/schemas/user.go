package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var UserSchema = openapi3.NewSchema().
	WithProperties(properties.UserProperties).
	WithProperty("roles", RolesSchema.Value).
	WithProperty("address", AddressSchema.Value.WithNullable()).
	WithProperty("bankDetails", BankDetailSchema.Value.WithNullable()).
	WithRequired([]string{
		"id",
		"name",
		"email",
		"phone",
		"mfaEnabled",
		"mfaVerified",
		"roles",
		"banned",
		"banReason",
		"activeOrganization",
		"type",
		"address",
		"bankDetails",
		"createdAt",
		"updatedAt",
	}).NewRef()

var UsersSchema = openapi3.NewArraySchema().WithItems(UserSchema.Value).NewRef()

var CreateUserSchema = openapi3.NewSchema().
	WithProperties(properties.CreateUserProperties).
	WithProperty("roles", openapi3.NewArraySchema().WithItems(RoleSchema.Value)).
	WithRequired([]string{
		"name",
		"email",
		"phone",
		"password",
		"roles",
		"type",
	}).NewRef()

var UpdateUserSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateUserProperties).
	WithProperty("roles", openapi3.NewArraySchema().WithItems(RoleSchema.Value)).
	NewRef()
