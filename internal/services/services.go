package services

import (
	"github.com/connor-davis/threereco-nextgen/internal/storage"
)

type Services struct {
	Storage *storage.Storage
}

func NewServices(storage *storage.Storage) *Services {
	return &Services{
		Storage: storage,
	}
}
