package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yoshis2/homeapi/domain"
)

type TemperatureRepository struct {
}

// List 過去の気温データを抽出する
func (repo *TemperatureRepository) List(db *gorm.DB) ([]domain.Temperature, error) {
	var temperatures []domain.Temperature
	err := db.Order("created_at desc").Limit(200).Find(&temperatures).Error
	return temperatures, err
}

// Insert 気温DBにデータを挿入
func (repo *TemperatureRepository) Insert(db *gorm.DB, temperature *domain.Temperature) error {
	return db.Create(&temperature).Error
}
