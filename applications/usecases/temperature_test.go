package usecases

import (
	"homeapi/applications/ports"
	repositorymock "homeapi/applications/repository/mock"
	"homeapi/applications/util"
	"homeapi/domain"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type serverMocks struct {
	temperatureRepository *repositorymock.MockTemperatureRepository
}

func newMocks(ctrl *gomock.Controller) (*TemperatureUsecase, *serverMocks) {
	mocks := &serverMocks{
		temperatureRepository: repositorymock.NewMockTemperatureRepository(ctrl),
	}

	temperatureUsecase := &TemperatureUsecase{
		TemperatureRepository: mocks.temperatureRepository,
	}

	return temperatureUsecase, mocks
}

func TestList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		nowTime := time.Now()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		temperatureUsecase, mocks := newMocks(ctrl)
		temperature := []domain.Temperature{
			{
				ID:        12,
				Temp:      "22",
				Humi:      "61",
				CreatedAt: nowTime,
			},
			{
				ID:        13,
				Temp:      "25",
				Humi:      "63",
				CreatedAt: nowTime,
			},
		}

		mocks.temperatureRepository.EXPECT().List().Return(temperature, nil)
		got, err := temperatureUsecase.List()
		require.NoError(t, err)
		if err != nil {
			t.Errorf("error message : %v", err)
		}
		want := []ports.TemperatureOutputPort{
			{
				ID:   12,
				Temp: "22",
				Humi: "61",
			},
			{
				ID:   13,
				Temp: "25",
				Humi: "63",
			},
		}
		assert.Equal(t, &want, got)
	})
}

func TestInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		nowTime, err := util.JapaneseNowTime()
		if err != nil {
			t.Error(err)
		}
		// nowTime := time.Now()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		temperatureUsecase, mocks := newMocks(ctrl)

		temperature := &domain.Temperature{
			Temp:      "20",
			Humi:      "55",
			CreatedAt: nowTime,
		}

		dofunc := func(temperature *domain.Temperature) *domain.Temperature {
			temperature.ID = 71
			temperature.CreatedAt = nowTime
			return temperature
		}

		mocks.temperatureRepository.EXPECT().Insert(temperature).Do(dofunc).Return(nil)

		request := &ports.TemperatureInputPort{
			Temp: "20",
			Humi: "55",
		}

		got, err := temperatureUsecase.Create(request)
		require.NoError(t, err)

		want := ports.TemperatureOutputPort{
			ID:   71,
			Temp: "20",
			Humi: "55",
		}
		assert.Equal(t, want, got)
	})
}
