package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type transactionsService interface {
	Materials() transactionMaterialsService
	Create(transaction models.CreateTransactionPayload) error
	Update(transactionId uuid.UUID, transaction models.UpdateTransactionPayload) error
	Delete(transactionId uuid.UUID) error
	Find(transactionId uuid.UUID) (*models.Transaction, error)
	List(clauses ...clause.Expression) ([]models.Transaction, error)
	Count(clauses ...clause.Expression) (int64, error)
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

func (s *transactions) Create(payload models.CreateTransactionPayload) error {
	if err := s.storage.Postgres.
		Model(&models.Transaction{}).
		Create(&payload).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactions) Update(transactionId uuid.UUID, payload models.UpdateTransactionPayload) error {
	var transaction models.Transaction

	if err := s.storage.Postgres.
		Where("id = ?", transactionId).
		First(&transaction).Error; err != nil {
		return err
	}

	if payload.SellerId != nil {
		transaction.SellerId = *payload.SellerId
	}

	if payload.BuyerId != nil {
		transaction.BuyerId = *payload.BuyerId
	}

	if err := s.storage.Postgres.
		Model(&models.Transaction{}).
		Where("id = ?", transactionId).
		Updates(&map[string]any{
			"seller_id": transaction.SellerId,
			"buyer_id":  transaction.BuyerId,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactions) Delete(transactionId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", transactionId).
		Delete(&models.Transaction{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactions) Find(transactionId uuid.UUID) (*models.Transaction, error) {
	var transaction *models.Transaction

	if err := s.storage.Postgres.
		Where("id = ?", transactionId).
		First(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactions) List(clauses ...clause.Expression) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *transactions) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.Transaction{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
