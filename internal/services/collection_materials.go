package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type collectionMaterialsService interface {
	Create(collectionMaterial models.CollectionMaterial) error
	Update(collectionMaterialId uuid.UUID, collectionMaterial models.CollectionMaterial) error
	Delete(collectionMaterialId uuid.UUID) error
	Find(collectionMaterialId uuid.UUID) (*models.CollectionMaterial, error)
	List() ([]models.CollectionMaterial, error)
	Count() (int64, error)
}

type collectionMaterials struct {
	storage storage.Storage
}

func newCollectionMaterialsService(storage storage.Storage) collectionMaterialsService {
	return &collectionMaterials{
		storage: storage,
	}
}

func (s *collectionMaterials) Create(collectionMaterial models.CollectionMaterial) error {
	if err := s.storage.Postgres.Create(&collectionMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (s *collectionMaterials) Update(collectionMaterialId uuid.UUID, collectionMaterial models.CollectionMaterial) error {
	if err := s.storage.Postgres.Where("id = ?", collectionMaterialId).Updates(&collectionMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (s *collectionMaterials) Delete(collectionMaterialId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", collectionMaterialId).Delete(&models.CollectionMaterial{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *collectionMaterials) Find(collectionMaterialId uuid.UUID) (*models.CollectionMaterial, error) {
	var collectionMaterial *models.CollectionMaterial

	if err := s.storage.Postgres.Where("id = ?", collectionMaterialId).First(&collectionMaterial).Error; err != nil {
		return nil, err
	}

	return collectionMaterial, nil
}

func (s *collectionMaterials) List() ([]models.CollectionMaterial, error) {
	var collectionMaterials []models.CollectionMaterial

	if err := s.storage.Postgres.Find(&collectionMaterials).Error; err != nil {
		return nil, err
	}

	return collectionMaterials, nil
}

func (s *collectionMaterials) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.CollectionMaterial{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
