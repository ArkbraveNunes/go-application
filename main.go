package main

import (
	"application-openings/config"
	"application-openings/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.ErrorWithValues("Config init error: %v", err)
		return
	}

	router.InitRouter()
}
