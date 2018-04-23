package interfaces

type MongoHandler interface {
	FindAll(interface{}) error
}
