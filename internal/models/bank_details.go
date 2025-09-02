package models

type BankDetails struct {
	Base
	AccountHolder string `json:"accountHolder" gorm:"type:text;not null"`
	AccountNumber string `json:"accountNumber" gorm:"type:text;not null"`
	BankName      string `json:"bankName" gorm:"type:text;not null"`
	BranchCode    string `json:"branchCode" gorm:"type:text;not null"`
}
