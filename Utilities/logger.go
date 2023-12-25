package Utilities

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func (logger *Logger) Log(message string) {
	if logger == nil {
		return
	}
	logger.logger.Println(message)
}

func CreateLogger(logFilePath string) *Logger {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	var logger *log.Logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	return &Logger{logger}
}
