package view

import (
	"github.com/chenxuan520/dousheng/controller"
	"github.com/gin-gonic/gin"
)

func InitRelationView(c *gin.RouterGroup) {
	//action follower or not follower
	c.POST("/action/", controller.UserFansAction)
	//get fans list
	c.GET("/follower/list/", controller.UserFans)
	//get follow list
	c.GET("/follow/list/", controller.UserFollowList)
}
