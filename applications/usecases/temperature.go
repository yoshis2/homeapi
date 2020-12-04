package usecases

import (
	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"

	"github.com/go-playground/validator/v10"
)

// TemperatureUsecase 気温のUsecase
type TemperatureUsecase struct {
	TemperatureRepository repository.TemperatureRepository
	Logging               logging.Logging
	Validator             *validator.Validate
}

func (usecase *TemperatureUsecase) List() (*[]ports.TemperatureOutputPort, error) {
	temperatures, err := usecase.TemperatureRepository.List()
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
	temperature.CreatedAt, err = util.JapaneseNowTime()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	if err := usecase.TemperatureRepository.Insert(temperature); err != nil {
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
