package model

import (
	mgo "gopkg.in/mgo.v2"
	"github.com/chenxuan520/dousheng/config"
	"log"
)

var (
	mongoSession *mgo.Session
)

func init(){
	log.Println("connecting db!")
	session, err := getMongoSession()
	if err != nil {
		log.Println("MongoDB init error!")
		log.Panic(err)
		return
	}
	mongoSession = session

	log.Println("Database init done!")
}

func getMongoSession() (*mgo.Session, error) {
	mgosession, err := mgo.Dial(config.GlobalConfig.Mongo.Host+":"+config.GlobalConfig.Mongo.Port);
	if err != nil {
		log.Println("Mongodb dial error!")
		log.Panic(err)
		return nil, err
	}
	mgosession.SetMode(mgo.Monotonic, true)
	mgosession.SetPoolLimit(300)
	myDb:=mgosession.DB(config.GlobalConfig.Mongo.DBname);
	err=myDb.Login(config.GlobalConfig.Mongo.User,config.GlobalConfig.Mongo.Pwd);
	if err!=nil{
		log.Println("Login wrong"+config.GlobalConfig.Mongo.User+config.GlobalConfig.Mongo.Pwd);
		log.Panic(err)
		return nil, err
	}
	return mgosession, nil
}


