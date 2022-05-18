package view

import (
	"github.com/chenxuan520/dousheng/controller"
	"github.com/gin-gonic/gin"
)

func InitFavView(c *gin.RouterGroup) {
	//video fav
	c.POST("/action/", controller.FavAdd)
	//get user fav list
	c.GET("/list/", controller.FavList)
}
