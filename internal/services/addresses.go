package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type addressesService interface {
	Create(address models.Address) error
	Update(addressId uuid.UUID, address models.Address) error
	Delete(addressId uuid.UUID) error
	Find(addressId uuid.UUID) (*models.Address, error)
	List() ([]models.Address, error)
	Count() (int64, error)
}

type addresses struct {
	storage storage.Storage
}

func newAddressesService(storage storage.Storage) addressesService {
	return &addresses{
		storage: storage,
	}
}

func (s *addresses) Create(address models.Address) error {
	if err := s.storage.Postgres.Create(&address).Error; err != nil {
		return err
	}

	return nil
}

func (s *addresses) Update(addressId uuid.UUID, address models.Address) error {
	if err := s.storage.Postgres.Where("id = ?", addressId).Updates(&address).Error; err != nil {
		return err
	}

	return nil
}

func (s *addresses) Delete(addressId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", addressId).Delete(&models.Address{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *addresses) Find(addressId uuid.UUID) (*models.Address, error) {
	var address *models.Address

	if err := s.storage.Postgres.Where("id = ?", addressId).First(&address).Error; err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addresses) List() ([]models.Address, error) {
	var addresses []models.Address

	if err := s.storage.Postgres.Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

func (s *addresses) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Address{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
