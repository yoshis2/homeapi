package api

import (
	"net/http"

	"homeapi/infrastructure/firebases"
	"homeapi/interfaces"
	"homeapi/interfaces/repository"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/usecases"

	"github.com/labstack/echo/v4"
)

// FirestoreConnectController はFirestoreコネクト用コントローラー
type FirestoreConnectController struct {
	Usecase *usecases.FirestoreConnectUsecase
}

// NewFirestoreController はfirestoreコネクト用Newコントローラー
func NewFirestoreController(logging logging.Logging) *FirestoreConnectController {
	return &FirestoreConnectController{
		Usecase: &usecases.FirestoreConnectUsecase{
			FirestoreRepository: &repository.FirestoreRepository{},
			Firestore:           &firebases.Firestore{},
			Logging:             logging,
		},
	}
}

// List はfirestoreから登録したデータを取得する
// @Tags Firebase firestore
// Temperature godoc
// @Summary firebaseで登録したデータを取得する
// @Description firebaseで登録したデータを取得する
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.FirestoreConnectOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /firestores [get]
func (controller *FirestoreConnectController) List(c echo.Context) error {
	outputs, err := controller.Usecase.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, outputs)
}

// Create はDBにFirestoreにデータを登録するハンドラー
// @Tags Firebase firestore
// Temperature godoc
// @Summary firestoreの登録テスト用
// @Description FirebaseのFirestoreの登録接続をする
// @Accept  json
// @Produce  json
// @Param temperature body ports.FirestoreConnectInputPort true "温度湿度情報"
// @Success 200 {object} ports.FirestoreConnectOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /firestores [post]
func (controller *FirestoreConnectController) Create(c echo.Context) error {
	var input ports.FirestoreConnectInputPort

	if err := c.Bind(&input); err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	outputs, err := controller.Usecase.Create(&input)
	if err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, outputs)
}
