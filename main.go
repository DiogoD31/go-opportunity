package main

import (
	"github.com/DiogoD31/go-opportunity/config"
	"github.com/DiogoD31/go-opportunity/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
