package main

import (
	"github.com/mnzt/coffer/internal/router"
	"github.com/mnzt/coffer/internal/store"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	level = os.Getenv("LOG_LEVEL")
)

func main() {

	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	// TODO: Implement a new store

	err = router.Start("5050", dataStore)
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

}
