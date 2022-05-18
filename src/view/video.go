package view

import (
	"github.com/chenxuan520/dousheng/controller"
	"github.com/gin-gonic/gin"
)

func InitVideoView(c *gin.RouterGroup) {
	//upload video
	c.POST("/action/", controller.VideoAdd)
	//get post list
	c.GET("/list/", controller.VideoList)
}

func InitVideoFeed(c *gin.RouterGroup) {
	//get feed video
	c.GET("/feed/", controller.VideoFeed)
}
