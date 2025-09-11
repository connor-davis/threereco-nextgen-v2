package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type bankDetailsService interface {
	Create(payload models.CreateBankDetailsPayload) (uuid.UUID, error)
	Update(bankDetailsId uuid.UUID, payload models.UpdateBankDetailsPayload) error
	Delete(bankDetailsId uuid.UUID) error
	Find(bankDetailsId uuid.UUID) (*models.BankDetails, error)
	List(clauses ...clause.Expression) ([]models.BankDetails, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type bankDetails struct {
	storage storage.Storage
}

func newBankDetailsService(storage storage.Storage) bankDetailsService {
	return &bankDetails{
		storage: storage,
	}
}

func (s *bankDetails) Create(payload models.CreateBankDetailsPayload) (uuid.UUID, error) {
	var bankDetails models.BankDetails

	bankDetails.AccountHolder = payload.AccountHolder
	bankDetails.AccountNumber = payload.AccountNumber
	bankDetails.BankName = payload.BankName
	bankDetails.BranchCode = payload.BranchCode

	if err := s.storage.Postgres.
		Create(&bankDetails).Error; err != nil {
		return uuid.Nil, err
	}

	return bankDetails.Id, nil
}

func (s *bankDetails) Update(bankDetailsId uuid.UUID, payload models.UpdateBankDetailsPayload) error {
	var bankDetails models.BankDetails

	if err := s.storage.Postgres.
		Where("id = ?", bankDetailsId).
		First(&bankDetails).Error; err != nil {
		return err
	}

	if payload.AccountHolder != nil {
		bankDetails.AccountHolder = *payload.AccountHolder
	}

	if payload.AccountNumber != nil {
		bankDetails.AccountNumber = *payload.AccountNumber
	}

	if payload.BankName != nil {
		bankDetails.BankName = *payload.BankName
	}

	if payload.BranchCode != nil {
		bankDetails.BranchCode = *payload.BranchCode
	}

	if err := s.storage.Postgres.
		Model(&models.BankDetails{}).
		Where("id = ?", bankDetailsId).
		Updates(&map[string]any{
			"account_holder": bankDetails.AccountHolder,
			"account_number": bankDetails.AccountNumber,
			"bank_name":      bankDetails.BankName,
			"branch_code":    bankDetails.BranchCode,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *bankDetails) Delete(bankDetailsId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", bankDetailsId).
		Delete(&models.BankDetails{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *bankDetails) Find(bankDetailsId uuid.UUID) (*models.BankDetails, error) {
	var bankDetails *models.BankDetails

	if err := s.storage.Postgres.
		Where("id = ?", bankDetailsId).
		First(&bankDetails).Error; err != nil {
		return nil, err
	}

	return bankDetails, nil
}

func (s *bankDetails) List(clauses ...clause.Expression) ([]models.BankDetails, error) {
	var bankDetails []models.BankDetails

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&bankDetails).Error; err != nil {
		return nil, err
	}

	return bankDetails, nil
}

func (s *bankDetails) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.BankDetails{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
