package usecases

import (
	"context"
	"fmt"
	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ThermometerUsecase 気温のUsecase
type ThermometerUsecase struct {
	ThermometerRepository repository.ThermometerRepository
	Logging               logging.Logging
	Validator             *validator.Validate
}

func (usecase *ThermometerUsecase) List(ctx context.Context) (*[]ports.ThermometerOutputPort, error) {
	temperatures, err := usecase.ThermometerRepository.List(ctx)
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	outputs := make([]ports.ThermometerOutputPort, len(temperatures))
	for i, temperature := range temperatures {
		outputs[i] = ports.ThermometerOutputPort{
			ID:          temperature.ID,
			Temperature: temperature.Temperature,
			Humidity:    temperature.Humidity,
		}
	}

	return &outputs, nil
}

func (usecase *ThermometerUsecase) Create(ctx context.Context, input *ports.ThermometerInputPort) (*ports.ThermometerOutputPort, error) {
	now, err := util.JapaneseNowTime()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	temperature := &domain.Thermometer{
		Temperature: input.Temperature,
		Humidity:    input.Humidity,
		CreatedAt:   now,
	}

	if err := usecase.ThermometerRepository.Insert(ctx, temperature); err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	output := &ports.ThermometerOutputPort{
		ID:          temperature.ID,
		Temperature: temperature.Temperature,
		Humidity:    temperature.Humidity,
	}

	return output, nil
}

func (usecase *ThermometerUsecase) Download(ctx context.Context) ([][]string, error) {
	temperatures, err := usecase.ThermometerRepository.List(ctx)
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	var newTemperatures []domain.Thermometer
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
