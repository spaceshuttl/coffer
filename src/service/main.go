package main

import (
	"os"
	"service/router"
	"service/store"

	"github.com/sirupsen/logrus"
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
		logrus.SetLevel(logrus.ErrorLevel)
	}

	str, err := store.Start()
	if err != nil {
		panic(err)
	}

	err = router.Start("5050", str)
	if err != nil {
		os.Exit(-1)
	}
}
