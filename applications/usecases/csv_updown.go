package usecases

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/yoshis2/homeapi/applications"
	"github.com/yoshis2/homeapi/applications/logging"
	"github.com/yoshis2/homeapi/applications/repository"
	"github.com/yoshis2/homeapi/domain"
)

type CsvUpdownUsecase struct {
	TemperatureRepository repository.TemperatureRepository
	Database              *gorm.DB
	Logging               logging.Logging
}

func (usecase *CsvUpdownUsecase) Download() ([][]string, *applications.UsecaseError) {
	temperatures, err := usecase.TemperatureRepository.List(usecase.Database)
	if usecaseError := applications.GetUErrorByError(err); usecaseError != nil {
		usecase.Logging.Error(err)
		return nil, usecaseError
	}

	var newTemperatures []domain.Temperature
	for _, temperature := range temperatures {
		newTemperatures = append(newTemperatures, temperature)
	}

	generateTemperatures := generateCSVRows(newTemperatures)
	return generateTemperatures, nil
}

// 構造体を配列の文字列に変更しCSVに変換できる状態にする。
func generateCSVRows(src interface{}) [][]string {
	slices := []interface{}{}
	if csvData := reflect.ValueOf(src); csvData.Kind() == reflect.Slice {
		log.Printf("vの正体 : %v", csvData.Kind())
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
