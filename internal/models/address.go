package models

type Address struct {
	Base
	LineOne  string  `json:"lineOne" gorm:"type:text;not null"`
	LineTwo  *string `json:"lineTwo" gorm:"type:text;"`
	City     string  `json:"city" gorm:"type:text;not null"`
	ZipCode  string  `json:"zipCode" gorm:"type:text;not null"`
	Province string  `json:"province" gorm:"type:text;not null"`
	Country  string  `json:"country" gorm:"type:text;not null"`
}
