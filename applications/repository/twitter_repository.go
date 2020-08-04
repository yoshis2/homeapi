package repository

import (
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

// TwitterRepository Twitter Repository
type TwitterRepository interface {
	Insert(*gorm.DB, *domain.Twitter) error
	Get(*gorm.DB, int) (*domain.Twitter, error)
	Last(*gorm.DB) (int, error)
}
