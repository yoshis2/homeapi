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

	"github.com/go-playground/validator/v10"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

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
	Validator       *validator.Validate
}

// ImageSize リサイズする画像サイズ
type ImageSize struct {
	MaxWidth  uint
	MaxHeight uint
}

// ResizeImageList go2goのファイルアップロードで作成するサイズの配列
func ResizeImageList() []ImageSize {
	var imageSizes []ImageSize

	imageSizes = append(imageSizes, ImageSize{150, 150})
	imageSizes = append(imageSizes, ImageSize{300, 300})
	imageSizes = append(imageSizes, ImageSize{500, 500})
	imageSizes = append(imageSizes, ImageSize{1242, 1242})
	imageSizes = append(imageSizes, ImageSize{1920, 1080})

	return imageSizes
}

func (usecase *ImagesUsecase) Upload(ctx context.Context, input *ports.ImageInputPort) (*ports.ImageOutputPort, error) {
	now, err := util.JapaneseNowTime()
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	imageInfo := domain.Image{
		Name:      input.Name,
		Path:      input.Path,
		CreatedAt: now,
	}

	// ファイルタイプを取得
	imageType := GetImageType(input.Data)
	usecase.Logging.Info(fmt.Sprintf("イメージタイプ : %v", string(imageType)))

	// base64をデコードファイルデータに変換
	image, err := DecodeBase64Image(imageType, input.Data) //[]byte
	if err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	if err := applications.Transaction(usecase.Database, func(tx *gorm.DB) error {
		if err := usecase.ImageRepository.Insert(ctx, &imageInfo); err != nil {
			return err
		}

		// 画像をGorutineでアップロード
		if err := ImageUploads(imageType, imageInfo.Name, image); err != nil {
			return err
		}
		return nil
	}); err != nil {
		usecase.Logging.Error(err)
		return nil, err
	}

	output := &ports.ImageOutputPort{
		Name:      imageInfo.Name,
		Path:      imageInfo.Path,
		CreatedAt: imageInfo.CreatedAt,
	}
	return output, nil
}

// DecodeBase64ImageStr Base64でエンコードされたImageの文字列からcontentTypeを抜いてDecodeする
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
	_, cancel := context.WithCancel(ctx)
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
