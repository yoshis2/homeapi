package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"gorm.io/gorm"
)

//go:generate mockgen -source=./temperature_repository.go -package=repositorymock -destination=./mock/temperature_repository.go

func NewTemperatureRepository(db *gorm.DB) repository.TemperatureRepository {
	return repository.TemperatureRepository{
		Database: db,
	}
}

// TemperatureRepository Temperature Repository
type TemperatureRepository interface {
	List(ctx context.Context) ([]domain.Temperature, error)
	Insert(ctx context.Context, temperature *domain.Temperature) error
}
