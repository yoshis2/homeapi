package repository

import (
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

// TemperatureRepository Temperature Repository
type TemperatureRepository interface {
	List(*gorm.DB) ([]domain.Temperature, error)
	Insert(*gorm.DB, *domain.Temperature) error
}
