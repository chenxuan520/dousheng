package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/model"
	"github.com/chenxuan520/dousheng/util"
	"net/http"
	"fmt"
)
func UserSignIn(c *gin.Context){
	email:=c.Query("username");
	pwd:=c.Query("password");
	pwd=util.MD5(pwd);
	user,err:=model.UserLoginByNamePwd(email,pwd);
	if err!=nil{
		util.Error(c,-4,err.Error());
		return;
	}
	c.JSON(http.StatusOK,gin.H{
		"status_code":0,
		"status_msg":"ok",
		"user_id":user.UserId,
		"token":user.ID,
	})
}
func UserSignUp(c *gin.Context){
	name:=c.Query("username");
	pwd:=c.Query("password");
	pwd=util.MD5(pwd);
	if len(name)==0||len(pwd)==0{
		util.Error(c,-1,"argv wrong");
		return;
	}
	user:=model.User{};
	user.Name=name;
	user.Pwd=pwd;
	err:=model.UserAdd(user);
	if err!=nil{
		util.Error(c,-2,err.Error());
		return;
	}
	user,err=model.UserLoginByNamePwd(name,pwd);
	if err!=nil{
		util.Error(c,-3,err.Error());
		return;
	}
	c.JSON(http.StatusOK,gin.H{
		"status_code":0,
		"status_msg":"ok",
		"user_id":user.UserId,
		"token":user.ID,
	})
} 
func UserMessage(c *gin.Context){
	id:=c.Query("token");
	if len(id)==0{
		util.Error(c,-1,"token wrong");
		return;
	}
	fmt.Println(id);
	user,err:=model.UserInfoByID(id);
	if err!=nil{
		util.Error(c,-2,err.Error());
		fmt.Print(err);
		return;
	}
	util.SuccUserInfo(c,"ok",user);
}
func UserFansAdd(c *gin.Context){

}
