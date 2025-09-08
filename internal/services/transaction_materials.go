package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/connor-davis/threereco-nextgen/internal/storage"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type transactionMaterialsService interface {
	Create(transactionMaterial models.CreateTransactionMaterialPayload) error
	Update(transactionMaterialId uuid.UUID, transactionMaterial models.UpdateTransactionMaterialPayload) error
	Delete(transactionMaterialId uuid.UUID) error
	Find(transactionMaterialId uuid.UUID) (*models.TransactionMaterial, error)
	List(clauses ...clause.Expression) ([]models.TransactionMaterial, error)
	Count(clauses ...clause.Expression) (int64, error)
}

type transactionMaterials struct {
	storage storage.Storage
}

func newTransactionMaterialsService(storage storage.Storage) transactionMaterialsService {
	return &transactionMaterials{
		storage: storage,
	}
}

func (s *transactionMaterials) Create(payload models.CreateTransactionMaterialPayload) error {
	if err := s.storage.Postgres.
		Model(&models.TransactionMaterial{}).
		Create(&payload).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactionMaterials) Update(transactionMaterialId uuid.UUID, payload models.UpdateTransactionMaterialPayload) error {
	var transactionMaterial models.TransactionMaterial

	if err := s.storage.Postgres.
		Where("id = ?", transactionMaterialId).
		First(&transactionMaterial).Error; err != nil {
		return err
	}

	if payload.TransactionId != nil {
		transactionMaterial.TransactionId = *payload.TransactionId
	}

	if payload.MaterialId != nil {
		transactionMaterial.MaterialId = *payload.MaterialId
	}

	if payload.Weight != nil {
		transactionMaterial.Weight = *payload.Weight
	}

	if payload.Value != nil {
		transactionMaterial.Value = *payload.Value
	}

	if err := s.storage.Postgres.
		Model(&models.TransactionMaterial{}).
		Where("id = ?", transactionMaterialId).
		Updates(&map[string]any{
			"transaction_id": transactionMaterial.TransactionId,
			"material_id":    transactionMaterial.MaterialId,
			"weight":         transactionMaterial.Weight,
			"value":          transactionMaterial.Value,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactionMaterials) Delete(transactionMaterialId uuid.UUID) error {
	if err := s.storage.Postgres.
		Where("id = ?", transactionMaterialId).
		Delete(&models.TransactionMaterial{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *transactionMaterials) Find(transactionMaterialId uuid.UUID) (*models.TransactionMaterial, error) {
	var transactionMaterial *models.TransactionMaterial

	if err := s.storage.Postgres.
		Where("id = ?", transactionMaterialId).
		First(&transactionMaterial).Error; err != nil {
		return nil, err
	}

	return transactionMaterial, nil
}

func (s *transactionMaterials) List(clauses ...clause.Expression) ([]models.TransactionMaterial, error) {
	var transactionMaterials []models.TransactionMaterial

	if err := s.storage.Postgres.
		Clauses(clauses...).
		Find(&transactionMaterials).Error; err != nil {
		return nil, err
	}

	return transactionMaterials, nil
}

func (s *transactionMaterials) Count(clauses ...clause.Expression) (int64, error) {
	var count int64

	if err := s.storage.Postgres.
		Model(&models.TransactionMaterial{}).
		Clauses(clauses...).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
