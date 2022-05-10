package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func GetUserInfo(c *gin.Context)
{
	c.string(http.StatusOK,"gin ok");
}

