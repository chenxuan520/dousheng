package model

import (
	"fmt"
	"github.com/chenxuan520/dousheng/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	ColUser       = "user"
	ColVideo      = "video"
	ColAssessment = "assessment"
)

var (
	DBName = config.GlobalConfig.Mongo.DBname
)

func getCollection(col string) (collection *mgo.Collection, cls func()) {
	s := mongoSession.Clone()
	c := s.DB(DBName).C(col)
	return c, s.Close
}

func getUserInfo(query, selector interface{}) (User, error) {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(ColUser)

	user := User{}
	q := c.Find(query)

	if selector != nil {
		q = q.Select(selector)
	}

	err := q.One(&user)

	return user, err
}
func getVideoInfo(query, selector interface{}) (Video, error) {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(ColVideo)

	user := Video{}
	q := c.Find(query)

	if selector != nil {
		q = q.Select(selector)
	}

	err := q.One(&user)

	return user, err
}
func changeData(col string, query, update interface{}) error {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(col)

	err := c.Update(query, update)
	return err
}

func getCountInfo(query, selector interface{}, col string) (int, error) {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(col)
	q := c.Find(query)

	if selector != nil {
		q = q.Select(selector)
	}
	return q.Count()
}
func deleteOne(query interface{}, col string) error {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(col)
	return c.Remove(query)
}

func getVideoList(query, selector interface{}, last int64, limit int) ([]VideoInfo, error) {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(ColVideo)
	pipe := []bson.M{
		{
			"$match": query,
		},
		{
			"$sort": bson.M{
				"post_time": -1,
			},
		},
		{
			"$limit": limit,
		},
	}
	pipe = append(pipe, bson.M{
		"$lookup": bson.M{
			"from":         ColUser,
			"localField":   "author_id",
			"foreignField": "id",
			"as":           "author",
		},
	})

	if selector != nil {
		pipe = append(pipe, bson.M{
			"$project": selector,
		})
	}

	var list []VideoInfo
	flag := false
	err := c.Pipe(pipe).All(&list)
	for i := 0; i < len(list); i++ {
		flag = false
		list[i].Author, _ = getUserInfo(bson.M{
			"_id": list[i].AuID,
		}, nil)
		for j := 0; j < len(list[i].Author.FavVideo); j++ {
			if list[i].Author.FavVideo[j] == list[i].VideoID {
				flag = true
				break
			}
		}
		list[i].IsFav = flag
	}
	return list, err
}
func getAssList(query, selector interface{}) ([]AssessmentInfo, error) {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(ColAssessment)
	pipe := []bson.M{
		{
			"$match": query,
		},
	}
	pipe = append(pipe, bson.M{
		"$lookup": bson.M{
			"from":         ColUser,
			"localField":   "author_id",
			"foreignField": "id",
			"as":           "user",
		},
	})

	if selector != nil {
		pipe = append(pipe, bson.M{
			"$project": selector,
		})
	}
	var list []AssessmentInfo
	err := c.Pipe(pipe).All(&list)
	for i := 0; i < len(list); i++ {
		_, month, day := time.Unix(list[i].Time, 0).Date()
		list[i].User, _ = getUserInfo(bson.M{
			"id": list[i].AuthorID,
		}, nil)
		list[i].Date = fmt.Sprintf("%d-%d", month, day)
	}
	return list, err
}

func insertData(query interface{}, col string) error {
	s := mongoSession.Copy()
	defer s.Close()
	c := s.DB(DBName).C(col)
	return c.Insert(query)
}
