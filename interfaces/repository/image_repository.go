package repository

import (
	"context"
	"homeapi/domain"

	"gorm.io/gorm"
)

type ImageRepository struct {
	Database *gorm.DB
}

func (repo *ImageRepository) List(ctx context.Context) ([]domain.Image, error) {
	var images []domain.Image
	if err := repo.Database.WithContext(ctx).Order("created_at desc").Limit(200).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

// Insert 気温DBにデータを挿入
func (repo *ImageRepository) Insert(ctx context.Context, image *domain.Image) error {
	return repo.Database.WithContext(ctx).Create(&image).Error
}
