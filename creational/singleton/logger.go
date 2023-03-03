package singleton

import (
	"fmt"
	"sync"
)


type MyLogger struct {
	logLevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.logLevel = level
}

var logger *MyLogger

var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		logger = &MyLogger{}
	})
	
	return logger
}