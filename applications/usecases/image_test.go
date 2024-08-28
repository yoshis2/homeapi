package usecases

import (
	"context"
	repositorymock "homeapi/applications/repository/mock"
	"homeapi/applications/util"
	"homeapi/domain"
	"homeapi/infrastructure/databases"
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

	db, _, _ := databases.GormMock()
	imagesUsecase := &ImagesUsecase{
		ImageRepository: mocks.imageRepository,
		Database:        db,
	}

	return imagesUsecase, mocks
}

func TestImageList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
	})
}

func TestImageInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		nowTime, err := util.JapaneseNowTime()
		if err != nil {
			t.Error(err)
		}

		image := &domain.Image{
			Name:      "マーティン",
			Path:      "/var/www/html",
			CreatedAt: nowTime,
		}

		_, mocks := newServerImageMocks(ctrl)
		// imagesUsecase, mocks := newServerImageMocks(ctrl)
		mocks.imageRepository.EXPECT().Insert(ctx, image).AnyTimes().Return(nil)

		// imageData := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAABhGlDQ1BJQ0MgcHJvZmlsZQAAKJF9kT1Iw0AcxV9TpVUqDmYoxSFDdbIgKuKoVShChVArtOpgPvoFTRqSFBdHwbXg4Mdi1cHFWVcHV0EQ/ABxdnBSdJES/5cUWsR4cNyPd/ced+8ArllVNKtnHNB028ykkkIuvyqEXhFGFDxi6JMUy5gTxTR8x9c9Amy9S7As/3N/jgG1YClAQCCeVQzTJt4gnt60Dcb7xLxSllTic+Ixky5I/Mh02eM3xiWXOZbJm9nMPDFPLJS6WO5ipWxqxFPEcVXTKZ/Leawy3mKsVetK+57shZGCvrLMdJrDSGERSxAhQEYdFVRhI0GrToqFDO0nffwx1y+SSyZXBQo5FlCDBsn1g/3B726t4uSElxRJAr0vjvMxAoR2gVbDcb6PHad1AgSfgSu94681gZlP0hsdLX4EDG4DF9cdTd4DLneA6JMhmZIrBWlyxSLwfkbflAeGboH+Na+39j5OH4AsdZW+AQ4OgdESZa/7vDvc3du/Z9r9/QBOCnKYu9amkgAAAAZiS0dEAP8A/wD/oL2nkwAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB+gIGBQeAFxwM40AAALeSURBVGje7dhvyF5zHMfx19yXyc3Gyp3UHlzdyijK/yf+/H788NCfxEQ0yUjcDcVmxYNp8r8VD4w9knUjSyYljs7ZGs2fJyNKi+2BtKYUQ7Nxe+A8uDqd6zZ26zqnzqeuJ99z/er37vP9dw6dOnXq1KltSjEsSTHEJtxl7HAOT/b7N+CVyX7/58l+f/u3u3aPDOSIwzx/JDbgOmxIMcxvKwjswaX4A0WK4cS2gsjyYn+WF7djGh+lGM5uY2r9OgC0DsvxVophadscmam48z4iVqcYHksxjLUmtWpSbScuxOnYlGJY2HSQHvYPgfkJV+JLfJhiOLnJIPNwcBZnZrK8WIW1yFMMqVWpVQO0EVfjxRTDVIphXtNAxnDgEGE+wwVYivUphqNaCVLCfF8Ozx6yFMNEq1KrZnjeitfxcYrhzCaAzJ+t2P8BaB3uxOYUw/WjBhkf1n4PEeZdJDycYliTYvjP9+mN8F1mHJM4A29jNRZgxShAetUVZZaLL8YUlpS/xfgEu/E5rsKno3Lk2H9RIwtwF67BDuzN8uLPuXK49z+lTa+sgYMD9fBViuEDTGR5saeJK0oVYgJb8WjN/x/HA02c7Avx+wDEadiGd7AsxXB8pUttwz5c0dhdK8VwMd7Dyiwv1uDVck7UubKykSAphpvwMq7N8mJTGX4Wd6cYjq64shknpBjObxLIcbgHq3BJlhfbBy68C1twS825J8ozjXJkHBdlefFNzbO1uL+muKdxVorh1KaALMflWV78OGQF+QJfl6v7YPwAnsZ9cwUy521wSBN4BudleTEzED8GO3FulhffNX6Nz/JiS9miL6vEf8ELuLdN7yNP4sGa+PO4McWwqC0gb+KkFMM5FVf24o0h86aRHx9m8BQeqtTPIvyG2xq5NA7RRjySYjjF359Zp3AzXqvWTyO7VsWBFbijXOlfwnNZXvzQ2DV+Fq0vl8bpLC/26dSpU6dOnRqivwBcOOiISmv7JQAAAABJRU5ErkJggg=="
		// request := ports.ImageInputPort{
		// 	Name: "マーティン",
		// 	Path: "/var/www/html",
		// 	Data: imageData,
		// }

		// got, err := imagesUsecase.Upload(ctx, &request)
		// require.NoError(t, err)

		// log.Println(got)

		// want := ports.ImageOutputPort{
		// 	Name:      "マーティン",
		// 	Path:      "/var/www/html",
		// 	CreatedAt: nowTime,
		// }
		// assert.Equal(t, want, got)
	})
}
