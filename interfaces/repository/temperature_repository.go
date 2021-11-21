package repository

import (
	"context"
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

// TemperatureController 気温制御 Controller
type TemperatureRepository struct {
	Database *gorm.DB
}

// List 過去の気温データを抽出する
func (repo *TemperatureRepository) List(ctx context.Context) ([]domain.Temperature, error) {
	var temperatures []domain.Temperature
	err := repo.Database.Order("created_at desc").Limit(200).Find(&temperatures).Error
	return temperatures, err
}

// Insert 気温DBにデータを挿入
func (repo *TemperatureRepository) Insert(ctx context.Context, temperature *domain.Temperature) error {
	return repo.Database.Create(&temperature).Error
}
