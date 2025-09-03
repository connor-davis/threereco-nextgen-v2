package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type organizationsService interface {
	Create(organization models.Organization) error
	Update(organizationId uuid.UUID, organization models.Organization) error
	Delete(organizationId uuid.UUID) error
	Find(organizationId uuid.UUID) (*models.Organization, error)
	List() ([]models.Organization, error)
	Count() (int64, error)
}

type organizations struct {
	storage storage.Storage
}

func newOrganizationsService(storage storage.Storage) organizationsService {
	return &organizations{
		storage: storage,
	}
}

func (s *organizations) Create(organization models.Organization) error {
	if err := s.storage.Postgres.Create(&organization).Error; err != nil {
		return err
	}

	return nil
}

func (s *organizations) Update(organizationId uuid.UUID, organization models.Organization) error {
	if err := s.storage.Postgres.Where("id = ?", organizationId).Updates(&organization).Error; err != nil {
		return err
	}

	return nil
}

func (s *organizations) Delete(organizationId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", organizationId).Delete(&models.Organization{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *organizations) Find(organizationId uuid.UUID) (*models.Organization, error) {
	var organization *models.Organization

	if err := s.storage.Postgres.Where("id = ?", organizationId).First(&organization).Error; err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *organizations) List() ([]models.Organization, error) {
	var organizations []models.Organization

	if err := s.storage.Postgres.Find(&organizations).Error; err != nil {
		return nil, err
	}

	return organizations, nil
}

func (s *organizations) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Organization{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
