package main

import (
	"log"

	"github.com/nakabonne/golintui/pkg/app"
	"github.com/nakabonne/golintui/pkg/config"
)

func main() {
	// TODO: Populate
	appConfig, err := config.New("", "", "", "", "", "", "", false)
	if err != nil {
		log.Fatal(err.Error())
	}
	a, err := app.New(appConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
