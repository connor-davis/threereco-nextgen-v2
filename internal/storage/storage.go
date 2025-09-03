package storage

import (
	"github.com/connor-davis/threereco-nextgen/env"
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Postgres *gorm.DB
}

func New() Storage {
	return Storage{}
}

func (s *Storage) ConnectPostgres() {
	database, err := gorm.Open(postgres.Open(string(env.POSTGRES_DSN)), &gorm.Config{})

	if err != nil {
		log.Infof("🔥 Failed to connect to Postgres: %v", err)
		return
	}

	log.Info("✅ Successfully connected to Postgres")

	s.Postgres = database
}

func (s *Storage) MigratePostgres() {
	if s.Postgres == nil {
		log.Error("❌ Postgres connection is not established, cannot migrate")

		return
	}

	log.Info("🔄 Running Postgres migrations...")

	if err := s.Postgres.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`).Error; err != nil {
		log.Errorf("❌ Failed to create schema or extensions: %v", err)

		return
	} else {
		log.Info("✅ Extensions created successfully")
	}

	log.Info("🔃 Running GORM migrations...")

	if err := s.Postgres.AutoMigrate(
		&models.User{},
		&models.Organization{},
		&models.Role{},
		&models.Address{},
		&models.BankDetails{},
		&models.Material{},
		&models.Collection{},
		&models.CollectionMaterial{},
		&models.Transaction{},
		&models.TransactionMaterial{},
	); err != nil {
		log.Errorf("❌ AutoMigrate failed: %v", err)

		return
	}

	log.Info("✅ Postgres migrations completed successfully")
}

func (s *Storage) SeedPostgres() {
	if s.Postgres == nil {
		log.Error("❌ Postgres connection is not established, cannot seed")

		return
	}

	log.Info("🔄 Seeding Postgres...")

	log.Info("✅ Postgres seeding completed successfully.")
}
