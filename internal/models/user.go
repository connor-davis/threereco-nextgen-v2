package models

import "github.com/google/uuid"

type UserType string

const (
	Standard  UserType = "standard"
	Collector UserType = "collector"
	Business  UserType = "business"
	System    UserType = "system"
)

type User struct {
	Base
	Name               string       `json:"name" gorm:"type:text;not null"`
	Email              string       `json:"email" gorm:"type:text;not null;unique"`
	Phone              string       `json:"phone" gorm:"type:text;not null;unique"`
	Password           []byte       `json:"-" gorm:"type:bytea;not null"`
	MfaSecret          []byte       `json:"-" gorm:"type:bytea"`
	MfaEnabled         bool         `json:"mfaEnabled" gorm:"type:boolean;not null;default:false"`
	MfaVerified        bool         `json:"mfaVerified" gorm:"type:boolean;not null;default:false"`
	Roles              []Role       `json:"roles" gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banned             bool         `json:"banned" gorm:"type:boolean;not null;default:false"`
	BanReason          *string      `json:"banReason" gorm:"type:text"`
	ActiveOrganization uuid.UUID    `json:"activeOrganization" gorm:"type:uuid;not null"`
	Type               UserType     `json:"-" gorm:"type:text;default:'standard';not null"`
	AddressId          *uuid.UUID   `json:"-" gorm:"type:uuid"`
	Address            *Address     `json:"address" gorm:"foreignKey:AddressId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BankDetailsId      *uuid.UUID   `json:"-" gorm:"type:uuid"`
	BankDetails        *BankDetails `json:"bankDetails" gorm:"foreignKey:BankDetailsId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateUserPayload struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Password string   `json:"password"`
	Roles    []Role   `json:"roles"`
	Type     UserType `json:"type"`
}

type UpdateUserPayload struct {
	Name     *string   `json:"name"`
	Email    *string   `json:"email"`
	Phone    *string   `json:"phone"`
	Password *string   `json:"password"`
	Roles    []Role    `json:"roles"`
	Type     *UserType `json:"type"`
}
