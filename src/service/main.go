package main

import (
	"os"

	"service/router"
	"service/store"
)

func main() {
	store, err := store.Init()
	if err != nil {
		os.Exit(-1)
	}

	if err := router.Init("5050", store); err != nil {
		os.Exit(-1)
	}
}
