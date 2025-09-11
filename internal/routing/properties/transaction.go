package properties

import "github.com/getkin/kin-openapi/openapi3"

var TransactionProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"sellerId":  openapi3.NewUUIDSchema(),
	"buyerId":   openapi3.NewUUIDSchema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}

var CreateTransactionProperties = map[string]*openapi3.Schema{
	"sellerId": openapi3.NewUUIDSchema(),
	"buyerId":  openapi3.NewUUIDSchema(),
}

var UpdateTransactionProperties = map[string]*openapi3.Schema{
	"sellerId": openapi3.NewUUIDSchema().WithNullable(),
	"buyerId":  openapi3.NewUUIDSchema().WithNullable(),
}

var TransactionMaterialProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"weight":    openapi3.NewFloat64Schema(),
	"value":     openapi3.NewFloat64Schema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}

var CreateTransactionMaterialProperties = map[string]*openapi3.Schema{
	"materialId": openapi3.NewUUIDSchema(),
	"weight":     openapi3.NewFloat64Schema(),
	"value":      openapi3.NewFloat64Schema(),
}

var UpdateTransactionMaterialProperties = map[string]*openapi3.Schema{
	"materialId": openapi3.NewUUIDSchema().WithNullable(),
	"weight":     openapi3.NewFloat64Schema().WithNullable(),
	"value":      openapi3.NewFloat64Schema().WithNullable(),
}
