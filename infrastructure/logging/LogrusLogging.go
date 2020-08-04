package logging

import (
	"fmt"
	"os"

	"homeapi/applications/logging"

	"github.com/sirupsen/logrus"
)

//LogrusLogging STDOUT,STDERRにStackdriverLogging形式のログを出力し、StackdriverLoggingに自動的に取り込んでもらう場合のlogger
type LogrusLogging struct {
	Client *logrus.Logger
}

//NewLogrusLogging New LogrusLogging
func NewLogrusLogging() *LogrusLogging {
	logger := logrus.New()
	accessLogFile, err := os.OpenFile("required/logs/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}
	logger.Out = accessLogFile
	//logger.Out = os.Stderr

	return &LogrusLogging{
		Client: logger,
	}
}

//Close 特になにもしていない
func (logrus *LogrusLogging) Close() {
}

//Error Errorレベルのアプリケーションログの出力
func (logrus *LogrusLogging) Error(data interface{}) {
	logrus.Client.Errorln(data)
}

//Warning Warningレベルのアプリケーションログの出力
func (logrus *LogrusLogging) Warning(data interface{}) {
	logrus.Client.Warningln(data)
}

//Info Infoレベルのアプリケーションログの出力
func (logrus *LogrusLogging) Info(data interface{}) {
	logrus.Client.Infoln(data)
}

//Debug Debugレベルのアプリケーションログの出力
func (logrus *LogrusLogging) Debug(data interface{}) {
	logrus.Client.Debugln(data)
}

//AccessLog ginのアクセスログの出力
func (logrus *LogrusLogging) AccessLog(data *logging.AccessLogEntry) {
	logrus.Client.Infoln(data)
}

//SQLLog gormのSQLログの出力
func (logrus *LogrusLogging) SQLLog(v1 interface{}, v2 interface{}, v3 interface{}) {
	logrus.Client.Debugln(v1)
	logrus.Client.Debugln(v2)
	logrus.Client.Debugln(v3)
}
