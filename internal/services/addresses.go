package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type addressesService interface {
	Create(payload models.CreateAddressPayload) (uuid.UUID, error)
	Update(addressId uuid.UUID, payload models.UpdateAddressPayload) error
	Delete(addressId uuid.UUID) error
	Find(addressId uuid.UUID) (*models.Address, error)
	List(clauses ...clause.Expression) ([]models.Address, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type addresses struct {
	storage storage.Storage
}

func newAddressesService(storage storage.Storage) addressesService {
	return &addresses{
		storage: storage,
	}
}

func (s *addresses) Create(payload models.CreateAddressPayload) (uuid.UUID, error) {
	var address models.Address

	address.LineOne = payload.LineOne
	address.LineTwo = payload.LineTwo
	address.City = payload.City
	address.ZipCode = payload.ZipCode
	address.Province = payload.Province
	address.Country = payload.Country

	if err := s.storage.Postgres.
		Create(&address).Error; err != nil {
		return uuid.Nil, err
	}

	return address.Id, nil
}

func (s *addresses) Update(addressId uuid.UUID, payload models.UpdateAddressPayload) error {
	var address models.Address

	if err := s.storage.Postgres.
		Where("id = ?", addressId).
		First(&address).Error; err != nil {
		return err
	}

	if payload.LineOne != nil {
		address.LineOne = *payload.LineOne
	}

	if payload.LineTwo != nil {
		address.LineTwo = payload.LineTwo
	}

	if payload.City != nil {
		address.City = *payload.City
	}

	if payload.ZipCode != nil {
		address.ZipCode = *payload.ZipCode
	}

	if payload.Province != nil {
		address.Province = *payload.Province
	}

	if payload.Country != nil {
		address.Country = *payload.Country
	}

	if err := s.storage.Postgres.
		Model(&models.Address{}).
		Where("id = ?", addressId).
		Updates(&map[string]any{
			"line_one": address.LineOne,
			"line_two": address.LineTwo,
			"city":     address.City,
			"zip_code": address.ZipCode,
			"province": address.Province,
			"country":  address.Country,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *addresses) Delete(addressId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", addressId).
		Delete(&models.Address{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *addresses) Find(addressId uuid.UUID) (*models.Address, error) {
	var address *models.Address

	if err := s.storage.Postgres.
		Where("id = ?", addressId).
		First(&address).Error; err != nil {
		return nil, err
	}

	return address, nil
}

func (s *addresses) List(clauses ...clause.Expression) ([]models.Address, error) {
	var addresses []models.Address

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

func (s *addresses) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Address{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
