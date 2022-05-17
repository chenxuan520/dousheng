package view

import (
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/controller"
)

func InitFavView(c *gin.RouterGroup){
	//video fav
	c.POST("/action/",controller.FavAdd);
	//get user fav list
	c.GET("/list/",controller.FavList);
}
