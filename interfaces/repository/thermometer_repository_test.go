package repository

import (
	"context"
	"homeapi/domain"
	"log"
	"regexp"
	"testing"
	"time"

	"homeapi/infrastructure/databases"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureList(t *testing.T) {
	ctx := context.Background()
	db, mock, err := databases.GormMock()
	if err != nil {
		t.Fatal(err)
	}

	repo := ThermometerRepository{
		Database: db,
	}

	nowTime := time.Now()

	room1 := &domain.Thermometer{
		ID:          3,
		Temperature: "22.5",
		Humidity:    "76",
		CreatedAt:   nowTime,
	}

	room2 := &domain.Thermometer{
		ID:          4,
		Temperature: "11.8",
		Humidity:    "60",
		CreatedAt:   nowTime,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `thermometers` ORDER BY created_at desc LIMIT ?")).
		WithArgs(200).
		WillReturnRows(sqlmock.NewRows([]string{"id", "temperature", "humidity", "created_at"}).
			AddRow(room1.ID, room1.Temperature, room1.Humidity, room1.CreatedAt).
			AddRow(4, "11.8", "60", nowTime),
		)

	res, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)

	assert.Equal(t, &res[0], room1)
	assert.Equal(t, &res[1], room2)
}

func TemperatureInsert(t *testing.T) {
	ctx := context.Background()
	db, mock, err := databases.GormMock()
	if err != nil {
		t.Fatal(err)
	}

	repo := ThermometerRepository{
		Database: db,
	}

	nowTime := time.Now()

	room1 := &domain.Thermometer{
		ID:          3,
		Temperature: "22.5",
		Humidity:    "76",
		CreatedAt:   nowTime,
	}

	insertTemperature := regexp.QuoteMeta(`INSERT INTO "temperatures" ("id","temp", "humi", "created_at") VALUES ($1, $2, $3, $4)`)
	mock.ExpectQuery(insertTemperature).
		WithArgs(room1.ID, room1.Temperature, room1.Humidity, room1.CreatedAt).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "temp", "humi", "created_at"}).
				AddRow(room1.ID, room1.Temperature, room1.Humidity, room1.CreatedAt),
		)

	// 実行
	err = repo.Insert(ctx, room1)
	if err != nil {
		t.Fatal(err)
	}
}
