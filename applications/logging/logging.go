package logging

import (
	"time"
)

type AccessLogEntry struct {
	Status    int
	Method    string
	Path      string
	IP        string
	Latency   time.Duration
	UserAgent string
	Time      time.Time
}

//Logging Loggingのインターフェース
type Logging interface {
	Close()
	Error(interface{})
	Warning(interface{})
	Info(interface{})
	Debug(interface{})
	AccessLog(*AccessLogEntry)
	SQLLog(interface{}, interface{}, interface{})
}
