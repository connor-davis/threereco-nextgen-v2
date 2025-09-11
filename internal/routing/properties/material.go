package properties

import "github.com/getkin/kin-openapi/openapi3"

var MaterialProperties = map[string]*openapi3.Schema{
	"id":           openapi3.NewUUIDSchema(),
	"name":         openapi3.NewStringSchema(),
	"gwCode":       openapi3.NewStringSchema(),
	"carbonFactor": openapi3.NewStringSchema(),
	"createdAt":    openapi3.NewDateTimeSchema(),
	"updatedAt":    openapi3.NewDateTimeSchema(),
}

var CreateMaterialProperties = map[string]*openapi3.Schema{
	"name":         openapi3.NewStringSchema(),
	"gwCode":       openapi3.NewStringSchema(),
	"carbonFactor": openapi3.NewStringSchema(),
}

var UpdateMaterialProperties = map[string]*openapi3.Schema{
	"name":         openapi3.NewStringSchema(),
	"gwCode":       openapi3.NewStringSchema(),
	"carbonFactor": openapi3.NewStringSchema(),
}
