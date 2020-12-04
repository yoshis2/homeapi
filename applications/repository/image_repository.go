package repository

import (
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source=./image_repository.go -package=repositorymock -destination=./mock/image_repository.go

// ImageRepository Image Repository
type ImageRepository interface {
	Insert(*gorm.DB, *domain.Images) error
}
