package repository

import (
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source=./temperature_repository.go -package=repositorymock -destination=./mock/temperature_repository.go

func NewTemperatureRepository(db *gorm.DB) repository.TemperatureRepository {
	return repository.TemperatureRepository{
		Database: db,
	}
}

// TemperatureRepository Temperature Repository
type TemperatureRepository interface {
	List() ([]domain.Temperature, error)
	Insert(*domain.Temperature) error
}
