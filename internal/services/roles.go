package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type rolesService interface {
	Create(role models.Role) error
	Update(roleId uuid.UUID, role models.Role) error
	Delete(roleId uuid.UUID) error
	Find(roleId uuid.UUID) (*models.Role, error)
	List() ([]models.Role, error)
	Count() (int64, error)
}

type roles struct {
	storage storage.Storage
}

func newRolesService(storage storage.Storage) rolesService {
	return &roles{
		storage: storage,
	}
}

func (s *roles) Create(role models.Role) error {
	if err := s.storage.Postgres.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func (s *roles) Update(roleId uuid.UUID, role models.Role) error {
	if err := s.storage.Postgres.Where("id = ?", roleId).Updates(&role).Error; err != nil {
		return err
	}

	return nil
}

func (s *roles) Delete(roleId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", roleId).Delete(&models.Role{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *roles) Find(roleId uuid.UUID) (*models.Role, error) {
	var role *models.Role

	if err := s.storage.Postgres.Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (s *roles) List() ([]models.Role, error) {
	var roles []models.Role

	if err := s.storage.Postgres.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roles) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Role{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
