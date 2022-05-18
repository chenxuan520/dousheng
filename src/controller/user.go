package controller

import (
	"fmt"
	"github.com/chenxuan520/dousheng/model"
	"github.com/chenxuan520/dousheng/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserSignIn(c *gin.Context) {
	email := c.Query("username")
	pwd := c.Query("password")
	pwd = util.MD5(pwd)
	user, err := model.UserLoginByNamePwd(email, pwd)
	if err != nil {
		util.Error(c, -4, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
		"user_id":     user.UserId,
		"token":       user.ID,
	})
}
func UserSignUp(c *gin.Context) {
	name := c.Query("username")
	pwd := c.Query("password")
	pwd = util.MD5(pwd)
	if len(name) == 0 || len(pwd) == 0 {
		util.Error(c, -1, "argv wrong")
		return
	}
	user := model.User{}
	user.Name = name
	user.Pwd = pwd
	err := model.UserAdd(user)
	if err != nil {
		util.Error(c, -2, err.Error())
		return
	}
	user, err = model.UserLoginByNamePwd(name, pwd)
	if err != nil {
		util.Error(c, -3, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
		"user_id":     user.UserId,
		"token":       user.ID,
	})
}
func UserMessage(c *gin.Context) {
	id := c.Query("token")
	if len(id) == 0 {
		util.Error(c, -1, "token wrong")
		return
	}
	fmt.Println(id)
	user, err := model.UserInfoByID(id)
	if err != nil {
		util.Error(c, -2, err.Error())
		fmt.Print(err)
		return
	}
	util.SuccUserInfo(c, "ok", user)
}
func UserFansAction(c *gin.Context) {
	token := c.Query("token")
	follower_str := c.Query("to_user_id")
	action := c.Query("action_type")
	if len(token) == 0 || len(follower_str) == 0 || len(action) == 0 {
		util.Error(c, -1, "message wrong")
		return
	}
	user, err := model.UserInfoByID(token)
	user_id := user.UserId
	if err != nil {
		util.Error(c, -1, err.Error())
		return
	}
	follower, err := strconv.Atoi(follower_str)
	if err != nil {
		util.Error(c, -2, err.Error())
		return
	}
	if action == "1" {
		model.UserFansAdd(follower, user_id)
	} else {
		model.UserFansDel(follower, user_id)
	}
	util.Success(c, "ok")
}

func UserFans(c *gin.Context) {
	user_id_str := c.Query("user_id")
	if len(user_id_str) == 0 {
		util.Error(c, -1, "message wrong")
		return
	}
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		util.Error(c, -1, err.Error())
		return
	}
	list, err := model.UserFansList(user_id)
	if err != nil {
		util.Error(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
		"user_list":   list,
	})
}
func UserFollowList(c *gin.Context) {
	user_id_str := c.Query("user_id")
	if len(user_id_str) == 0 {
		util.Error(c, -1, "message wrong")
		return
	}
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		util.Error(c, -1, err.Error())
		return
	}
	list, err := model.UserFansList(user_id)
	if err != nil {
		util.Error(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
		"user_list":   list,
	})
}
