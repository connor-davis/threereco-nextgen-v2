package properties

import "github.com/getkin/kin-openapi/openapi3"

var TransactionProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}

var TransactionMaterialProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"weight":    openapi3.NewFloat64Schema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}
