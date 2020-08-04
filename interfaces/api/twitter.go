package api

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/yoshis2/homeapi/applications/logging"
	"github.com/yoshis2/homeapi/applications/ports"
	"github.com/yoshis2/homeapi/applications/usecases"
	"github.com/yoshis2/homeapi/interfaces"
	"github.com/yoshis2/homeapi/interfaces/repository"
)

type TwitterController struct {
	Usecase *usecases.TwitterUsecase
}

func NewTwitterController(db *gorm.DB, redisClient *redis.Client, twitterClient *twitter.Client, logging logging.Logging) *TwitterController {
	return &TwitterController{
		Usecase: &usecases.TwitterUsecase{
			TwitterRepository: &repository.TwitterRepository{},
			DB:                db,
			RedisClient:       redisClient,
			TwitterClient:     twitterClient,
			Logging:           logging,
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
	if err := controller.Usecase.Get(); err != nil {
		return c.JSON(interfaces.GetErrorResponse(err))
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
// @Param twitter body ports.TwitterInputPort true "twitter""
// @Success 200 {object} ports.TwiterOutputPort
// @Failure 400 {object} interfaces.ErrorResponseObject
// @Failure 404 {object} interfaces.ErrorResponseObject
// @Failure 500 {object} interfaces.ErrorResponseObject
// @Router /twitters [post]
func (controller *TwitterController) Create(c echo.Context) error {
	var input ports.TwitterInputPort

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
