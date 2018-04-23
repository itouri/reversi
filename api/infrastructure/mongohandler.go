package infrastructure

import (
	"fmt"
	"log"

	"../interfaces/database"
	"./config"
	mgo "gopkg.in/mgo.v2"
)

type MongoHandler struct {
	Db      *mgo.Database
	Session *mgo.Session
}

func NewMongoHandler() database.MongoHandler {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", config.DBCONFIG.Host, config.DBCONFIG.Port))
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error:%v", err))
	}
	database = session.DB(config.DBCONFIG.Database)
	if err = session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error:%v", err))
	}
	log.Println("MongoDB initialize success.")
	mgo.SetDebug(APPCONFIG.Debug)
}

func (mh *MongoHandler) FindAll(collection string) (domain.Rooms, error) {
	rooms := new(domain.Rooms)
	err := mh.Db.C(collection).Find(nil).All(rooms)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
