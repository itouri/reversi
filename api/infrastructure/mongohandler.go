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

func (mh *MongoHandler) FindOne(collection string, res *interface{}) (interface{}, error) {
	err := mh.Db.C(collection).Find(nil).One(res)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

func (mh *MongoHandler) FindAll(collection string, res *interface{}) (interface{}, error) {
	err := mh.Db.C(collection).Find(nil).All(res)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

func (mh *MongoHandler) Upsert(collection string, query interface{}, upsert interface{}) error {
	_, err := mh.Db.C(collection).Upsert(query, upsert)
}

func (mh *MongoHandler) Insert(collection string, object interface{}) error {
	err := mh.Db.C(collection).Insert(object)
}

func (mh *MongoHandler) Delete(collection string, query interface{}) error {
	err := mh.Db.C(collection).Remove(query)
}
