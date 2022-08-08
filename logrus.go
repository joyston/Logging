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
	log.SetOutput(file)
	log.Warning("Just a warning")

	log.Trace("Something very low level.")
	// log.SetLevel(log.DebugLevel)
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")

	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

}
