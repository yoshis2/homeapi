package repository

import (
	"context"
	"homeapi/domain"

	"gorm.io/gorm"
)

// ThermometerController 気温制御 Controller
type ThermometerRepository struct {
	Database *gorm.DB
}

// List 過去の気温データを抽出する
func (repo *ThermometerRepository) List(ctx context.Context) ([]domain.Temperature, error) {
	var temperatures []domain.Temperature
	err := repo.Database.Order("created_at desc").Limit(200).Find(&temperatures).Error
	return temperatures, err
}

// Insert 気温DBにデータを挿入
func (repo *ThermometerRepository) Insert(ctx context.Context, temperature *domain.Temperature) error {
	return repo.Database.Create(&temperature).Error
}
