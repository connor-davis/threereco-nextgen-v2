package models

import "github.com/google/uuid"

type User struct {
	Base
	Name               string    `json:"name" gorm:"type:text;not null"`
	Email              string    `json:"email" gorm:"type:text;not null;unique"`
	Phone              string    `json:"phone" gorm:"type:text;not null;unique"`
	Password           []byte    `json:"password" gorm:"type:bytea;not null"`
	MfaSecret          *string   `json:"mfaSecret" gorm:"type:text"`
	MfaEnabled         bool      `json:"mfaEnabled" gorm:"type:boolean;default:false"`
	MfaVerified        bool      `json:"mfaVerified" gorm:"type:boolean;default:false"`
	Roles              []Role    `json:"roles" gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banned             bool      `json:"banned" gorm:"type:boolean;default:false"`
	BanReason          *string   `json:"banReason" gorm:"type:text"`
	ActiveOrganization uuid.UUID `json:"activeOrganization" gorm:"type:uuid"`
}
