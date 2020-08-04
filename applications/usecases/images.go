package usecases

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"homeapi/applications"
	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/repository"
	"homeapi/applications/util"
	"homeapi/domain"
)

type ImagesUsecase struct {
	ImageRepository repository.ImageRepository
	Database        *gorm.DB
	Logging         logging.Logging
}

// ImageSize リサイズする画像サイズ
type ImageSize struct {
	MaxWidth  uint
	MaxHeight uint
}

//ResizeImageList go2goのファイルアップロードで作成するサイズの配列
func ResizeImageList() []ImageSize {
	var imageSizes []ImageSize

	imageSizes = append(imageSizes, ImageSize{150, 150})
	imageSizes = append(imageSizes, ImageSize{300, 300})
	imageSizes = append(imageSizes, ImageSize{500, 500})
	imageSizes = append(imageSizes, ImageSize{1242, 1242})
	imageSizes = append(imageSizes, ImageSize{1920, 1080})

	return imageSizes
}

func (usecase *ImagesUsecase) Upload(input *ports.ImagesInputPort) (*ports.ImagesOutputPort, error) {
	images := domain.Images{
		ImageName: input.ImageName,
		ImagePath: input.ImagePath,
	}

	var err error
	images.CreatedAt, err = util.JapaneseNowTime()

	// ファイルタイプを取得
	imageType := GetImageType(input.ImageData)
	usecase.Logging.Info(fmt.Sprintf("イメージタイプ : %v", string(imageType)))

	// base64をデコードファイルデータに変換
	image, err := DecodeBase64Image(imageType, input.ImageData) //[]byte
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	if err := applications.Transaction(usecase.Database, func(tx *gorm.DB) error {
		if err := usecase.ImageRepository.Insert(usecase.Database, &images); err != nil {
			return err
		}

		// 画像をGorutineでアップロード
		if err := ImageUploads(imageType, images.ImageName, image); err != nil {
			return err
		}
		return nil
	}); err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	output := &ports.ImagesOutputPort{
		ImageName: images.ImageName,
		ImagePath: images.ImagePath,
		CreatedAt: images.CreatedAt,
	}
	return output, nil
}

//DecodeBase64ImageStr Base64でエンコードされたImageの文字列からcontentTypeを抜いてDecodeする
func DecodeBase64Image(imageType, base64ImageStr string) (image.Image, error) {
	coI := strings.Index(base64ImageStr, ",")
	rawImage := base64ImageStr[coI+1:]

	imageData, err := base64.StdEncoding.DecodeString(rawImage)
	if err != nil {
		return nil, err
	}
	//デコードしたbyteデータを、image.Imageに変換
	var img image.Image

	switch imageType {
	case "png":
		img, err = png.Decode(bytes.NewReader(imageData))
		if err != nil {
			return nil, err
		}
	case "jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageData))
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("指定のファイル種類以外のファイルです。")
	}
	return img, nil
}

func ImageUploads(imageType, imageName string, image image.Image) error {
	// 一つでもエラーが出ると全て戻す処理。
	errorGroup, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, size := range ResizeImageList() {
		size := size // エラーグループの中に値を正しく持っていくために必要
		errorGroup.Go(func() error {
			// 複数サイズにリサイズ(Go Routine)
			m := resize.Thumbnail(size.MaxWidth, size.MaxHeight, image, resize.Lanczos3)

			out, err := os.Create(fmt.Sprintf("img/%s_%d_%d.%s", imageName, size.MaxWidth, size.MaxHeight, imageType))
			if err != nil {
				return err
			}
			defer out.Close()
			// 画像をフォルダに格納する
			switch imageType {
			case "png":
				png.Encode(out, m)
			case "jpeg":
				jpeg.Encode(out, m, nil)
			default:
				return errors.New("指定のファイル種類以外のファイルです。")
			}

			return nil
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return err
	}

	return nil
}

// GetImageType Base64でエンコードされたImageの文字列のcontentTypeからimagetypeを取得
func GetImageType(base64ImageStr string) string {
	start := strings.Index(base64ImageStr, "/") + 1
	end := strings.Index(base64ImageStr, ";")

	return base64ImageStr[start:end]
}
