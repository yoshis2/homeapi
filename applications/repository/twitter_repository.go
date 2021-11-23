package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/interfaces/repository"

	"gorm.io/gorm"
)

func NewTwitterRepository(db *gorm.DB) repository.TwitterRepository {
	return repository.TwitterRepository{
		Database: db,
	}
}

// TwitterRepository Twitter Repository
type TwitterRepository interface {
	Insert(ctx context.Context, twitter *domain.Twitter) error
	Get(ctx context.Context, ID int) (*domain.Twitter, error)
	Last(ctx context.Context) (int, error)
}
