package properties

import "github.com/getkin/kin-openapi/openapi3"

var CollectionProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}

var CreateCollectionProperties = map[string]*openapi3.Schema{
	"sellerId": openapi3.NewUUIDSchema(),
	"buyerId":  openapi3.NewUUIDSchema(),
}

var UpdateCollectionProperties = map[string]*openapi3.Schema{
	"sellerId": openapi3.NewUUIDSchema().WithNullable(),
	"buyerId":  openapi3.NewUUIDSchema().WithNullable(),
}

var CollectionMaterialProperties = map[string]*openapi3.Schema{
	"id":        openapi3.NewUUIDSchema(),
	"weight":    openapi3.NewFloat64Schema(),
	"createdAt": openapi3.NewDateTimeSchema(),
	"updatedAt": openapi3.NewDateTimeSchema(),
}

var CreateCollectionMaterialProperties = map[string]*openapi3.Schema{
	"collectionId": openapi3.NewUUIDSchema(),
	"materialId":   openapi3.NewUUIDSchema(),
	"weight":       openapi3.NewFloat64Schema(),
	"value":        openapi3.NewFloat64Schema(),
}

var UpdateCollectionMaterialProperties = map[string]*openapi3.Schema{
	"collectionId": openapi3.NewUUIDSchema().WithNullable(),
	"materialId":   openapi3.NewUUIDSchema().WithNullable(),
	"weight":       openapi3.NewFloat64Schema().WithNullable(),
	"value":        openapi3.NewFloat64Schema().WithNullable(),
}
