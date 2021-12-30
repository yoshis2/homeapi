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

// ThermometerController 気温制御 Controller
type ThermometerController struct {
	Usecase *usecases.ThermometerUsecase
}

// NewThermometerController Create New Temperature Controller
func NewThermometerController(db *gorm.DB, logging logging.Logging, validate *validator.Validate) *ThermometerController {
	repository := &repository.ThermometerRepository{
		Database: db,
	}
	return &ThermometerController{
		Usecase: &usecases.ThermometerUsecase{
			ThermometerRepository: repository,
			Logging:               logging,
			Validator:             validate,
		},
	}
}

// List はDBに温度湿度データを配列で出力するハンドラー
// @Tags 自宅の気温
// Thermometer godoc
// @Summary 家の温度と湿度のデータをデータベースから抽出する
// @Description 欲しいタイミングで過去の温度を出力し、グラフにできるようにする
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.ThermometerOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /thermometers [get]
func (controller *ThermometerController) List(c echo.Context) error {
	ctx := c.Request().Context()
	output, err := controller.Usecase.List(ctx)
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}
	return c.JSON(http.StatusOK, output)
}

// Create はDBに温度湿度データを登録するハンドラー
// @Tags 自宅の気温
// Thermometer godoc
// @Summary 家の温度と湿度のデータをデータベースに格納する
// @Description １時間ごとに家の温度と湿度をデータベースに格納する
// @Accept  json
// @Produce  json
// @Param thermometer body ports.ThermometerInputPort true "温度湿度情報"
// @Success 200 {object} ports.ThermometerOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /thermometers [post]
func (controller *ThermometerController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.ThermometerInputPort

	if err := c.Bind(&input); err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	if err := controller.Usecase.Validator.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	output, err := controller.Usecase.Create(ctx, &input)
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, output)
}
