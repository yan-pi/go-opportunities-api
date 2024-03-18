package main

import (
	"github.com/yan-pi/go-opportunities-api/config"
	"github.com/yan-pi/go-opportunities-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	router.Initialize()
	println("Server is running on port 8080...")
}
