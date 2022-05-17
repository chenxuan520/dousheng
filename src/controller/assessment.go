package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/model"
	"github.com/chenxuan520/dousheng/util"
	"strconv"
	"time"
	"net/http"
)

func AssessmentAdd(c *gin.Context){
	token:=c.Query("token");
	video_id_str:=c.Query("video_id");
	action_type_str:=c.Query("action_type");
	if len(token)==0||len(video_id_str)==0||len(action_type_str)==0{
		util.Error(c,-1,"token wrong");
		return;
	}
	video_id,err:=strconv.Atoi(video_id_str);
	if err!=nil{
		util.Error(c,-1,err.Error());
		return;
	}
	user,err:=model.UserInfoByID(token);
	if err!=nil{
		util.Error(c,-2,err.Error());
		return;
	}
	if action_type_str=="1"{
		var ass model.Assessment;
		content:=c.Query("comment_text");
		ass.AuID=user.ID;
		ass.Content=content;
		ass.Time=time.Now().Unix();
		ass.AuthorID=user.UserId;
		ass.VideoID=video_id;
		model.AssAdd(ass);
	}else{
		ass_id_str:=c.Query("comment_id");
		ass_id,err:=strconv.Atoi(ass_id_str);
		if err!=nil{
			util.Error(c,-2,err.Error());
			return;
		}
		err=model.AssDel(ass_id);
		if err!=nil{
			util.Error(c,-3,err.Error());
			return;
		}
	}
	util.Success(c,"ok");
}
func AssessmentList(c *gin.Context){
	video_id_str:=c.Query("video_id");
	video_id,err:=strconv.Atoi(video_id_str);
	if err!=nil{
		util.Error(c,-1,err.Error());
		return;
	}
	list,err:=model.AssMegGet(video_id);
	if err!=nil{
		util.Error(c,-2,err.Error());
		return;
	}
	c.JSON(http.StatusOK,gin.H{
		"status_code":0,
		"status_msg":"ok",
		"comment_list":list,
	})
}
