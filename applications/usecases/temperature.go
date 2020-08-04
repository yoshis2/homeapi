package usecases

import (
	"fmt"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"

	"github.com/jinzhu/gorm"
)

// TemperatureUsecase 気温のUsecase
type TemperatureUsecase struct {
	TemperatureRepository repository.TemperatureRepository
	DB                    *gorm.DB
	Logging               logging.Logging
}

func (usecase *TemperatureUsecase) List() (*[]ports.TemperatureOutputPort, error) {
	temperatures, err := usecase.TemperatureRepository.List(usecase.DB)
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	outputs := make([]ports.TemperatureOutputPort, len(temperatures))
	for i, temperature := range temperatures {
		outputs[i] = ports.TemperatureOutputPort{
			ID:   temperature.ID,
			Temp: temperature.Temp,
			Humi: temperature.Humi,
		}
	}

	return &outputs, nil
}

func (usecase *TemperatureUsecase) Create(input *ports.TemperatureInputPort) (*ports.TemperatureOutputPort, error) {
	temperature := &domain.Temperature{
		Temp: input.Temp,
		Humi: input.Humi,
	}

	var err error
	if temperature.Temp == "" {
		return nil, fmt.Errorf("BadRequest 温度が入っていません")
	}

	if temperature.Humi == "" {
		return nil, fmt.Errorf("BadRequest 湿度の値が入っていません")
	}

	// 時間フォーマット yyyy-mm-dd
	temperature.CreatedAt, err = util.JapaneseNowTime()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	if err = usecase.TemperatureRepository.Insert(usecase.DB, temperature); err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	output := &ports.TemperatureOutputPort{
		ID:   temperature.ID,
		Temp: temperature.Temp,
		Humi: temperature.Humi,
	}

	return output, nil
}
