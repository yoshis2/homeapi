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
func (repo *ThermometerRepository) List(ctx context.Context) ([]domain.Thermometer, error) {
	var thermometers []domain.Thermometer
	if err := repo.Database.WithContext(ctx).Order("created_at desc").Limit(200).Find(&thermometers).Error; err != nil {
		return nil, err
	}
	return thermometers, nil
}

// Insert 気温DBにデータを挿入
func (repo *ThermometerRepository) Insert(ctx context.Context, thermometer *domain.Thermometer) error {
	return repo.Database.WithContext(ctx).Create(&thermometer).Error
}
