package main

import (
	"fmt"
	"log"
	"os"

	"homeapi/docs"
	"homeapi/infrastructure/databases"
	"homeapi/infrastructure/ex_api"
	"homeapi/infrastructure/logging"
	"homeapi/infrastructure/server"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// @contact.name スリーネクスト　サポート
// @contact.url https://www.threenext.com
// @contact.email seki@threenext.com
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load(fmt.Sprintf("infrastructure/config/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	newLogging := logging.NewLogrusLogging() // Logrus
	//newLogging := logging.NewStackdriverLogging() // stack driver
	defer newLogging.Close()

	docs.SwaggerInfo.Title = "Home APIのSwagger"
	docs.SwaggerInfo.Description = "家の中の温度や監視カメラの状態を取得するAPI"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SWAGGER_SCHEMA")}

	//db
	var newDB databases.DatabaseInterface
	newDB = databases.NewMysql() //mysql
	db := newDB.Open()
	defer db.Close()

	var newRedis databases.RedisInterface
	newRedis = databases.NewRedis()
	redisClient := newRedis.Open()

	var newTwitter ex_api.TwitterInterface
	newTwitter = ex_api.NewTwitter()
	twitterClient := newTwitter.Open()

	validate := validator.New()

	//TODO dbの設定系の調整
	db.LogMode(true)
	db.SetLogger(&logging.GormLogger{Logging: newLogging})

	server.Run(db, redisClient, twitterClient, newLogging, validate)
}
