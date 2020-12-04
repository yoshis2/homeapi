package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/usecases"
	"homeapi/interfaces"
	"homeapi/interfaces/repository"
)

// TemperatureController 気温制御 Controller
type TemperatureController struct {
	Usecase *usecases.TemperatureUsecase
}

// NewTemperatureController Create New Temperature Controller
func NewTemperatureController(db *gorm.DB, logging logging.Logging, validate *validator.Validate) *TemperatureController {
	repository := &repository.TemperatureRepository{
		Database: db,
	}
	return &TemperatureController{
		Usecase: &usecases.TemperatureUsecase{
			TemperatureRepository: repository,
			Logging:               logging,
			Validator:             validate,
		},
	}
}

// List はDBに温度湿度データを配列で出力するハンドラー
// @Tags 自宅の気温
// Temperature godoc
// @Summary 家の温度と湿度のデータをデータベースから抽出する
// @Description 欲しいタイミングで過去の温度を出力し、グラフにできるようにする
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.TemperatureOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /temperatures [get]
func (controller *TemperatureController) List(c echo.Context) error {
	output, err := controller.Usecase.List()
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}
	return c.JSON(http.StatusOK, output)
}

// Create はDBに温度湿度データを登録するハンドラー
// @Tags 自宅の気温
// Temperature godoc
// @Summary 家の温度と湿度のデータをデータベースに格納する
// @Description １時間ごとに家の温度と湿度をデータベースに格納する
// @Accept  json
// @Produce  json
// @Param temperature body ports.TemperatureInputPort true "温度湿度情報"
// @Success 200 {object} ports.TemperatureOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /temperatures [post]
func (controller *TemperatureController) Create(c echo.Context) error {
	var input ports.TemperatureInputPort

	if err := c.Bind(&input); err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	if err := controller.Usecase.Validator.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	output, err := controller.Usecase.Create(&input)
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, output)
}
