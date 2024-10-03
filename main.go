package main

import (
	"github.com/marciopriebe/crud-go.git/config"
	"github.com/marciopriebe/crud-go.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization: %v", err)
		return
	}

	// initiaizer router
	router.Initializer()

}
