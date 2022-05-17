package model

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type User struct{
	ID			bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId		int			  `bson:"id" json:"id"`
	Name		string        `bson:"name" json:"name"`
	Pwd			string        `bson:"password" json:"password"`
	FollCount	int			  `bson:"follow_count" json:"follow_conut"`
	FansCount	int			  `bson:"follower_count" json:"follower_count"`
	Follower	[]string      `bson:"follower" json:"follower"`
	Fans		[]string      `bson:"fans" json:"fans"`
	FavVideo	[]int		  `bson:"fav_video" json:"fav_video"`
	IsFollow	bool		  `bson:"is_follow" json:"is_follow"`
}
func UserAdd(user User)error{
	query:=bson.M{
		"username":user.Name,
	}
	count,err:=getCountInfo(query,nil,ColUser);
	num:=UserCount();
	user.UserId=num;
	if err!=nil{
		return err;
	}
	if count!=0{
		return errors.New("user exist");
	}
	return insertData(user,ColUser);
}
func UserLoginByNamePwd(email,pwd string)(User,error){
	query:=bson.M{
		"username":email,
		"password":pwd,
	}
	return getUserInfo(query,nil);
}
func UserInfoByID(id string)(User,error){
	query:=bson.M{
		"_id":bson.ObjectIdHex(id),
	}
	return getUserInfo(query,nil);
}
func UserCount()(int){
	num,_:=getCountInfo(bson.M{},nil,ColUser);
	return num;
}
func UserAddFav(token string,video int,action int)error{
	if action==1{
		return changeData(ColUser,bson.M{
			"_id":bson.ObjectIdHex(token),
		},bson.M{
			"$addToSet":bson.M{
				"fav_video":video,
			},
		})
	}else{
		return changeData(ColUser,bson.M{
			"_id":bson.ObjectIdHex(token),
		},bson.M{
			"$pull":bson.M{
				"fav_video":video,
			},
		})
	}
}
func UserFavList(token string)([]VideoInfo,error){
	user,err:=UserInfoByID(token);
	if err!=nil{
		return []VideoInfo{},err;
	}
	return getVideoList(bson.M{
		"id":bson.M{
			"$in":user.FavVideo,
		},
	},nil,100,100);
}
