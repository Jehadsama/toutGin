package main

import (
	"github.com/gin-gonic/gin"
	"toutGin/app/config"
	"toutGin/app/config/route"
)

func main() {
	gin.SetMode(gin.DebugMode)
	new(config.Database).GetConnect()
	defer config.DB.Close()
	engine := gin.New()
	(&route.Route{Engine: engine}).Run()
	engine.Run(":8080")
}
