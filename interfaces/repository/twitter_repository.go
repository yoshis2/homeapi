package repository

import (
	"context"
	"homeapi/domain"

	"gorm.io/gorm"
)

type TwitterRepository struct {
	Database *gorm.DB
}

func (repo *TwitterRepository) Insert(ctx context.Context, twitter *domain.Twitter) error {
	return repo.Database.Create(&twitter).Error
}

func (repo *TwitterRepository) Get(ctx context.Context, ID int) (*domain.Twitter, error) {
	twitter := domain.Twitter{}
	err := repo.Database.Where("id = ?", ID).First(&twitter).Error
	return &twitter, err
}

func (repo *TwitterRepository) Last(ctx context.Context) (int, error) {
	twitter := domain.Twitter{}
	err := repo.Database.Last(&twitter).Error
	return twitter.ID, err
}
