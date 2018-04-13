package main

import (
	"demo/config"
	"demo/router"

	"demo/trouter"

	"demo/ticker"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.GetENV() == "test" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	if config.GetENV() == "test" {
		trouter.Run(r)
	} else {
		router.Run(r)
	}

	go ticker.RunSyncCoin()

	r.Run(":" + config.GetPORT())

}
