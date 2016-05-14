package main

import (
	"os"
	"service/router"
	"service/store"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	str, err := store.Start()
	if err != nil {
		panic(err)
	}

	err = router.Start("5050", str)
	if err != nil {
		os.Exit(-1)
	}
}
