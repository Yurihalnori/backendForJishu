package controller

import (
	"Jishu/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LogController struct {
}

// 用户注册
func (s LogController) Register(c *gin.Context) {

	///获取数据，用户名以及密码并打包
	name := c.PostForm("name")
	password := c.PostForm("password")
	data := model.User{
		Name:     name,
		Password: password,
	}

	//校验输入数据  (暂时仅有检测是否都存在，如果有需求要改model
	validate := validator.New()
	errs := validate.Struct(data)
	if errs != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "输入信息格式有误",
		})
		return
	}

	//校验该用户名是否重复
	var count int64
	model.DB.Where("name = ?", name).Find(&data).Count(&count).Table("users")
	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "该用户名已注册",
		})
	} else {

		//开始注册写入数据
		model.DB.Create(&data)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"userId":  data.ID,
			"name":    name,
		})

		//记录登录状态
		session := sessions.Default(c)
		session.Set("userId", data.ID)
		session.Save()
	}
}

// 用户登录
func (s LogController) Login(c *gin.Context) {
	//获取数据
	name := c.PostForm("name")
	password := c.PostForm("password")

	//查询是否存在该用户名
	var count int64
	var user model.User
	model.DB.Where("name = ?", name).First(&user).Count(&count).Table("users")

	if count != 0 {

		//密码正确 登录成功
		if password == user.Password {
			c.JSON(http.StatusOK, gin.H{
				"Success": true,
				"UserId":  user.ID,
				"name":    user.Name,
			})
			//记录登录状态
			session := sessions.Default(c)
			session.Set("userId", user.ID)
			session.Save()

			//密码错误
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Success": false,
				"data":    "用户名或密码错误",
			})
		}

		//未找到该用户
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"data":    "用户不存在",
		})
	}
}

// 用户登出
func (s LogController) Logout(c *gin.Context) {

	//获取当前登录状态
	sessions := sessions.Default(c)
	userId := sessions.Get("userId")

	//登出成功
	if userId != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})

		//清除登录状态
		sessions.Clear()
		sessions.Save()

		//登出失败
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}
