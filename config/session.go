package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitSession(r *gin.Engine) {
	r.Use(sessions.Sessions("user", cookie.NewStore([]byte(Config.AppSecret))))
}
