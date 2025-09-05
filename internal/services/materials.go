package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type materialsService interface {
	Create(payload models.CreateMaterialPayload) error
	Update(materialId uuid.UUID, payload models.UpdateMaterialPayload) error
	Delete(materialId uuid.UUID) error
	Find(materialId uuid.UUID) (*models.Material, error)
	List(clauses ...clause.Expression) ([]models.Material, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type materials struct {
	storage storage.Storage
}

func newMaterialsService(storage storage.Storage) materialsService {
	return &materials{
		storage: storage,
	}
}

func (s *materials) Create(payload models.CreateMaterialPayload) error {
	if err := s.storage.Postgres.
		Model(&models.Material{}).
		Create(&payload).Error; err != nil {
		return err
	}

	return nil
}

func (s *materials) Update(materialId uuid.UUID, payload models.UpdateMaterialPayload) error {
	if err := s.storage.Postgres.
		Model(&models.Material{}).
		Where("id = ?", materialId).
		Updates(&payload).Error; err != nil {
		return err
	}

	return nil
}

func (s *materials) Delete(materialId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", materialId).
		Delete(&models.Material{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *materials) Find(materialId uuid.UUID) (*models.Material, error) {
	var material *models.Material

	if err := s.storage.Postgres.
		Where("id = ?", materialId).
		First(&material).Error; err != nil {
		return nil, err
	}

	return material, nil
}

func (s *materials) List(clauses ...clause.Expression) ([]models.Material, error) {
	var materials []models.Material

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&materials).Error; err != nil {
		return nil, err
	}

	return materials, nil
}

func (s *materials) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.Material{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
