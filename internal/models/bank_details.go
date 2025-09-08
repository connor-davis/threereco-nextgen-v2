package models

type BankDetails struct {
	Base
	AccountHolder string `json:"accountHolder" gorm:"type:text;not null"`
	AccountNumber string `json:"accountNumber" gorm:"type:text;not null"`
	BankName      string `json:"bankName" gorm:"type:text;not null"`
	BranchCode    string `json:"branchCode" gorm:"type:text;not null"`
}

type CreateBankDetailsPayload struct {
	AccountHolder string `json:"accountHolder"`
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	BranchCode    string `json:"branchCode"`
}

type UpdateBankDetailsPayload struct {
	AccountHolder *string `json:"accountHolder"`
	AccountNumber *string `json:"accountNumber"`
	BankName      *string `json:"bankName"`
	BranchCode    *string `json:"branchCode"`
}
