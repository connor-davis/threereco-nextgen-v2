package models

type Material struct {
	Base
	Name         string `json:"name" gorm:"type:text;not null"`
	GWCode       string `json:"gwCode" gorm:"type:text;not null"`
	CarbonFactor string `json:"carbonFactor" gorm:"type:float;not null"`
}

type CreateMaterialPayload struct {
	Name         string `json:"name"`
	GWCode       string `json:"gwCode"`
	CarbonFactor string `json:"carbonFactor"`
}

type UpdateMaterialPayload struct {
	Name         *string `json:"name"`
	GWCode       *string `json:"gwCode"`
	CarbonFactor *string `json:"carbonFactor"`
}
