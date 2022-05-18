package view

import (
	"github.com/chenxuan520/dousheng/controller"
	"github.com/gin-gonic/gin"
)

func InitAssessmentView(c *gin.RouterGroup) {
	//commit assessment
	c.POST("/action/", controller.AssessmentAdd)
	//get list
	c.GET("/list/", controller.AssessmentList)
}
