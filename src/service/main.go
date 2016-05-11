package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"service/router"
	"service/store"
)

var (
	level = os.Getenv("LEVEL")
)

func main() {

	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.Panic("envar $LEVEL was not defined.")
	}

	logrus.Info("Starting store...")
	store, err := store.Init()
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

	logrus.Info("Starting router...")
	if err := router.Init("5050", store); err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
