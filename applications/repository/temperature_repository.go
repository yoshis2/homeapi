package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yoshis2/homeapi/domain"
)

// TemperatureRepository Temperature Repository
type TemperatureRepository interface {
	List(*gorm.DB) ([]domain.Temperature, error)
	Insert(*gorm.DB, *domain.Temperature) error
}
