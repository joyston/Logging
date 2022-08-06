package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func ExecuteLogrus() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.Warning("Just a warning")
	log.SetOutput(file)
}
