package models

import "github.com/google/uuid"

type Collector struct {
	Base
	Name          string       `json:"name" gorm:"type:text;not null"`
	AddressId     uuid.UUID    `json:"addressId" gorm:"type:uuid;not null"`
	Address       Address      `json:"address" gorm:"foreignKey:AddressId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BankDetailsId uuid.UUID    `json:"bankDetailsId" gorm:"type:uuid"`
	BankDetails   *BankDetails `json:"bankDetails" gorm:"foreignKey:BankDetailsId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
