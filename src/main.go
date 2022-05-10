package main

import ( 
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/controller"
	"net/http"
)

func main(){
	r:=gin.Default();
	r.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"gin ok");
	});
	r.Run(":5200");
}
