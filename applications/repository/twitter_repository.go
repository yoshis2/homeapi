package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yoshis2/homeapi/domain"
)

// TwitterRepository Twitter Repository
type TwitterRepository interface {
	Insert(*gorm.DB, *domain.Twitter) error
	Get(*gorm.DB, int) (*domain.Twitter, error)
	Last(*gorm.DB) (int, error)
}
