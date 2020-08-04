package api

import (
	"encoding/csv"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"homeapi/applications/logging"
	"homeapi/applications/usecases"
	"homeapi/interfaces"
	"homeapi/interfaces/repository"
)

// CsvUpdownController はusecaseのコントローラー
type CsvUpdownController struct {
	Usecase *usecases.CsvUpdownUsecase
}

// NewCsvController はnewコントローラー
func NewCsvController(database *gorm.DB, logging logging.Logging) *CsvUpdownController {
	return &CsvUpdownController{
		Usecase: &usecases.CsvUpdownUsecase{
			TemperatureRepository: &repository.TemperatureRepository{},
			Database:              database,
			Logging:               logging,
		},
	}
}

// Download はDBの温度湿度データをCSVで出力するハンドラー
// @Tags CSV 自宅の気温
// Temperature godoc
// @Summary 家の温度と湿度のデータをデータベースからCSVで抽出する
// @Description 欲しいタイミングで過去の温度を出力し、グラフにできるようにする
// @Accept  json
// @Produce  json
// @Success 200 {string} ok
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /csv_updown [get]
func (controller *CsvUpdownController) Download(c echo.Context) error {
	generateTemperatures, err := controller.Usecase.Download()
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	response := c.Response()
	header := response.Header()
	header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	header.Set(echo.HeaderContentDisposition, "attachment; filename="+"temperature"+".csv")
	header.Set("Content-Transfer-Encoding", "binary")
	header.Set("Expires", "0")
	response.WriteHeader(http.StatusOK)

	writerTemperature := csv.NewWriter(response)
	for _, generateTemperature := range generateTemperatures {
		writerTemperature.Write(generateTemperature)
	}
	writerTemperature.Flush()

	return c.JSON(http.StatusOK, "ok")
}
