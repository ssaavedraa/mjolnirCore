package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	infoLogger *log.Logger
	errorLogger *log.Logger
)

func init () {
	infoLogger = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR]: ", log.Ldate|log.Ltime)
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2) // Adjust the depth as needed
	if !ok {
			return "unknown file:0"
	}
	return fmt.Sprintf("%s line:%d", file, line)
}

func Info (message string) {
	infoLogger.Println(getCallerInfo(), message)
}

func Error (err error) {
	errorLogger.Println(getCallerInfo(), err)
}