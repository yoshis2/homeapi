package repository

import (
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source=./image_repository.go -package=repositorymock -destination=./mock/image_repository.go

func NewImageRepository(db *gorm.DB) repository.ImageRepository {
	return repository.ImageRepository{
		Database: db,
	}
}

// ImageRepository Image Repository
type ImageRepository interface {
	Insert(*domain.Images) error
}
