package ports

import "time"

type CsvUpdownInputPort struct {
	Name      string    `json:"name" validate:"required" example:"テスト　太郎"`
	Address   string    `json:"address" validate:"required" example:"東京都港区六本木１−１−１"`
	CreatedAt time.Time `json:"created_at" validate:"required" example:"2018-11-11 11:11:11"`
}

type CsvUpdownOutputPort struct {
	ID   uint   `json:"id"`
	Temp string `json:"temp"`
	Humi string `json:"humi" `
}
