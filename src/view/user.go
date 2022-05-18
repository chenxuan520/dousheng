package view

import (
	"github.com/chenxuan520/dousheng/controller"
	"github.com/gin-gonic/gin"
)

func InitUserView(c *gin.RouterGroup) {
	// user register
	c.POST("/register/", controller.UserSignUp)
	// user sign in
	c.POST("/login/", controller.UserSignIn)
	// user info
	c.GET("/", controller.UserMessage)
}
