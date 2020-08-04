package repository

import (
	"github.com/yoshis2/homeapi/domain"
	"github.com/jinzhu/gorm"
)

// ImageRepository Image Repository
type ImageRepository interface {
	Insert(*gorm.DB, *domain.Images) error
}
