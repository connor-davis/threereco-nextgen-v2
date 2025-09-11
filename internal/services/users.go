package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type usersService interface {
	Create(payload models.CreateUserPayload) (uuid.UUID, error)
	Update(userId uuid.UUID, payload models.UpdateUserPayload) error
	Delete(userId uuid.UUID) error
	Find(userId uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	List(clauses ...clause.Expression) ([]models.User, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type users struct {
	storage storage.Storage
}

func newUsersService(storage storage.Storage) usersService {
	return &users{
		storage: storage,
	}
}

func (s *users) Create(payload models.CreateUserPayload) (uuid.UUID, error) {
	var user models.User

	user.Name = payload.Name
	user.Email = payload.Email
	user.Phone = payload.Phone

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return uuid.Nil, err
	}

	user.Password = hashedPassword
	user.Type = payload.Type

	if err := s.storage.Postgres.
		Model(&models.User{}).
		Create(&user).Error; err != nil {
		return uuid.Nil, err
	}

	if payload.Roles != nil {
		if err := s.storage.Postgres.
			Model(&user).Association("Roles").Append(payload.Roles); err != nil {
			return uuid.Nil, err
		}
	}

	return user.Id, nil
}

func (s *users) Update(userId uuid.UUID, payload models.UpdateUserPayload) error {
	var user models.User

	if err := s.storage.Postgres.
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		return err
	}

	if payload.Name != nil {
		user.Name = *payload.Name
	}

	if payload.Email != nil {
		user.Email = *payload.Email
	}

	if payload.Phone != nil {
		user.Phone = *payload.Phone
	}

	if payload.Type != nil {
		user.Type = *payload.Type
	}

	if err := s.storage.Postgres.
		Model(&models.User{}).
		Where("id = ?", userId).
		Updates(&map[string]any{
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
			"type":  user.Type,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *users) Delete(userId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", userId).
		Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *users) Find(userId uuid.UUID) (*models.User, error) {
	var user *models.User

	if err := s.storage.Postgres.
		Where("id = ?", userId).
		Preload("Roles").
		Preload("Address").
		Preload("BankDetails").
		First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *users) FindByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := s.storage.Postgres.
		Where("email = ?", email).
		Preload("Roles").
		Preload("Address").
		Preload("BankDetails").
		First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *users) FindByPhone(phone string) (*models.User, error) {
	var user *models.User

	if err := s.storage.Postgres.
		Where("phone = ?", phone).
		Preload("Roles").
		Preload("Address").
		Preload("BankDetails").
		First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *users) List(clauses ...clause.Expression) ([]models.User, error) {
	var users []models.User

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *users) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.User{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
