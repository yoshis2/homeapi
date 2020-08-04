package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/usecases"
	"homeapi/interfaces"
	"homeapi/interfaces/repository"
)

// ImagesController は画像アップロードコントローラー
type ImagesController struct {
	Usecase *usecases.ImagesUsecase
}

// NewImagesController は画像アップロードコントローラー
func NewImagesController(database *gorm.DB, logging logging.Logging) *ImagesController {
	return &ImagesController{
		Usecase: &usecases.ImagesUsecase{
			ImageRepository: &repository.ImageRepository{},
			Database:        database,
			Logging:         logging,
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
// @Param image body ports.ImagesInputPort true "画像情報"
// @Success 200 {object} ports.ImagesOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /images [post]
func (controller *ImagesController) Upload(c echo.Context) error {
	var input ports.ImagesInputPort
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, interfaces.ErrorResponseObject{
			Message: err.Error(),
		})
	}
	output, err := controller.Usecase.Upload(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, output)
}
