package domain

import "time"

// Images は画像アップロードの構造体
type Images struct {
	ImageName string    `gorm:"image_name"`
	ImagePath string    `gorm:"image_path"`
	CreatedAt time.Time `gorm:"created_at"`
}
