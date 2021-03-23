package repository

import (
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

type ImageRepository struct {
	Database *gorm.DB
}

// Insert 気温DBにデータを挿入
func (repo *ImageRepository) Insert(image *domain.Images) error {
	return repo.Database.Create(&image).Error
}
