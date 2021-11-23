package api

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"homeapi/applications/logging"
	"homeapi/applications/ports"
	"homeapi/applications/usecases"
	"homeapi/interfaces"
	"homeapi/interfaces/repository"
)

type TwitterController struct {
	Usecase *usecases.TwitterUsecase
}

func NewTwitterController(db *gorm.DB, redisClient *redis.Client, twitterClient *twitter.Client, logging logging.Logging, validate *validator.Validate) *TwitterController {
	return &TwitterController{
		Usecase: &usecases.TwitterUsecase{
			TwitterRepository: &repository.TwitterRepository{},
			DB:                db,
			RedisClient:       redisClient,
			TwitterClient:     twitterClient,
			Logging:           logging,
			Validator:         validate,
		},
	}
}

// Get はDBにTwitterにDBに登録されているデータをツイートさせる機能
// @Tags Twitter
// Twtter godoc
// @Summary Twitterの自分のアカウントにツイートさせる
// @Description 欲しいタイミングで過去の温度を出力し、グラフにできるようにする
// @Accept  json
// @Produce  json
// @Success 200 {string} ok
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /twitters [get]
func (controller *TwitterController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	if err := controller.Usecase.Get(ctx); err != nil {
		return c.JSON(interfaces.ErrorResponse(err))
	}
	return c.JSON(http.StatusOK, "ok")
}

// Create はDBにツイートさせる情報を登録させる
// @Tags Twitter
// Twtter godoc
// @Summary 家の温度と湿度のデータをデータベースから抽出する
// @Description 欲しいタイミングで過去の温度を出力し、グラフにできるようにする
// @Accept  json
// @Produce  json
// @Param twitter body ports.TwitterInputPort true "twitter"
// @Success 200 {object} ports.TwiterOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /twitters [post]
func (controller *TwitterController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.TwitterInputPort

	if err := c.Bind(&input); err != nil {
		c.JSON(interfaces.ErrorResponse(err))
	}

	if err := controller.Usecase.Validator.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	output, err := controller.Usecase.Create(ctx, &input)
	if err != nil {
		c.JSON(interfaces.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, output)
}
