package ports

import "time"

type CsvUpdownInputPort struct {
	Name      string    `json:"name" example:"テスト　太郎"`
	Address   string    `json:"address" example:"東京都港区六本木１−１−１"`
	CreatedAt time.Time `json:"created_at" example:"2018-11-11 11:11:11"`
}

type CsvUpdownOutputPort struct {
	ID   uint
	Temp string
	Humi string
}
