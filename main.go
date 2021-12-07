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
	if err := godotenv.Load(fmt.Sprintf("infrastructure/config/%s.env", os.Getenv("GO_ENV"))); err != nil {
		log.Fatal("Error loading .env file")
	}
	validate := validator.New()
	newLogging := logging.NewLogrusLogging() // Logrus
	//newLogging := logging.NewStackdriverLogging() // stack driver
	defer newLogging.Close()

	swaggerSet()

	var newDB databases.DatabaseInterface
	var newRedis databases.RedisInterface
	var newTwitter ex_api.TwitterInterface

	newDB = databases.NewMysql() //mysql
	mysqlDb := newDB.Open()
	db, err := mysqlDb.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db.LogMode(true)
	// db.SetLogger(&logging.GormLogger{Logging: newLogging})

	newRedis = databases.NewRedis()
	redisClient := newRedis.Open()
	defer redisClient.Close()

	newTwitter = ex_api.NewTwitter()
	twitterClient := newTwitter.Open()

	server.Run(mysqlDb, redisClient, twitterClient, newLogging, validate)
}

func swaggerSet() {
	docs.SwaggerInfo.Title = "Home APIのSwagger"
	docs.SwaggerInfo.Description = "家の中の温度や監視カメラの状態を取得するAPI"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SWAGGER_SCHEMA")}
}
