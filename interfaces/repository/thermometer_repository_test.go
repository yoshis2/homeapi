package repository

import (
	"context"
	"homeapi/domain"
	"regexp"
	"testing"
	"time"

	"homeapi/infrastructure/databases"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureList(t *testing.T) {
	ctx := context.Background()
	db, mock, err := databases.MySQLMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	repo := ThermometerRepository{
		Database: db,
	}

	nowTime := time.Now()

	room1 := &domain.Temperature{
		ID:          3,
		Temperature: "22.5",
		Humidity:    "76",
		CreatedAt:   nowTime,
	}

	room2 := &domain.Temperature{
		ID:          4,
		Temperature: "11.8",
		Humidity:    "60",
		CreatedAt:   nowTime,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `temperatures` ORDER BY created_at desc LIMIT 200")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "temp", "humi", "created_at"}).
			AddRow(room1.ID, room1.Temperature, room1.Humidity, room1.CreatedAt).
			AddRow(4, "11.8", "60", nowTime),
		)

	res, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &res[0], room1)
	assert.Equal(t, &res[1], room2)

}

func TemperatureInsert(t *testing.T) {
	ctx := context.Background()
	db, mock, err := databases.MySQLMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	repo := ThermometerRepository{
		Database: db,
	}

	nowTime := time.Now()

	room1 := &domain.Temperature{
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
