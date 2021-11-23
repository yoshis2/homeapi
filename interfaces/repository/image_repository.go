package repository

import (
	"context"
	"homeapi/domain"

	"gorm.io/gorm"
)

type ImageRepository struct {
	Database *gorm.DB
}

// Insert 気温DBにデータを挿入
func (repo *ImageRepository) Insert(ctx context.Context, image *domain.Images) error {
	return repo.Database.WithContext(ctx).Create(&image).Error
}
