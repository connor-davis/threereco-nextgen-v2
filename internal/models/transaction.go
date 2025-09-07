package models

import "github.com/google/uuid"

type Transaction struct {
	Base
	Materials []TransactionMaterial `json:"materials" gorm:"foreignKey:TransactionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SellerId  uuid.UUID             `json:"-" gorm:"type:uuid;not null"`
	Seller    Organization          `json:"seller" gorm:"foreignKey:SellerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BuyerId   uuid.UUID             `json:"-" gorm:"type:uuid;not null"`
	Buyer     Organization          `json:"buyer" gorm:"foreignKey:BuyerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateTransactionPayload struct {
	SellerId uuid.UUID `json:"sellerId"`
	BuyerId  uuid.UUID `json:"buyerId"`
}

type UpdateTransactionPayload struct {
	SellerId *uuid.UUID `json:"sellerId"`
	BuyerId  *uuid.UUID `json:"buyerId"`
}

type TransactionMaterial struct {
	Base
	TransactionId uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	MaterialId    uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	Material      Material  `json:"material" gorm:"foreignKey:MaterialId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Weight        float64   `json:"weight" gorm:"type:decimal(10,2);not null"`
	Value         float64   `json:"value" gorm:"type:decimal(10,2);not null"`
}

type CreateTransactionMaterialPayload struct {
	MaterialId uuid.UUID `json:"materialId"`
	Weight     float64   `json:"weight"`
	Value      float64   `json:"value"`
}

type UpdateTransactionMaterialPayload struct {
	MaterialId *uuid.UUID `json:"materialId"`
	Weight     *float64   `json:"weight"`
	Value      *float64   `json:"value"`
}
