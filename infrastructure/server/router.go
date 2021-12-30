package server

import (
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	"homeapi/applications/logging"
	"homeapi/interfaces/api"
)

func Run(
	db *gorm.DB,
	redisClient *redis.Client,
	twitterClient *twitter.Client,
	logging logging.Logging,
	validate *validator.Validate,
) {
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://www.seldnext.com", "http://www.seldnext.com"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	{
		// firestoreデータ登録参照
		firestoreController := api.NewFirestoreController(logging, validate)
		v1.GET("/firestores", firestoreController.List)
		v1.POST("/firestores", firestoreController.Create)

		// 自宅の温度
		thermometerController := api.NewThermometerController(db, logging, validate)
		v1.GET("/thermometers", thermometerController.List)
		v1.POST("/thermometers", thermometerController.Create)

		csvupdownController := api.NewCsvController(db, logging, validate)
		v1.GET("/csv_updown", csvupdownController.Download)

		imagesController := api.NewImagesController(db, logging, validate)
		v1.POST("/images", imagesController.Upload)

		twitterController := api.NewTwitterController(db, redisClient, twitterClient, logging, validate)
		v1.GET("/twitters", twitterController.Get)
		v1.POST("/twitters", twitterController.Create)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1323"
	}

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + port))
}
