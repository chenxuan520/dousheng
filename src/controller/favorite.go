package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/model"
	"github.com/chenxuan520/dousheng/util"
	"strconv"
	"net/http"
)

func FavAdd(c *gin.Context){
	id_str:=c.Query("video_id");
	token:=c.Query("token");
	action_str:=c.Query("action_type");
	if len(id_str)==0||len(token)==0||len(action_str)==0{
		util.Error(c,-1,"messagewrong");
		return;
	}
	id,err:=strconv.Atoi(id_str);
	if err!=nil{
		util.Error(c,-2,err.Error());
		return;
	}
	action,err:=strconv.Atoi(action_str);
	if err!=nil{
		util.Error(c,-2,err.Error());
		return;
	}
	err=model.UserAddFav(token,id,action);
	if err!=nil{
		util.Error(c,-3,err.Error());
		return;
	}
	err=model.VideoFav(id,action);
	if err!=nil{
		util.Error(c,-3,err.Error());
		return;
	}
	util.Success(c,"ok");
}

func FavList(c *gin.Context){
	token:=c.Query("token");
	if len(token)==0{
		util.Error(c,-1,"token wrong");
		return;
	}
	list,err:=model.UserFavList(token);
	if err!=nil{
		util.Error(c,-2,err.Error());
		return;
	}
	if list==nil{
		c.JSON(http.StatusOK,gin.H{
			"status_code":0,
			"status_msg":"ok",
			"video_list":[]model.VideoInfo{},
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status_code":0,
			"status_msg":"ok",
			"video_list":list,
		})
	}
}
