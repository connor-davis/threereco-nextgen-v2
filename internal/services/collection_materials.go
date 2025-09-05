package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type collectionMaterialsService interface {
	Create(collectionId uuid.UUID, payload models.CreateCollectionMaterialPayload) error
	Update(collectionId uuid.UUID, collectionMaterialId uuid.UUID, payload models.UpdateCollectionMaterialPayload) error
	Delete(collectionMaterialId uuid.UUID) error
	Find(collectionMaterialId uuid.UUID) (*models.CollectionMaterial, error)
	List(clauses ...clause.Expression) ([]models.CollectionMaterial, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type collectionMaterials struct {
	storage storage.Storage
}

func newCollectionMaterialsService(storage storage.Storage) collectionMaterialsService {
	return &collectionMaterials{
		storage: storage,
	}
}

func (s *collectionMaterials) Create(collectionId uuid.UUID, payload models.CreateCollectionMaterialPayload) error {
	var collectionMaterial models.CollectionMaterial

	collectionMaterial.CollectionId = collectionId
	collectionMaterial.MaterialId = payload.MaterialId
	collectionMaterial.Weight = payload.Weight
	collectionMaterial.Value = payload.Value

	if err := s.storage.Postgres.
		Create(&collectionMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (s *collectionMaterials) Update(collectionId uuid.UUID, collectionMaterialId uuid.UUID, payload models.UpdateCollectionMaterialPayload) error {
	var collectionMaterial models.CollectionMaterial

	if err := s.storage.Postgres.
		Where("id = ?", collectionMaterialId).
		First(&collectionMaterial).Error; err != nil {
		return err
	}

	if payload.MaterialId != nil {
		collectionMaterial.MaterialId = *payload.MaterialId
	}

	if payload.Weight != nil {
		collectionMaterial.Weight = *payload.Weight
	}

	if payload.Value != nil {
		collectionMaterial.Value = *payload.Value
	}

	if err := s.storage.Postgres.
		Model(&models.CollectionMaterial{}).
		Where("id = ?", collectionMaterialId).
		Updates(&map[string]any{
			"material_id": collectionMaterial.MaterialId,
			"weight":      collectionMaterial.Weight,
			"value":       collectionMaterial.Value,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *collectionMaterials) Delete(collectionMaterialId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", collectionMaterialId).
		Delete(&models.CollectionMaterial{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *collectionMaterials) Find(collectionMaterialId uuid.UUID) (*models.CollectionMaterial, error) {
	var collectionMaterial *models.CollectionMaterial

	if err := s.storage.Postgres.
		Where("id = ?", collectionMaterialId).
		First(&collectionMaterial).Error; err != nil {
		return nil, err
	}

	return collectionMaterial, nil
}

func (s *collectionMaterials) List(clauses ...clause.Expression) ([]models.CollectionMaterial, error) {
	var collectionMaterials []models.CollectionMaterial

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&collectionMaterials).Error; err != nil {
		return nil, err
	}

	return collectionMaterials, nil
}

func (s *collectionMaterials) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.CollectionMaterial{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
