package logging

import (
	"log"
	"os"
	"time"
)

func Log(message string, err error) {
	logFile, logErr := os.OpenFile("logging/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if logErr != nil {
		log.Fatal(logErr)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	errorTime := time.Now().Format("2006-01-02 15:04:05")
	errorType := "Error"

	if err != nil {
		logger.Printf("[%s][%s] %v\n", errorTime, errorType, message, err)
	} else {
		logger.Printf("[%s][%s] Unknown error occurred\n", errorTime, message, errorType)
	}
}
