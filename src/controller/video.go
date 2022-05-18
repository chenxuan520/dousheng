package controller

import (
	"fmt"
	"github.com/chenxuan520/dousheng/config"
	"github.com/chenxuan520/dousheng/model"
	"github.com/chenxuan520/dousheng/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func VideoAdd(c *gin.Context) {
	token := c.PostForm("token")
	if len(token) == 0 {
		util.Error(c, -1, "token wrong")
		fmt.Println("token wrong")
		return
	}
	user, err := model.UserInfoByID(token)
	if err != nil {
		fmt.Println(err)
		util.Error(c, -2, "wrong")
		return
	}
	video := model.Video{}
	video.AuID = user.ID
	video.IsFav = false
	fmt.Println(user)
	header, err := c.FormFile("data")
	if err != nil {
		util.Error(c, -2, err.Error())
		fmt.Println(err)
		return
	}
	t := time.Now().Unix()
	str := strconv.FormatInt(t, 15)
	dst := config.GlobalConfig.SavePath + user.ID.String() + header.Filename + str
	fmt.Println(dst)
	if err := c.SaveUploadedFile(header, dst); err != nil {
		util.Error(c, -3, err.Error())
		fmt.Println(err)
		return
	}
	util.Success(c, "ok")
}
func VideoFeed(c *gin.Context) {
	token := c.Query("token")
	if len(token) == 0 {
		token = "-1"
	}
	list, err := model.VideoFeedList(0, 200, token)
	if err != nil {
		util.Error(c, -1, err.Error())
		return
	}
	num := len(list)
	var lastTime int64 = 0
	if num != 0 {
		lastTime = list[num-1].Time - 1
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
		"next_time":   lastTime,
		"video_list":  list,
	})
}

func VideoList(c *gin.Context) {
	id := c.Query("token")
	if len(id) == 0 {
		util.Error(c, -1, "token wrong")
		return
	}
	list, err := model.VideoList(id)
	if err != nil {
		util.Error(c, -2, err.Error())
		fmt.Print(err)
		return
	}
	if list != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "ok",
			"video_list":  list,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "ok",
			"video_list":  []model.VideoInfo{},
		})
	}
}
