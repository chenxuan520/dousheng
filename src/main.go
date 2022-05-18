package main

import ( 
	"github.com/gin-gonic/gin"
	"github.com/chenxuan520/dousheng/view"
)

func initWebApi(r *gin.Engine){
	userApi:=r.Group("/douyin/user")
	view.InitUserView(userApi)

	videoApi:=r.Group("/douyin/publish");
	view.InitVideoView(videoApi);

	feedApi:=r.Group("/douyin");
	view.InitVideoFeed(feedApi);

	favApi:=r.Group("/douyin/favorite");
	view.InitFavView(favApi);

	assApi:=r.Group("/douyin/comment");
	view.InitAssessmentView(assApi);

	relationApi:=r.Group("/douyin/relation");
	view.InitRelationView(relationApi);
}
func main(){
	r:=gin.Default();
	initWebApi(r);
	r.Run(":5200");
}
