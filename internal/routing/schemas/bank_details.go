package schemas

import (
	"github.com/connor-davis/threereco-nextgen/internal/routing/properties"
	"github.com/getkin/kin-openapi/openapi3"
)

var BankDetailSchema = openapi3.NewSchema().
	WithProperties(properties.BankDetailsProperties).
	WithRequired([]string{
		"id",
		"accountNumber",
		"accountHolder",
		"bankName",
		"branchName",
		"createdAt",
		"updatedAt",
	}).NewRef()

var BankDetailsSchema = openapi3.NewArraySchema().WithItems(BankDetailSchema.Value).NewRef()

var CreateBankDetailSchema = openapi3.NewSchema().
	WithProperties(properties.CreateBankDetailsProperties).
	WithRequired([]string{
		"accountNumber",
		"accountHolder",
		"bankName",
		"branchName",
	}).NewRef()

var UpdateBankDetailSchema = openapi3.NewSchema().
	WithProperties(properties.UpdateBankDetailsProperties).
	NewRef()
