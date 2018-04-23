package database

type MongoHandler interface {
	FindOne(string, interface{}, interface{}) error
	FindAll(string, interface{}) error
	Insert(string, interface{}) error
	Upsert(string, interface{}, interface{}) error
	Delete(string, interface{}) error
}
