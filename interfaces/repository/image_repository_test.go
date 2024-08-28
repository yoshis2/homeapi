package repository

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	"homeapi/domain"
	"homeapi/infrastructure/databases"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestImageList(t *testing.T) {
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

	db, mock, err := databases.GormMock()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `images` ORDER BY created_at desc LIMIT ?")).
		WithArgs(200).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "path", "created_at"}).
			AddRow(image1.ID, image1.Name, image1.Path, image1.CreatedAt),
		)

	ctx := context.Background()
	repo := ImageRepository{
		Database: db,
	}

	res, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)

	assert.Equal(t, &res[0], image1)
}

func TestImageInsert(t *testing.T) {
	strVal := "2024/8/10"
	layout := "2006/1/2"
	timeVal, err := time.Parse(layout, strVal)
	if err != nil {
		t.Fatal(err)
	}

	image2 := &domain.Image{
		ID:        1,
		Name:      "test no image",
		Path:      "img/2024-09/cursol2",
		CreatedAt: timeVal,
	}

	// _, mock, err := databases.GormMock()
	db, mock, err := databases.GormMock()
	if err != nil {
		t.Fatal(err)
	}

	insertImage := regexp.QuoteMeta("INSERT INTO `images` (`name`,`path`,`created_at`,`id`) VALUES (?,?,?,?)")
	mock.ExpectBegin()
	mock.ExpectExec(insertImage).
		WithArgs(image2.Name, image2.Path, image2.CreatedAt, 1).
		WillReturnResult(
			sqlmock.NewResult(1, 100),
		)
	mock.ExpectCommit()

	ctx := context.Background()
	repo := ImageRepository{
		Database: db,
	}

	if err := repo.Insert(ctx, image2); err != nil {
		t.Fatal(err)
	}
}
