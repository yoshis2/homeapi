package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"gorm.io/gorm"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE
func NewThermometerRepository(db *gorm.DB) repository.ThermometerRepository {
	return repository.ThermometerRepository{
		Database: db,
	}
}

// ThermometerRepository Temperature Repository
type ThermometerRepository interface {
	List(ctx context.Context) ([]domain.Thermometer, error)
	Insert(ctx context.Context, thermometer *domain.Thermometer) error
}
