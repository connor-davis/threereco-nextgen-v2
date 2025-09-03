package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type transactionMaterialsService interface {
	Create(transactionMaterial models.TransactionMaterial) error
	Update(transactionMaterialId uuid.UUID, transactionMaterial models.TransactionMaterial) error
	Delete(transactionMaterialId uuid.UUID) error
	Find(transactionMaterialId uuid.UUID) (*models.TransactionMaterial, error)
	List() ([]models.TransactionMaterial, error)
	Count() (int64, error)
}

type transactionMaterials struct {
	storage storage.Storage
}

func newTransactionMaterialsService(storage storage.Storage) transactionMaterialsService {
	return &transactionMaterials{
		storage: storage,
	}
}

func (s *transactionMaterials) Create(transactionMaterial models.TransactionMaterial) error {
	if err := s.storage.Postgres.Create(&transactionMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactionMaterials) Update(transactionMaterialId uuid.UUID, transactionMaterial models.TransactionMaterial) error {
	if err := s.storage.Postgres.Where("id = ?", transactionMaterialId).Updates(&transactionMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactionMaterials) Delete(transactionMaterialId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", transactionMaterialId).Delete(&models.TransactionMaterial{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactionMaterials) Find(transactionMaterialId uuid.UUID) (*models.TransactionMaterial, error) {
	var transactionMaterial *models.TransactionMaterial

	if err := s.storage.Postgres.Where("id = ?", transactionMaterialId).First(&transactionMaterial).Error; err != nil {
		return nil, err
	}

	return transactionMaterial, nil
}

func (s *transactionMaterials) List() ([]models.TransactionMaterial, error) {
	var transactionMaterials []models.TransactionMaterial

	if err := s.storage.Postgres.Find(&transactionMaterials).Error; err != nil {
		return nil, err
	}

	return transactionMaterials, nil
}

func (s *transactionMaterials) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.TransactionMaterial{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
