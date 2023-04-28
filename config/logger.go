package config

import (
	"io"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func createLogger(name string) io.Writer {
	fileName := path.Join("log", name+".log")
	logger, err := rotatelogs.New(fileName+".%Y%m%d",
		// rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour))
	if err != nil {
		panic(err)
	}
	return logger
}

var GinLogger io.Writer
var DatabaseLogger io.Writer

func initLogger() {
	if Config.AppProd {
		GinLogger = createLogger("gin")
		DatabaseLogger = createLogger("database")

		gin.DisableConsoleColor()
		gin.DefaultWriter = GinLogger
	}
}
