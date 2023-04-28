package controller

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    uint64      `json:"code,omitempty"`
}

type IDUriForm struct {
	ID int `uri:"id" binding:"min=1"`
}

type PagerForm struct {
	Page  int `form:"page" binding:"min=1"`
	Limit int `form:"limit" binding:"min=1,max=20"`
}

const (
	ParamErr gin.ErrorType = iota + 3
	SysErr
	OpErr
	AuthErr
	LevelErr
)

var ErrorMapper = map[uint64]string{
	1: "内部错误",
	2: "公开错误",
	3: "参数错误",
	4: "系统错误",
	5: "操作错误",
	6: "鉴权错误",
	7: "权限错误",
}
