package logging

import (
	"context"
	"os"

	ulog "github.com/yoshis2/homeapi/applications/logging"
	"cloud.google.com/go/logging"
	"google.golang.org/api/option"
)

var applicationLogName string
var accessLogName string
var databaseLogName string

//StackdriverLogging StackdriverLoggingのAPIを叩いて、直接ログを送る場合のlogger
type StackdriverLogging struct {
	Client *logging.Client
}

func init() {
	applicationLogName = os.Getenv("APPLICATION_LOG")
	accessLogName = os.Getenv("ACCESS_LOG")
	databaseLogName = os.Getenv("DATABASE_LOG")
}

//NewStackdriverLogging New StackdriverLogging
func NewStackdriverLogging() *StackdriverLogging {
	opt := option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT_KEY"))
	projectID := os.Getenv("GCS_PROJECT_ID")

	client, err := logging.NewClient(context.Background(), projectID, opt)
	if err != nil {
		panic(err)
	}

	return &StackdriverLogging{
		Client: client,
	}
}

//Close Close処理
func (stlog *StackdriverLogging) Close() {
	if err := stlog.Client.Close(); err != nil {
		panic(err)
	}
}

//Error Errorレベルのアプリケーションログの出力
func (stlog *StackdriverLogging) Error(data interface{}) {
	logger := stlog.Client.Logger(applicationLogName).StandardLogger(logging.Error)
	logger.Println(data)
}

//Warning Warningレベルのアプリケーションログの出力
func (stlog *StackdriverLogging) Warning(data interface{}) {
	logger := stlog.Client.Logger(applicationLogName).StandardLogger(logging.Warning)
	logger.Println(data)
}

//Info Infoレベルのアプリケーションログの出力
func (stlog *StackdriverLogging) Info(data interface{}) {
	logger := stlog.Client.Logger(applicationLogName).StandardLogger(logging.Info)
	logger.Println(data)
}

//Debug Debugレベルのアプリケーションログの出力
func (stlog *StackdriverLogging) Debug(data interface{}) {
	logger := stlog.Client.Logger(applicationLogName).StandardLogger(logging.Debug)
	logger.Println(data)
}

//AccessLog ginのアクセスログの出力
func (stlog *StackdriverLogging) AccessLog(data *ulog.AccessLogEntry) {
	logger := stlog.Client.Logger(accessLogName).StandardLogger(logging.Notice)
	logger.Println(data)
}

//SQLLog gormのSQLログの出力
func (stlog *StackdriverLogging) SQLLog(v1 interface{}, v2 interface{}, v3 interface{}) {
	logger := stlog.Client.Logger(databaseLogName).StandardLogger(logging.Info)
	logger.Println(v1)
	logger.Println(v2)
	logger.Println(v3)
}
