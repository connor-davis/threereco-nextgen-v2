package models

import "github.com/google/uuid"

type Collection struct {
	Base
	Materials      []CollectionMaterial `json:"materials" gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CollectorId    uuid.UUID            `json:"-" gorm:"type:uuid;not null"`
	Collector      User                 `json:"collector" gorm:"foreignKey:CollectorId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizationId uuid.UUID            `json:"-" gorm:"type:uuid;not null"`
	Organization   Organization         `json:"organization" gorm:"foreignKey:OrganizationId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCollectionPayload struct {
	CollectorId    uuid.UUID `json:"collectorId"`
	OrganizationId uuid.UUID `json:"organizationId"`
}

type UpdateCollectionPayload struct {
	CollectorId    *uuid.UUID `json:"collectorId"`
	OrganizationId *uuid.UUID `json:"organizationId"`
}

type CollectionMaterial struct {
	Base
	CollectionId uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	MaterialId   uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	Material     Material  `json:"material" gorm:"foreignKey:MaterialId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Weight       float64   `json:"weight" gorm:"type:decimal(10,2);not null"`
	Value        float64   `json:"value" gorm:"type:decimal(10,2);not null"`
}

type CreateCollectionMaterialPayload struct {
	MaterialId uuid.UUID `json:"materialId"`
	Weight     float64   `json:"weight"`
	Value      float64   `json:"value"`
}

type UpdateCollectionMaterialPayload struct {
	MaterialId *uuid.UUID `json:"materialId"`
	Weight     *float64   `json:"weight"`
	Value      *float64   `json:"value"`
}
