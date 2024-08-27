package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"gorm.io/gorm"
)

//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE
func NewImageRepository(db *gorm.DB) repository.ImageRepository {
	return repository.ImageRepository{
		Database: db,
	}
}

// ImageRepository Image Repository
type ImageRepository interface {
	List(ctx context.Context) ([]domain.Image, error)
	Insert(ctx context.Context, image *domain.Image) error
}
