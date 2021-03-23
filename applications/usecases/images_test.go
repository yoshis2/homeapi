package usecases

import (
	repositorymock "homeapi/applications/repository/mock"
	"homeapi/applications/util"
	"homeapi/domain"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
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
		nowTime, err := util.JapaneseNowTime()
		if err != nil {
			t.Error(err)
		}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		_, mocks := newServerImageMocks(ctrl)
		image := &domain.Images{
			ImageName: "マーティン",
			ImagePath: "/var/www/html",
			CreatedAt: nowTime,
		}

		log.Printf(" : %v : %v", mocks, image)
		mocks.imageRepository.EXPECT().Insert(image).AnyTimes().Return(nil)

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
