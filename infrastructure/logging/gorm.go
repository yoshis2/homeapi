package logging

import "github.com/yoshis2/homeapi/applications/logging"

// GormLogger Gormのloggerをラップする
type GormLogger struct {
	Logging logging.Logging
}

// Print Gormのログをloggerに送る
func (gl *GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		//TODO もう少し意味のあるものにしたいところ
		gl.Logging.SQLLog(v[1], v[2], v[3])
		return
	}
}
