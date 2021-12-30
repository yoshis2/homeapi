package domain

import "time"

// Temperature 温度の入出力構造体
type Twitter struct {
	ID        int       `gorm:"primary_key"`
	Tweet     string    `gorm:"tweet"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
