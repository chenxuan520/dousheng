package model

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId    int           `bson:"id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Pwd       string        `bson:"password" json:"password"`
	FollCount int           `bson:"follow_count" json:"follow_count"`
	FansCount int           `bson:"follower_count" json:"follower_count"`
	Follower  []int         `bson:"follower" json:"follower"`
	Fans      []int         `bson:"fans" json:"fans"`
	FavVideo  []int         `bson:"fav_video" json:"fav_video"`
	IsFollow  bool          `bson:"is_follow" json:"is_follow"`
}

func UserAdd(user User) error {
	query := bson.M{
		"name": user.Name,
	}
	count, err := getCountInfo(query, nil, ColUser)
	num := UserCount()
	user.UserId = num
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("user exist")
	}
	return insertData(user, ColUser)
}
func UserLoginByNamePwd(email, pwd string) (User, error) {
	query := bson.M{
		"name":     email,
		"password": pwd,
	}
	return getUserInfo(query, nil)
}
func UserInfoByID(id string) (User, error) {
	query := bson.M{
		"_id": bson.ObjectIdHex(id),
	}
	return getUserInfo(query, nil)
}
func UserInfoByUserID(id int) (User, error) {
	query := bson.M{
		"id": id,
	}
	return getUserInfo(query, nil)
}
func UserCount() int {
	num, _ := getCountInfo(bson.M{}, nil, ColUser)
	return num
}
func UserAddFav(token string, video int, action int) error {
	if action == 1 {
		return changeData(ColUser, bson.M{
			"_id": bson.ObjectIdHex(token),
		}, bson.M{
			"$addToSet": bson.M{
				"fav_video": video,
			},
		})
	} else {
		return changeData(ColUser, bson.M{
			"_id": bson.ObjectIdHex(token),
		}, bson.M{
			"$pull": bson.M{
				"fav_video": video,
			},
		})
	}
}
func UserFavList(token string) ([]VideoInfo, error) {
	user, err := UserInfoByID(token)
	if err != nil {
		return []VideoInfo{}, err
	}
	return getVideoList(bson.M{
		"id": bson.M{
			"$in": user.FavVideo,
		},
	}, nil, 100, 100)
}
func UserFansDel(user_id, fan_id int) error {
	err := changeData(ColUser, bson.M{
		"id": user_id,
	}, bson.M{
		"$pull": bson.M{
			"fans": fan_id,
		},
	})
	user, err := getUserInfo(bson.M{
		"id": user_id,
	}, nil)
	num := len(user.Fans)
	err = changeData(ColUser, bson.M{
		"id": user_id,
	}, bson.M{
		"$set": bson.M{
			"follower_count": num,
		},
	})
	err = changeData(ColUser, bson.M{
		"id": fan_id,
	}, bson.M{
		"$pull": bson.M{
			"follower": fan_id,
		},
	})
	user, err = getUserInfo(bson.M{
		"id": fan_id,
	}, nil)
	num = len(user.Fans)
	err = changeData(ColUser, bson.M{
		"id": fan_id,
	}, bson.M{
		"$set": bson.M{
			"follow_count": num,
		},
	})
	return err
}

func UserFansAdd(user_id, fan_id int) error {
	err := changeData(ColUser, bson.M{
		"id": user_id,
	}, bson.M{
		"$addToSet": bson.M{
			"fans": fan_id,
		},
	})
	user, err := getUserInfo(bson.M{
		"id": user_id,
	}, nil)
	num := len(user.Fans)
	err = changeData(ColUser, bson.M{
		"id": user_id,
	}, bson.M{
		"$set": bson.M{
			"follower_count": num,
		},
	})
	err = changeData(ColUser, bson.M{
		"id": fan_id,
	}, bson.M{
		"$addToSet": bson.M{
			"follower": fan_id,
		},
	})
	user, err = getUserInfo(bson.M{
		"id": fan_id,
	}, nil)
	num = len(user.Fans)
	err = changeData(ColUser, bson.M{
		"id": fan_id,
	}, bson.M{
		"$set": bson.M{
			"follow_count": num,
		},
	})
	return err
}
func UserFansList(id int) ([]User, error) {
	user, err := UserInfoByUserID(id)
	if err != nil {
		return []User{}, err
	}
	var list []User
	fmt.Println(len(user.Fans))
	for i := 0; i < len(user.Fans); i++ {
		temp, err := getUserInfo(bson.M{
			"id": user.Fans[i],
		}, nil)
		if err != nil {
			return []User{}, err
		}
		temp.IsFollow = true
		list = append(list, temp)
	}
	return list, err
}
func UserFollerList(id int) ([]User, error) {
	user, err := UserInfoByUserID(id)
	if err != nil {
		return []User{}, err
	}
	var list []User
	fmt.Println(len(user.Follower))
	for i := 0; i < len(user.Follower); i++ {
		temp, err := getUserInfo(bson.M{
			"id": user.Follower[i],
		}, nil)
		if err != nil {
			return []User{}, err
		}
		temp.IsFollow = true
		list = append(list, temp)
	}
	return list, err
}
