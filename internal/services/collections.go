package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type collectionsService interface {
	Materials() collectionMaterialsService
	Create(collection models.Collection) error
	Update(collectionId uuid.UUID, collection models.Collection) error
	Delete(collectionId uuid.UUID) error
	Find(collectionId uuid.UUID) (*models.Collection, error)
	List() ([]models.Collection, error)
	Count() (int64, error)
}

type collections struct {
	storage   storage.Storage
	materials collectionMaterialsService
}

func newCollectionsService(storage storage.Storage) collectionsService {
	materials := newCollectionMaterialsService(storage)

	return &collections{
		storage:   storage,
		materials: materials,
	}
}

func (s *collections) Materials() collectionMaterialsService {
	return s.materials
}

func (s *collections) Create(collection models.Collection) error {
	if err := s.storage.Postgres.Create(&collection).Error; err != nil {
		return err
	}

	return nil
}

func (s *collections) Update(collectionId uuid.UUID, collection models.Collection) error {
	if err := s.storage.Postgres.Where("id = ?", collectionId).Updates(&collection).Error; err != nil {
		return err
	}

	return nil
}

func (s *collections) Delete(collectionId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", collectionId).Delete(&models.Collection{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *collections) Find(collectionId uuid.UUID) (*models.Collection, error) {
	var collection *models.Collection

	if err := s.storage.Postgres.Where("id = ?", collectionId).First(&collection).Error; err != nil {
		return nil, err
	}

	return collection, nil
}

func (s *collections) List() ([]models.Collection, error) {
	var collections []models.Collection

	if err := s.storage.Postgres.Find(&collections).Error; err != nil {
		return nil, err
	}

	return collections, nil
}

func (s *collections) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Collection{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
