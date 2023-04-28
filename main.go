package main

import (
	"Jishu/config"
	"Jishu/middleware"
	"Jishu/model"
	"Jishu/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Config.AppMode)
	r := gin.Default()
	model.InitModel()
	config.InitSession(r)

	r.Use(middleware.Error)
	router.InitRouter(r)

	r.Run(":8080")
}
