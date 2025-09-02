package models

import "github.com/google/uuid"

type Transaction struct {
	Base
	Materials      []TransactionMaterial `json:"materials" gorm:"foreignKey:TransactionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CollectorId    uuid.UUID             `json:"collectorId" gorm:"type:uuid;not null"`
	Collector      Collector             `json:"collector" gorm:"foreignKey:CollectorId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizationId uuid.UUID             `json:"organizationId" gorm:"type:uuid;not null"`
	Organization   Organization          `json:"organization" gorm:"foreignKey:OrganizationId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TransactionMaterial struct {
	Base
	TransactionId uuid.UUID `json:"transactionId" gorm:"type:uuid;not null"`
	MaterialId    uuid.UUID `json:"materialId" gorm:"type:uuid;not null"`
	Material      Material  `json:"material" gorm:"foreignKey:MaterialId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Weight        float64   `json:"weight" gorm:"type:decimal(10,2);not null"`
}
