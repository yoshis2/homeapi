package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"gorm.io/gorm"
)

//go:generate mockgen -source=./image_repository.go -package=repositorymock -destination=./mock/image_repository.go

func NewImageRepository(db *gorm.DB) repository.ImageRepository {
	return repository.ImageRepository{
		Database: db,
	}
}

// ImageRepository Image Repository
type ImageRepository interface {
	Insert(ctx context.Context, image *domain.Image) error
}
