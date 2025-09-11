package storage

import (
	"github.com/connor-davis/threereco-nextgen/env"
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	if err := s.Postgres.Where("name = ?", string(env.DEFAULT_ORGANIZATION_NAME)).First(&models.Organization{}).Error; err == nil {
		if err != gorm.ErrRecordNotFound {
			log.Info("✅ Default organization admin user already exists, skipping seeding")

			return
		}

		log.Errorf("❌ Failed to check for existing organization admin user: %v", err)

		return
	}

	newUserId := uuid.New()
	newOrganizationId := uuid.New()
	newOrganizationAdminRoleId := uuid.New()
	newOrganizationUserRoleId := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(env.DEFAULT_ORGANIZATION_ADMIN_PASSWORD), bcrypt.DefaultCost)

	if err != nil {
		log.Errorf("🔥 Error hashing password: %s", err.Error())

		return
	}

	organizationAdminRoleDescription := "Has full access to the organization, including user management, settings, and data."
	organizationUserRoleDescription := "Has limited access to the organization, primarily for their own profile information or data related to them that has been created by other users with their information."

	organizationAdminRole := models.Role{
		Base: models.Base{
			Id: newOrganizationAdminRoleId,
		},
		Name:        "Administrator",
		Description: &organizationAdminRoleDescription,
		Permissions: []string{"*"},
	}

	organizationUserRole := models.Role{
		Base: models.Base{
			Id: newOrganizationUserRoleId,
		},
		Name:        "User",
		Description: &organizationUserRoleDescription,
		Permissions: []string{"users.view.self", "users.update.self", "users.delete.self"},
	}

	newUser := models.User{
		Base: models.Base{
			Id: newUserId,
		},
		Name:               string(env.DEFAULT_ORGANIZATION_ADMIN_NAME),
		Email:              string(env.DEFAULT_ORGANIZATION_ADMIN_EMAIL),
		Phone:              string(env.DEFAULT_ORGANIZATION_ADMIN_PHONE),
		Password:           hashedPassword,
		ActiveOrganization: newOrganizationId,
		Roles:              []models.Role{organizationAdminRole},
		Type:               models.System,
	}

	newOrganization := models.Organization{
		Base: models.Base{
			Id: newOrganizationId,
		},
		Name:  string(env.DEFAULT_ORGANIZATION_NAME),
		Users: []models.User{newUser},
		Roles: []models.Role{organizationAdminRole, organizationUserRole},
	}

	if err := s.Postgres.
		Set("one:ignore_audit_log", true).
		Assign(&newOrganization).
		FirstOrCreate(&newOrganization).Error; err != nil {
		log.Errorf("❌ Failed to create organization admin user: %v", err)

		return
	}

	log.Info("✅ Postgres seeding completed successfully.")
}
