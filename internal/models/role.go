package models

import "github.com/lib/pq"

type Role struct {
	Base
	Name        string         `json:"name" gorm:"type:text;not null;"`
	Description *string        `json:"description" gorm:"type:text;"`
	Permissions pq.StringArray `json:"permissions" gorm:"type:text[];not null;default:'{}'"`
}

type CreateRolePayload struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Permissions []string `json:"permissions"`
}

type UpdateRolePayload struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Permissions []string `json:"permissions"`
}
