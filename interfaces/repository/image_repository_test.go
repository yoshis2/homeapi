package repository

import (
	"context"
	"homeapi/domain"
	"homeapi/infrastructure/databases"
	"log"
	"regexp"
	"testing"

	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestImageList(t *testing.T) {
	ctx := context.Background()
	db, mock, err := databases.GormMock()
	if err != nil {
		t.Fatal(err)
	}

	repo := ImageRepository{
		Database: db,
	}

	strVal := "2024/8/10"
	layout := "2006/1/2"
	timeVal, err := time.Parse(layout, strVal)
	if err != nil {
		t.Error(err)
	}

	image1 := &domain.Image{
		ID:        20,
		Name:      "test image",
		Path:      "img/2024-08/cursol",
		CreatedAt: timeVal,
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `images` ORDER BY created_at desc LIMIT ?")).
		WithArgs(200).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "path", "created_at"}).
			AddRow(image1.ID, image1.Name, image1.Path, image1.CreatedAt),
		)

	res, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)

	assert.Equal(t, &res[0], image1)
}
