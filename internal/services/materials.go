package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type materialsService interface {
	Create(material models.Material) error
	Update(materialId uuid.UUID, material models.Material) error
	Delete(materialId uuid.UUID) error
	Find(materialId uuid.UUID) (*models.Material, error)
	List() ([]models.Material, error)
	Count() (int64, error)
}

type materials struct {
	storage storage.Storage
}

func newMaterialsService(storage storage.Storage) materialsService {
	return &materials{
		storage: storage,
	}
}

func (s *materials) Create(material models.Material) error {
	if err := s.storage.Postgres.Create(&material).Error; err != nil {
		return err
	}

	return nil
}

func (s *materials) Update(materialId uuid.UUID, material models.Material) error {
	if err := s.storage.Postgres.Where("id = ?", materialId).Updates(&material).Error; err != nil {
		return err
	}

	return nil
}

func (s *materials) Delete(materialId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", materialId).Delete(&models.Material{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *materials) Find(materialId uuid.UUID) (*models.Material, error) {
	var material *models.Material

	if err := s.storage.Postgres.Where("id = ?", materialId).First(&material).Error; err != nil {
		return nil, err
	}

	return material, nil
}

func (s *materials) List() ([]models.Material, error) {
	var materials []models.Material

	if err := s.storage.Postgres.Find(&materials).Error; err != nil {
		return nil, err
	}

	return materials, nil
}

func (s *materials) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Material{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
