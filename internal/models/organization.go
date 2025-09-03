package models

import "github.com/google/uuid"

type Organization struct {
	Base
	Name          string       `json:"name" gorm:"type:text;not null"`
	Roles         []Role       `json:"roles" gorm:"many2many:organization_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Users         []User       `json:"users" gorm:"many2many:organization_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AddressId     uuid.UUID    `json:"-" gorm:"type:uuid;not null"`
	Address       *Address     `json:"address" gorm:"foreignKey:AddressId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BankDetailsId uuid.UUID    `json:"-" gorm:"type:uuid"`
	BankDetails   *BankDetails `json:"bankDetails" gorm:"foreignKey:BankDetailsId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
