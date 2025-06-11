package api

import (
	"homeapi/applications/logging"
	"homeapi/interfaces"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// FirestoreConnectController はFirestoreコネクト用コントローラー
type WalletController struct {
	// Usecase *usecases.FirestoreConnectUsecase
}

// NewFirestoreController はfirestoreコネクト用Newコントローラー
func NewWalletController(logging logging.Logging, validate *validator.Validate) *WalletController {
	return &WalletController{}
}

// List はDBにウォレット情報を配列で出力するハンドラー
// @Tags ウォレット情報
// Wallet godoc
// @Summary ウォレットの情報をデータベースから抽出する
// @Description 仮想通貨のウォレットの情報を出力し、グラフにできるようにする
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.WalletOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /wallets [get]
func (controller *WalletController) List(c echo.Context) error {
	ctx := c.Request().Context()
	output, err := controller.Usecase.List(ctx)
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}
	return c.JSON(http.StatusOK, output)
}

// Create はDBにウォレット情報を登録するハンドラー
// @Tags ウォレット情報
// Wallet godoc
// @Summary ウォレット情報をデータベースに格納する
// @Description ウォレット情報をデータベースに格納する
// @Accept  json
// @Produce  json
// @Param wallet body ports.WalletInputPort true "温度湿度情報"
// @Success 200 {object} ports.WalletOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /wallets [post]
func (controller *WalletController) Create(c echo.Context) error {
	// ctx := c.Request().Context()
	// var input ports.ThermometerInputPort

	// if err := c.Bind(&input); err != nil {
	// 	return c.JSON(interfaces.ErrorResponse(err))
	// }

	// if err := controller.Usecase.Validator.Struct(&input); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	// output, err := controller.Usecase.Create(ctx, &input)
	// if err != nil {
	// 	return c.JSON(interfaces.ErrorResponse(err))
	// }

	return c.JSON(http.StatusOK, nil)
}
