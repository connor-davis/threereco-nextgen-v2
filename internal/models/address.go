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

type CreateAddressPayload struct {
	LineOne  string  `json:"lineOne"`
	LineTwo  *string `json:"lineTwo"`
	City     string  `json:"city"`
	ZipCode  string  `json:"zipCode"`
	Province string  `json:"province"`
	Country  string  `json:"country"`
}

type UpdateAddressPayload struct {
	LineOne  *string `json:"lineOne"`
	LineTwo  *string `json:"lineTwo"`
	City     *string `json:"city"`
	ZipCode  *string `json:"zipCode"`
	Province *string `json:"province"`
	Country  *string `json:"country"`
}
