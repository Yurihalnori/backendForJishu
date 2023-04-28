package config

import (
	"Jishu/service/validator"
)

func init() {
	initConfig()
	initLogger()
	validator.InitValidator()
}
