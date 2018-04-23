package database

type MongoHandler interface {
	Find(interface{}) error
	FindAll(interface{}) error
}
