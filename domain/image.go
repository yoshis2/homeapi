package domain

import "time"

// Image は画像アップロードの構造体
type Image struct {
	ID        int       `gorm:"id"`
	Name      string    `gorm:"name"`
	Path      string    `gorm:"path"`
	CreatedAt time.Time `gorm:"created_at"`
}
