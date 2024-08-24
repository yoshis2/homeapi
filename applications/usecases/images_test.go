package usecases

import (
	repositorymock "homeapi/applications/repository/mock"
	"log"
	"testing"

	"go.uber.org/mock/gomock"
)

type serverImageMocks struct {
	imageRepository *repositorymock.MockImageRepository
}

func newServerImageMocks(ctrl *gomock.Controller) (*ImagesUsecase, *serverImageMocks) {
	mocks := &serverImageMocks{
		imageRepository: repositorymock.NewMockImageRepository(ctrl),
	}

	imagesUsecase := &ImagesUsecase{
		ImageRepository: mocks.imageRepository,
	}

	return imagesUsecase, mocks
}

func TestImageInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		_, mocks := newServerImageMocks(ctrl)
		log.Printf(" : %v ", mocks)

		// nowTime, err := util.JapaneseNowTime()
		// if err != nil {
		// 	t.Error(err)
		// }

		// image := &domain.Image{
		// 	Name:      "マーティン",
		// 	Path:      "/var/www/html",
		// 	CreatedAt: nowTime,
		// }

		// mocks.imageRepository.EXPECT().Insert(image).AnyTimes().Return(nil)

		// request := &ports.ImagesInputPort{
		// 	ImageName: "マーティン",
		// 	ImagePath: "/var/www/html",
		// 	ImageData: "ponpon",
		// }

		// got, err := imagesUsecase.Upload(request)
		// require.NoError(t, err)

		// want := ports.ImagesOutputPort{
		// 	ImageName: "マーティン",
		// 	ImagePath: "/var/www/html",
		// 	CreatedAt: nowTime,
		// }
		// assert.Equal(t, want, got)
	})
}
