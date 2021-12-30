package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/usecases"
	"homeapi/interfaces"
	"homeapi/interfaces/repository"
)

// ImageController は画像アップロードコントローラー
type ImageController struct {
	Usecase *usecases.ImagesUsecase
}

// NewImageController は画像アップロードコントローラー
func NewImageController(database *gorm.DB, logging logging.Logging, validate *validator.Validate) *ImageController {
	repository := &repository.ImageRepository{
		Database: database,
	}
	return &ImageController{
		Usecase: &usecases.ImagesUsecase{
			ImageRepository: repository,
			Database:        database,
			Logging:         logging,
			Validator:       validate,
		},
	}
}

// Upload は画像データをGorutineを使って複数並列アップロードするハンドラー
// @Tags images Gorutine
// Image godoc
// @Summary 画像データをGorutineを使って複数並列アップロードする
// @Description 画像データをGorutineを使って複数並列アップロードする
// @Accept  json
// @Produce  json
// @Param image body ports.ImageInputPort true "画像情報"
// @Success 200 {object} ports.ImageOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /images [post]
func (controller *ImageController) Upload(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.ImageInputPort
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.ErrorResponseObject{
			Message: err.Error(),
		})
	}
	if err := controller.Usecase.Validator.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	output, err := controller.Usecase.Upload(ctx, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, output)
}
