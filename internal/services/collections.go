package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type collectionsService interface {
	Materials() collectionMaterialsService
	Create(payload models.CreateCollectionPayload) error
	Update(collectionId uuid.UUID, payload models.UpdateCollectionPayload) error
	Delete(collectionId uuid.UUID) error
	Find(collectionId uuid.UUID) (*models.Collection, error)
	List(clauses ...clause.Expression) ([]models.Collection, error)
	Count(clauses ...clause.Expression) (int64, error)
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

func (s *collections) Create(payload models.CreateCollectionPayload) error {
	if err := s.storage.Postgres.
		Model(&models.Collection{}).
		Create(&payload).Error; err != nil {
		return err
	}

	return nil
}

func (s *collections) Update(collectionId uuid.UUID, payload models.UpdateCollectionPayload) error {
	var collection models.Collection

	if err := s.storage.Postgres.Where("id = ?", collectionId).First(&collection).Error; err != nil {
		return err
	}

	if payload.SellerId != nil {
		collection.SellerId = *payload.SellerId
	}

	if payload.BuyerId != nil {
		collection.BuyerId = *payload.BuyerId
	}

	if err := s.storage.Postgres.
		Model(&models.Collection{}).
		Where("id = ?", collectionId).
		Updates(&map[string]any{
			"seller_id": collection.SellerId,
			"buyer_id":  collection.BuyerId,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *collections) Delete(collectionId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", collectionId).
		Delete(&models.Collection{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *collections) Find(collectionId uuid.UUID) (*models.Collection, error) {
	var collection *models.Collection

	if err := s.storage.Postgres.
		Where("id = ?", collectionId).
		First(&collection).Error; err != nil {
		return nil, err
	}

	return collection, nil
}

func (s *collections) List(clauses ...clause.Expression) ([]models.Collection, error) {
	var collections []models.Collection

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&collections).Error; err != nil {
		return nil, err
	}

	return collections, nil
}

func (s *collections) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.Collection{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
