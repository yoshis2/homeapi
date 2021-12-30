package domain

import "time"

// Temperature 温度の入出力構造体
type Thermometer struct {
	ID          uint      `gorm:"primary_key" csv:"ID"`
	Temperature string    `gorm:"not null" csv:"温度"`
	Humidity    string    `gorm:"not null" csv:"湿度"`
	CreatedAt   time.Time `gorm:"created_at" csv:"作成日時"`
}
