package model

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Assessment struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	AssID    int           `bson:"id" json:"id"`
	AuID     bson.ObjectId `bson:"aid" json:"aid"`
	AuthorID int           `bson:"author_id" json:"author_id"`
	VideoID  int           `bson:"video_id" json:"video_id"`
	Content  string        `bson:"content" json:"content"`
	Time     int64         `bson:"date" json:"date"`
}
type AssessmentInfo struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	AssID    int           `bson:"id" json:"id"`
	AuID     bson.ObjectId `bson:"aid" json:"aid"`
	AuthorID int           `bson:"author_id" json:"author_id"`
	VideoID  int           `bson:"video_id" json:"video_id"`
	Content  string        `bson:"content" json:"content"`
	Time     int64         `bson:"date" json:"date"`
	Date     string        `bson:"create_date" json:"create_date"`
	User     User          `bson:"user" json:"user"`
}

func AssAdd(ass Assessment) error {
	num, err := getCountInfo(bson.M{}, nil, ColAssessment)
	if err != nil {
		return err
	}
	ass.AssID = num
	err = changeData(ColVideo, bson.M{
		"id": ass.VideoID,
	}, bson.M{
		"$inc": bson.M{
			"comment_count": 1,
		},
	})
	if err != nil {
		return err
	}
	return insertData(ass, ColAssessment)
}
func AssDel(ass_id int) error {
	return deleteOne(bson.M{
		"id": ass_id,
	}, ColAssessment)
}

func AssMegGet(video_id int) ([]AssessmentInfo, error) {
	list, err := getAssList(bson.M{
		"video_id": video_id,
	}, nil)
	fmt.Println(video_id)
	if err != nil || list == nil {
		return []AssessmentInfo{}, err
	}
	return list, err
}
