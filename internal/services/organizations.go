package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type organizationsService interface {
	Create(payload models.CreateOrganizationPayload) error
	Update(organizationId uuid.UUID, payload models.UpdateOrganizationPayload) error
	Delete(organizationId uuid.UUID) error
	Find(organizationId uuid.UUID) (*models.Organization, error)
	List(clauses clause.Expression) ([]models.Organization, error)
	Count(clauses clause.Expression) (int64, error)
}

type organizations struct {
	storage storage.Storage
}

func newOrganizationsService(storage storage.Storage) organizationsService {
	return &organizations{
		storage: storage,
	}
}

func (s *organizations) Create(payload models.CreateOrganizationPayload) error {
	var organization models.Organization

	organization.Name = payload.Name

	if err := s.storage.Postgres.Create(&organization).Error; err != nil {
		return err
	}

	if payload.Users != nil {
		if err := s.storage.Postgres.
			Model(&organization).Association("Users").Append(payload.Users); err != nil {
			return err
		}
	}

	if payload.Roles != nil {
		if err := s.storage.Postgres.
			Model(&organization).Association("Roles").Append(payload.Roles); err != nil {
			return err
		}
	}

	return nil
}

func (s *organizations) Update(organizationId uuid.UUID, payload models.UpdateOrganizationPayload) error {
	var organization models.Organization

	if err := s.storage.Postgres.Where("id = ?", organizationId).First(&organization).Error; err != nil {
		return err
	}

	if payload.Name != nil {
		organization.Name = *payload.Name
	}

	if err := s.storage.Postgres.
		Model(&models.Organization{}).
		Where("id = ?", organizationId).
		Updates(&map[string]any{
			"name": organization.Name,
		}).Error; err != nil {
		return err
	}

	if payload.Users != nil {
		if err := s.storage.Postgres.
			Model(&organization).Association("Users").Replace(payload.Users); err != nil {
			return err
		}
	}

	if payload.Roles != nil {
		if err := s.storage.Postgres.
			Model(&organization).Association("Roles").Replace(payload.Roles); err != nil {
			return err
		}
	}

	return nil
}

func (s *organizations) Delete(organizationId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", organizationId).
		Delete(&models.Organization{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *organizations) Find(organizationId uuid.UUID) (*models.Organization, error) {
	var organization *models.Organization

	if err := s.storage.Postgres.
		Where("id = ?", organizationId).
		First(&organization).Error; err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *organizations) List(clauses clause.Expression) ([]models.Organization, error) {
	var organizations []models.Organization

	if err := s.storage.Postgres.
		Clauses(clauses).
		Find(&organizations).Error; err != nil {
		return nil, err
	}

	return organizations, nil
}

func (s *organizations) Count(clauses clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.Organization{}).
		Clauses(clauses).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
