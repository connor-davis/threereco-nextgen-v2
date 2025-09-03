package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type usersService interface {
	Create(user models.User) error
	Update(userId uuid.UUID, user models.User) error
	Delete(userId uuid.UUID) error
	Find(userId uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	List() ([]models.User, error)
	Count() (int64, error)
}

type users struct {
	storage storage.Storage
}

func newUsersService(storage storage.Storage) usersService {
	return &users{
		storage: storage,
	}
}

func (s *users) Create(user models.User) error {
	if err := s.storage.Postgres.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *users) Update(userId uuid.UUID, user models.User) error {
	if err := s.storage.Postgres.Where("id = ?", userId).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *users) Delete(userId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", userId).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *users) Find(userId uuid.UUID) (*models.User, error) {
	var user *models.User

	if err := s.storage.Postgres.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *users) FindByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := s.storage.Postgres.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *users) FindByPhone(phone string) (*models.User, error) {
	var user *models.User

	if err := s.storage.Postgres.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *users) List() ([]models.User, error) {
	var users []models.User

	if err := s.storage.Postgres.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *users) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
