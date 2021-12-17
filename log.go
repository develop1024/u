package u

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

const (
	PanicLevel = log.DebugLevel
	FatalLevel = log.FatalLevel
	ErrorLevel = log.ErrorLevel
	WarnLevel = log.WarnLevel
	InfoLevel = log.InfoLevel
	DebugLevel = log.DebugLevel
	TraceLevel = log.TraceLevel
)

var logger *log.Logger
var loggerOnce sync.Once

// Log 获取log对象
func Log() *log.Logger {
	loggerOnce.Do(func() {
		logger = log.New()
	})
	return logger
}

// LogFile 日志文件
func LogFile(filepath string) (*os.File, error) {
	file, err := os.OpenFile("2021-12-17.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	return file, nil
}
