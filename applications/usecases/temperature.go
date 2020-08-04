package usecases

import (
	"log"
	"net/http"
	"reflect"

	"homeapi/applications"
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

func (usecase *TemperatureUsecase) List() (*[]ports.TemperatureOutputPort, *applications.UsecaseError) {
	temperatures, err := usecase.TemperatureRepository.List(usecase.DB)
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
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

func (usecase *TemperatureUsecase) Create(input *ports.TemperatureInputPort) (*ports.TemperatureOutputPort, *applications.UsecaseError) {
	temperature := &domain.Temperature{
		Temp: input.Temp,
		Humi: input.Humi,
	}

	if temperature.Temp == "" {
		usecaseError := &applications.UsecaseError{
			Code: http.StatusBadRequest,
			Msg:  "温度が入っていません",
		}
		return nil, usecaseError
	}

	if temperature.Humi == "" {
		usecaseError := &applications.UsecaseError{
			Code: http.StatusBadRequest,
			Msg:  "湿度の値が入っていません",
		}
		return nil, usecaseError
	}

	var err error
	// 時間フォーマット yyyy-mm-dd
	temperature.CreatedAt, err = util.JapaneseNowTime()
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	log.Printf("temperatureの方 : %v", reflect.TypeOf(temperature))
	err = usecase.TemperatureRepository.Insert(usecase.DB, temperature)
	if uerr := applications.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	output := &ports.TemperatureOutputPort{
		ID:   temperature.ID,
		Temp: temperature.Temp,
		Humi: temperature.Humi,
	}

	return output, nil
}
