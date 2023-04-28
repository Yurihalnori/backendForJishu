package router

import (
	"Jishu/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// r.Use(middleware.Error)
	apiRouter := r.Group("/api")
	{
		LogController := controller.LogController{}
		apiRouter.POST("/register", LogController.Register)
		apiRouter.POST("/login", LogController.Login)
		apiRouter.DELETE("/logout", LogController.Logout)
	}
}
