package logger

import (
	"fmt"
	"sync"
	"time"
)

const (
	INFO    string = "INFO"
	WARNING string = "WARNING"
	ERROR   string = "ERROR"
)

type Logger struct{}

var logger *Logger
var mtx = &sync.Mutex{}

func NewInstance() *Logger {
	if logger == nil {
		mtx.Lock()
		defer mtx.Unlock()
		if logger == nil {
			fmt.Println("Creating new logger")
			logger = &Logger{}
		}
	} else {
		fmt.Println("Logger already created")
	}
	return logger
}

func (l *Logger) Info(message string) {
	fmt.Printf("%s - %s: %s\n", getFormatedTimeNow(), INFO, message)
}

func (l *Logger) Warning(message string) {
	fmt.Printf("%s - %s: %s\n", getFormatedTimeNow(), WARNING, message)
}

func (l *Logger) Error(message string) {
	fmt.Printf("%s - %s: %s\n", getFormatedTimeNow(), ERROR, message)
}

func getFormatedTimeNow() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
