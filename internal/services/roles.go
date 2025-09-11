package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type rolesService interface {
	Create(payload models.CreateRolePayload) (uuid.UUID, error)
	Update(roleId uuid.UUID, payload models.UpdateRolePayload) error
	Delete(roleId uuid.UUID) error
	Find(roleId uuid.UUID) (*models.Role, error)
	List(clauses ...clause.Expression) ([]models.Role, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type roles struct {
	storage storage.Storage
}

func newRolesService(storage storage.Storage) rolesService {
	return &roles{
		storage: storage,
	}
}

func (s *roles) Create(payload models.CreateRolePayload) (uuid.UUID, error) {
	var role models.Role

	role.Name = payload.Name
	role.Description = payload.Description
	role.Permissions = payload.Permissions

	if err := s.storage.Postgres.
		Create(&role).Error; err != nil {
		return uuid.Nil, err
	}

	return role.Id, nil
}

func (s *roles) Update(roleId uuid.UUID, payload models.UpdateRolePayload) error {
	var role models.Role

	if err := s.storage.Postgres.
		Where("id = ?", roleId).
		First(&role).Error; err != nil {
		return err
	}

	if payload.Name != nil {
		role.Name = *payload.Name
	}

	role.Description = payload.Description

	if payload.Permissions != nil {
		role.Permissions = payload.Permissions
	}

	if err := s.storage.Postgres.
		Model(&models.Role{}).
		Where("id = ?", roleId).
		Updates(&map[string]any{
			"name":        role.Name,
			"description": role.Description,
			"permissions": role.Permissions,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *roles) Delete(roleId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", roleId).
		Delete(&models.Role{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *roles) Find(roleId uuid.UUID) (*models.Role, error) {
	var role *models.Role

	if err := s.storage.Postgres.
		Where("id = ?", roleId).
		First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (s *roles) List(clauses ...clause.Expression) ([]models.Role, error) {
	var roles []models.Role

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roles) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.Role{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
