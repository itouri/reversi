package infrastructure

import (
	"fmt"
	"log"

	"github.com/itouri/reversi/api/infrastructure/config"
	"github.com/itouri/reversi/api/interfaces/database"
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
	database := session.DB(config.DBCONFIG.Database)
	if err = session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error:%v", err))
	}
	log.Println("MongoDB initialize success.")
	mgo.SetDebug(config.APPCONFIG.Debug)
	mongoHandler := &MongoHandler{
		Db:      database,
		Session: session,
	}
	return mongoHandler

}

func (mh *MongoHandler) FindOne(collection string, query interface{}, res interface{}) error {
	return mh.Db.C(collection).Find(query).One(res)
}

func (mh *MongoHandler) FindAll(collection string, res interface{}) error {
	return mh.Db.C(collection).Find(nil).All(res)
}

func (mh *MongoHandler) Upsert(collection string, query interface{}, upsert interface{}) error {
	_, err := mh.Db.C(collection).Upsert(query, upsert)
	return err
}

func (mh *MongoHandler) Insert(collection string, object interface{}) error {
	return mh.Db.C(collection).Insert(object)
}

func (mh *MongoHandler) Delete(collection string, query interface{}) error {
	return mh.Db.C(collection).Remove(query)
}
