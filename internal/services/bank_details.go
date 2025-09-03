package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type bankDetailsService interface {
	Create(bankDetails models.BankDetails) error
	Update(bankDetailsId uuid.UUID, bankDetails models.BankDetails) error
	Delete(bankDetailsId uuid.UUID) error
	Find(bankDetailsId uuid.UUID) (*models.BankDetails, error)
	List() ([]models.BankDetails, error)
	Count() (int64, error)
}

type bankDetails struct {
	storage storage.Storage
}

func newBankDetailsService(storage storage.Storage) bankDetailsService {
	return &bankDetails{
		storage: storage,
	}
}

func (s *bankDetails) Create(bankDetails models.BankDetails) error {
	if err := s.storage.Postgres.Create(&bankDetails).Error; err != nil {
		return err
	}

	return nil
}

func (s *bankDetails) Update(bankDetailsId uuid.UUID, bankDetails models.BankDetails) error {
	if err := s.storage.Postgres.Where("id = ?", bankDetailsId).Updates(&bankDetails).Error; err != nil {
		return err
	}

	return nil
}

func (s *bankDetails) Delete(bankDetailsId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", bankDetailsId).Delete(&models.BankDetails{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *bankDetails) Find(bankDetailsId uuid.UUID) (*models.BankDetails, error) {
	var bankDetails *models.BankDetails

	if err := s.storage.Postgres.Where("id = ?", bankDetailsId).First(&bankDetails).Error; err != nil {
		return nil, err
	}

	return bankDetails, nil
}

func (s *bankDetails) List() ([]models.BankDetails, error) {
	var bankDetails []models.BankDetails

	if err := s.storage.Postgres.Find(&bankDetails).Error; err != nil {
		return nil, err
	}

	return bankDetails, nil
}

func (s *bankDetails) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.BankDetails{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
