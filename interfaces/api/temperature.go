package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/yoshis2/homeapi/applications/logging"
	"github.com/yoshis2/homeapi/applications/ports"
	"github.com/yoshis2/homeapi/applications/usecases"
	"github.com/yoshis2/homeapi/interfaces"
	"github.com/yoshis2/homeapi/interfaces/repository"
)

// TemperatureController 気温制御 Controller
type TemperatureController struct {
	Usecase *usecases.TemperatureUsecase
}

// NewTemperatureController Create New Temperature Controller
func NewTemperatureController(db *gorm.DB, logging logging.Logging) *TemperatureController {
	return &TemperatureController{
		Usecase: &usecases.TemperatureUsecase{
			TemperatureRepository: &repository.TemperatureRepository{},
			DB:                    db,
			Logging:               logging,
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
		return c.JSON(interfaces.GetErrorResponse(err))
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
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponseObject{
			Message: err.Error(),
		})
	}
	output, err := controller.Usecase.Create(&input)
	if err != nil {
		c.JSON(interfaces.GetErrorResponse(err))
	}

	return c.JSON(http.StatusOK, output)
}
