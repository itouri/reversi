package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session  *mgo.Session
	database *mgo.Database
)

type DbBase struct{}

func init() {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", DBCONFIG.Host, DBCONFIG.Port))
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error:%v", err))
	}
	database = session.DB(DBCONFIG.Database)
	if err = session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error:%v", err))
	}
	log.Println("MongoDB initialize success.")
	mgo.SetDebug(APPCONFIG.Debug)
}

func (d *DbBase) Session() *mgo.Session {
	return session
}

func (d *DbBase) Database() *mgo.Database {
	return database
}

func (d *DbBase) Collection(collectionName string) *mgo.Collection {
	return d.Database().C(collectionName)
}

func (d *DbBase) Find(collectionName string, query interface{}) *mgo.Query {
	return d.Database().C(collectionName).Find(query)
}

func (d *DbBase) FindSelect(collectionName string, query, selector interface{}) *mgo.Query {
	return d.Database().C(collectionName).Find(query).Select(selector)
}

func (d *DbBase) FindId(collectionName string, id bson.ObjectId) *mgo.Query {
	return d.Database().C(collectionName).FindId(id)
}
