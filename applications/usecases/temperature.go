package usecases

import (
	"context"
	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"

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
