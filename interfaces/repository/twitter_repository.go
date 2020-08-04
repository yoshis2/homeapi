package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yoshis2/homeapi/domain"
)

type TwitterRepository struct {
}

func (repo *TwitterRepository) Insert(db *gorm.DB, twitter *domain.Twitter) error {
	return db.Create(&twitter).Error
}

func (repo *TwitterRepository) Get(db *gorm.DB, ID int) (*domain.Twitter, error) {
	twitter := domain.Twitter{}
	err := db.Where("id = ?", ID).First(&twitter).Error
	return &twitter, err
}

func (repo *TwitterRepository) Last(db *gorm.DB) (int, error) {
	twitter := domain.Twitter{}
	err := db.Last(&twitter).Error
	return twitter.ID, err
}
