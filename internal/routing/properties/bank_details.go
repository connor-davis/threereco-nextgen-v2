package properties

import "github.com/getkin/kin-openapi/openapi3"

var BankDetailsProperties = map[string]*openapi3.Schema{
	"id":            openapi3.NewUUIDSchema(),
	"accountNumber": openapi3.NewStringSchema(),
	"accountHolder": openapi3.NewStringSchema(),
	"bankName":      openapi3.NewStringSchema(),
	"branchName":    openapi3.NewStringSchema(),
	"createdAt":     openapi3.NewDateTimeSchema(),
	"updatedAt":     openapi3.NewDateTimeSchema(),
}

var CreateBankDetailsProperties = map[string]*openapi3.Schema{
	"accountNumber": openapi3.NewStringSchema(),
	"accountHolder": openapi3.NewStringSchema(),
	"bankName":      openapi3.NewStringSchema(),
	"branchName":    openapi3.NewStringSchema(),
}

var UpdateBankDetailsProperties = map[string]*openapi3.Schema{
	"accountNumber": openapi3.NewStringSchema().WithNullable(),
	"accountHolder": openapi3.NewStringSchema().WithNullable(),
	"bankName":      openapi3.NewStringSchema().WithNullable(),
	"branchName":    openapi3.NewStringSchema().WithNullable(),
}
