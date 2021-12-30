package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"gorm.io/gorm"
)

//go:generate mockgen -source=./temperature_repository.go -package=repositorymock -destination=./mock/temperature_repository.go

func NewThermometerRepository(db *gorm.DB) repository.ThermometerRepository {
	return repository.ThermometerRepository{
		Database: db,
	}
}

// ThermometerRepository Temperature Repository
type ThermometerRepository interface {
	List(ctx context.Context) ([]domain.Thermometer, error)
	Insert(ctx context.Context, temperature *domain.Thermometer) error
}
