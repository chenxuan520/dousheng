package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 返回值
type Response struct {
	Status  int         `json:"status_code"`
	Msg		string		`json:"status_msg"`
}
type ResLogin struct {
	Status  int         `json:"status_code"`
	Msg		string		`json:"status_msg"`
	UserId int			`json:"user_id"`
	Token string		`json:"token"`
}
type ResUserInfo struct {
	Status  int         `json:"status_code"`
	Msg		string		`json:"status_msg"`
	User	interface{}	`json:"user"`
}
type ResVideoList struct {
	Status  int         `json:"status_code"`
	Msg		string		`json:"status_msg"`

}

func Success(c *gin.Context,msg string) {
	c.JSON(http.StatusOK, Response{
		Status: 0,
		Msg: msg,
	})
}

func SuccLogin(c *gin.Context,msg string,id int,token string){
	c.JSON(http.StatusOK,ResLogin{
		Status:0,
		Msg:msg,
		UserId:id,
		Token:token,
	})
}
func SuccUserInfo(c *gin.Context,msg string,user interface{}){
	c.JSON(http.StatusOK,ResUserInfo{
		Status:0,
		Msg:msg,
		User:user,
	})
}

// Error 错误
func Error(c *gin.Context, status int, msg string){
	c.JSON(http.StatusOK, Response{
		Status: status,
		Msg: msg,
	})
}

