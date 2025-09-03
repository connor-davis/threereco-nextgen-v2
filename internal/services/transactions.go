package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
)

type transactionsService interface {
	Materials() transactionMaterialsService
	Create(transaction models.Transaction) error
	Update(transactionId uuid.UUID, transaction models.Transaction) error
	Delete(transactionId uuid.UUID) error
	Find(transactionId uuid.UUID) (*models.Transaction, error)
	List() ([]models.Transaction, error)
	Count() (int64, error)
}

type transactions struct {
	storage   storage.Storage
	materials transactionMaterialsService
}

func newTransactionsService(storage storage.Storage) transactionsService {
	materials := newTransactionMaterialsService(storage)

	return &transactions{
		storage:   storage,
		materials: materials,
	}
}

func (s *transactions) Materials() transactionMaterialsService {
	return s.materials
}

func (s *transactions) Create(transaction models.Transaction) error {
	if err := s.storage.Postgres.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactions) Update(transactionId uuid.UUID, transaction models.Transaction) error {
	if err := s.storage.Postgres.Where("id = ?", transactionId).Updates(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactions) Delete(transactionId uuid.UUID) error {
	if err := s.storage.Postgres.Where("id = ?", transactionId).Delete(&models.Transaction{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactions) Find(transactionId uuid.UUID) (*models.Transaction, error) {
	var transaction *models.Transaction

	if err := s.storage.Postgres.Where("id = ?", transactionId).First(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactions) List() ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := s.storage.Postgres.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *transactions) Count() (int64, error) {
	var count int64

	if err := s.storage.Postgres.Model(&models.Transaction{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
