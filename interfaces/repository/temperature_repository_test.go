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

	repo := TemperatureRepository{
		Database: db,
	}

	nowTime := time.Now()

	temperature1 := &domain.Temperature{
		ID:        3,
		Temp:      "22.5",
		Humi:      "76",
		CreatedAt: nowTime,
	}

	temperature2 := &domain.Temperature{
		ID:        4,
		Temp:      "11.8",
		Humi:      "60",
		CreatedAt: nowTime,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `temperatures` ORDER BY created_at desc LIMIT 200")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "temp", "humi", "created_at"}).
			AddRow(temperature1.ID, temperature1.Temp, temperature1.Humi, temperature1.CreatedAt).
			AddRow(4, "11.8", "60", nowTime),
		)

	res, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &res[0], temperature1)
	assert.Equal(t, &res[1], temperature2)

}

func TemperatureInsert(t *testing.T) {
	ctx := context.Background()
	db, mock, err := databases.MySQLMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	repo := TemperatureRepository{
		Database: db,
	}

	nowTime := time.Now()

	temperature1 := &domain.Temperature{
		ID:        3,
		Temp:      "22.5",
		Humi:      "76",
		CreatedAt: nowTime,
	}

	insertTemperature := regexp.QuoteMeta(`INSERT INTO "temperatures" ("id","temp", "humi", "created_at") VALUES ($1, $2, $3, $4)`)
	mock.ExpectQuery(insertTemperature).
		WithArgs(temperature1.ID, temperature1.Temp, temperature1.Humi, temperature1.CreatedAt).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "temp", "humi", "created_at"}).
				AddRow(temperature1.ID, temperature1.Temp, temperature1.Humi, temperature1.CreatedAt),
		)

	// 実行
	err = repo.Insert(ctx, temperature1)
	if err != nil {
		t.Fatal(err)
	}
}
