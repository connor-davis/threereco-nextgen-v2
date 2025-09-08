package models

import "github.com/google/uuid"

type Collection struct {
	Base
	Materials []CollectionMaterial `json:"materials" gorm:"foreignKey:CollectionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SellerId  uuid.UUID            `json:"-" gorm:"type:uuid;not null"`
	Seller    User                 `json:"seller" gorm:"foreignKey:SellerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BuyerId   uuid.UUID            `json:"-" gorm:"type:uuid;not null"`
	Buyer     Organization         `json:"buyer" gorm:"foreignKey:BuyerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCollectionPayload struct {
	SellerId uuid.UUID `json:"sellerId"`
	BuyerId  uuid.UUID `json:"buyerId"`
}

type UpdateCollectionPayload struct {
	SellerId *uuid.UUID `json:"sellerId"`
	BuyerId  *uuid.UUID `json:"buyerId"`
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
