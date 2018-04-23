package database

type MongoHandler interface {
	Find(interface{}) (interface{}, error)
	FindOne(string, interface{}, interface{}) error
	FindAll(string, interface{}) error
	Insert(string, interface{}) error
	Upsert(string, interface{}, interface{}) error
	Delete(string, interface{}) error
}
