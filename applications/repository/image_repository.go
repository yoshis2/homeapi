package repository

import (
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

// ImageRepository Image Repository
type ImageRepository interface {
	Insert(*gorm.DB, *domain.Images) error
}
