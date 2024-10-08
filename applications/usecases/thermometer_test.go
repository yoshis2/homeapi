package usecases

import (
	"context"
	"homeapi/applications/ports"
	repositorymock "homeapi/applications/repository/mock"
	"homeapi/applications/util"
	"homeapi/domain"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type serverMocks struct {
	temperatureRepository *repositorymock.MockThermometerRepository
}

func newMocks(ctrl *gomock.Controller) (*ThermometerUsecase, *serverMocks) {
	mocks := &serverMocks{
		temperatureRepository: repositorymock.NewMockThermometerRepository(ctrl),
	}

	temperatureUsecase := &ThermometerUsecase{
		ThermometerRepository: mocks.temperatureRepository,
	}

	return temperatureUsecase, mocks
}

func TestList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		nowTime := time.Now()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		temperatureUsecase, mocks := newMocks(ctrl)
		temperature := []domain.Thermometer{
			{
				ID:          12,
				Temperature: "22",
				Humidity:    "61",
				CreatedAt:   nowTime,
			},
			{
				ID:          13,
				Temperature: "25",
				Humidity:    "63",
				CreatedAt:   nowTime,
			},
		}

		mocks.temperatureRepository.EXPECT().List(ctx).Return(temperature, nil)
		got, err := temperatureUsecase.List(ctx)
		require.NoError(t, err)
		if err != nil {
			t.Errorf("error message : %v", err)
		}
		want := []ports.ThermometerOutputPort{
			{
				ID:          12,
				Temperature: "22",
				Humidity:    "61",
			},
			{
				ID:          13,
				Temperature: "25",
				Humidity:    "63",
			},
		}
		assert.Equal(t, &want, got)
	})
}

func TestInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		nowTime, err := util.JapaneseNowTime()
		if err != nil {
			t.Error(err)
		}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		temperatureUsecase, mocks := newMocks(ctrl)

		temperature := &domain.Thermometer{
			Temperature: "20",
			Humidity:    "55",
			CreatedAt:   nowTime,
		}

		dofunc := func(ctx context.Context, temperature *domain.Thermometer) (context.Context, *domain.Thermometer) {
			temperature.ID = 71
			temperature.Temperature = "20"
			temperature.Humidity = "55"
			return ctx, temperature
		}

		mocks.temperatureRepository.EXPECT().Insert(ctx, temperature).Do(dofunc).Return(nil)

		request := &ports.ThermometerInputPort{
			Temperature: "20",
			Humidity:    "55",
		}

		got, err := temperatureUsecase.Create(ctx, request)
		require.NoError(t, err)

		want := ports.ThermometerOutputPort{
			ID:          71,
			Temperature: "20",
			Humidity:    "55",
		}
		assert.Equal(t, want, got)
	})
}
