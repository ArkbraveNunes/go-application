package main

import (
	config "application-openings/config"
	router "application-openings/router"
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
