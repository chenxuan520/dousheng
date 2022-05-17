package model

import (
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "github.com/chenxuan520/dousheng/config"
)
type Video struct{
	ID			bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	AuID		bson.ObjectId `bson:"aid" json:"aid"`
	AuthorID	int			  `bson:"author_id" json:"author_id"`
	VideoID		int           `bson:"id" json:"id"`
	PlayUrl		string        `bson:"play_url" json:"play_url"`
	CountUrl	string        `bson:"cover_url" json:"cover_url"`
	FavCount	int           `bson:"favorite_count" json:"favorite_count"`
	ComCount	int           `bson:"comment_count" json:"comment_count"`
	IsFav		bool          `bson:"is_favorite" json:"is_favorite"`
	Time		int64		  `bson:"post_time" json:"post_time"`
} 
type VideoInfo struct{
	ID			bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	AuID		bson.ObjectId `bson:"aid" json:"aid"`
	VideoID		int           `bson:"id" json:"id"`
	AuthorID	int			  `bson:"author_id" json:"author_id"`
	PlayUrl		string        `bson:"play_url" json:"play_url"`
	CountUrl	string        `bson:"cover_url" json:"cover_url"`
	FavCount	int           `bson:"favorite_count" json:"favorite_count"`
	ComCount	int           `bson:"comment_count" json:"comment_count"`
	IsFav		bool          `bson:"is_favorite" json:"is_favorite"`
	Time		int64		  `bson:"post_time" json:"post_time"`
	Author		User		  `bson:"author" json:"author"`
} 

func VideoAdd(video Video)error{
	num,err:=VideoCount();
	if err!=nil{
		return err;
	}
	video.VideoID=num;
	return insertData(video,ColVideo);
}

func VideoCount()(int,error){
	return getCountInfo(bson.M{},nil,ColVideo);
}

func VideoFeedList(lastTime int64,limit int)([]VideoInfo,error){
	return getVideoList(bson.M{},nil,lastTime,limit)
}

func VideoList(id string)([]VideoInfo,error){
	return getVideoList(bson.M{
		"aid":bson.ObjectIdHex(id),
	},nil,100,100)
}

func VideoFav(videoId ,action int)error{
	if action==1{
		return changeData(ColVideo,bson.M{
			"id":videoId,
		},bson.M{
			"$inc":bson.M{
				"favorite_count":1,
			},
		})
	}else{
		return changeData(ColVideo,bson.M{
			"id":videoId,
		},bson.M{
			"$inc":bson.M{
				"favorite_count":-1,
			},
		})
	}
}
