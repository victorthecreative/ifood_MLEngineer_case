package utils

import (
	"log"
	"os"
)

func LogRequestResponse(request, response string) {
	file, err := os.OpenFile("api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Printf("Request: %s | Response: %s", request, response)
}
