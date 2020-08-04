package repository

import (
	"github.com/yoshis2/homeapi/domain"
	"github.com/jinzhu/gorm"
)

type ImageRepository struct {
}

// Insert 気温DBにデータを挿入
func (repo *ImageRepository) Insert(db *gorm.DB, image *domain.Images) error {
	return db.Create(&image).Error
}
