package view

import (
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/controller"
)
func InitAssessmentView(c *gin.RouterGroup){
	//commit assessment
	c.POST("/action/",controller.AssessmentAdd);
	//get list
	c.GET("/list/",controller.AssessmentList);
}
