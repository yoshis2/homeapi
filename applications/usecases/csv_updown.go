package usecases

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"homeapi/applications/logging"
	"homeapi/applications/repository"
	"homeapi/domain"

	"github.com/go-playground/validator/v10"
)

type CsvUpdownUsecase struct {
	TemperatureRepository repository.TemperatureRepository
	Logging               logging.Logging
	Validator             *validator.Validate
}

func (usecase *CsvUpdownUsecase) Download(ctx context.Context) ([][]string, error) {
	temperatures, err := usecase.TemperatureRepository.List(ctx)
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	var newTemperatures []domain.Temperature
	newTemperatures = append(newTemperatures, temperatures...)

	return generateCSVRows(newTemperatures), nil
}

// 構造体を配列の文字列に変更しCSVに変換できる状態にする。
func generateCSVRows(src interface{}) [][]string {
	slices := []interface{}{}
	if csvData := reflect.ValueOf(src); csvData.Kind() == reflect.Slice {
		for i := 0; i < csvData.Len(); i++ {
			slices = append(slices, csvData.Index(i).Interface())
		}
	} else {
		slices = append(slices, csvData.Interface())
	}

	rows := make([][]string, 1)
	ignoreColIndex := map[int]bool{}

	for number, d := range slices {
		rows = append(rows, []string{})
		v := reflect.ValueOf(d)
		for i := 0; i < v.NumField(); i++ {
			if number == 0 {
				colName := v.Type().Field(i).Tag.Get("csv")
				if colName == "" {
					colName = strings.ToLower(v.Type().Field(i).Name)
				} else if colName == "-" {
					ignoreColIndex[i] = true
					continue
				}
				rows[0] = append(rows[0], colName)
			}
			if !ignoreColIndex[i] {
				rows[len(rows)-1] = append(rows[len(rows)-1], fmt.Sprint(v.Field(i).Interface()))
			}
		}
	}
	return rows
}
