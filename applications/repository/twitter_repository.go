package repository

import (
	"context"
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

// TwitterRepository Twitter Repository
type TwitterRepository interface {
	Insert(ctx context.Context, db *gorm.DB, twitter *domain.Twitter) error
	Get(ctx context.Context, db *gorm.DB, ID int) (*domain.Twitter, error)
	Last(ctx context.Context, db *gorm.DB) (int, error)
}
