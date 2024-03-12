package main

import (
	"github.com/Yan-pi/go-opportunities-api/config"
	"github.com/Yan-pi/go-opportunities-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	
	router.Initialize()
	println("Server is running on port 8080...")
}